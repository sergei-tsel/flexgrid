package repository

import (
	"context"
	"encoding/json"
	"errors"
	"flexgrid/internal/db"
	"flexgrid/internal/model"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type UserRepository interface {
	FindById(id int) (*model.User, error)
	FindByEmail(email string) (*model.User, error)
	Create(entity *model.User) error
}

type UserRepo struct{}

func (repo *UserRepo) FindById(id int) (*model.User, error) {
	userCache, err := getUserCache(id)

	if userCache != nil {
		return userCache, nil
	}

	err = createUserTable()

	query := `
		SELECT *
		FROM users
		WHERE id = $1;
	`

	row, err := db.Postgres.Query(query, id)
	defer row.Close()

	if !row.Next() {
		return nil, fmt.Errorf("user not found: %w", err)
	}

	var user model.User

	err = row.Scan(&user.Id, &user.Email, &user.Password, &user.CreatedAt)

	if err != nil {
		return nil, fmt.Errorf("failing to read data from database: %w", err)
	}

	setUserCache(&user)

	return &user, nil
}

func (repo *UserRepo) FindByEmail(email string) (*model.User, error) {
	err := createUserTable()

	query := `
		SELECT *
		FROM users
		WHERE email = $1;
	`

	row, err := db.Postgres.Query(query, email)
	defer row.Close()

	if !row.Next() {
		return nil, fmt.Errorf("failed to find user: %w", err)
	}

	var user model.User

	err = row.Scan(&user.Id, &user.Email, &user.Password, &user.CreatedAt)

	if err != nil {
		return nil, fmt.Errorf("failing to read data from database: %w", err)
	}

	return &user, nil
}

func (repo *UserRepo) Create(entity *model.User) error {
	err := createUserTable()

	query := `
		INSERT INTO users (email, password, created_at) 
		VALUES ($1, $2, $3)
	`

	_, err = db.Postgres.Query(
		query,
		entity.Email,
		entity.Password,
		entity.CreatedAt,
	)

	if err != nil {
		return fmt.Errorf("failed to insert user: %w", err)
	}

	setUserCache(entity)

	return nil
}

func createUserTable() error {
	query := `
		CREATE TABLE IF NOT EXISTS users (
    		id SERIAL PRIMARY KEY,
    		email TEXT NOT NULL,                    	-- эмейл
    		password TEXT NOT NULL,                 	-- пароль
    		created_at TIMESTAMPTZ DEFAULT now(),

    		CONSTRAINT users_email_unique UNIQUE(email)
		)
	`

	_, err := db.Postgres.Query(query)

	if err != nil {
		return fmt.Errorf("failed to create users table if not exist: %w", err)
	}

	return nil
}

func setUserCache(user *model.User) {
	jsonBytes, _ := json.Marshal(user)

	db.Redis.Set(
		context.Background(),
		fmt.Sprintf("user:%d", user.Id),
		jsonBytes,
		3*time.Hour,
	)
}

func getUserCache(userId int) (*model.User, error) {
	result, err := db.Redis.Get(
		context.Background(),
		fmt.Sprintf("user:%d", userId),
	).Result()

	if errors.Is(err, redis.Nil) {
		return nil, err
	}

	var user model.User

	err = json.Unmarshal([]byte(result), &user)

	if err != nil {
		return nil, err
	}

	return &user, nil
}
