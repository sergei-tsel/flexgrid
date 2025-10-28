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

type ArticleRepository interface {
	FindById(id int, userId int) (*model.Article, error)
	FindMany(search string, userId int) ([]model.Article, error)
	Create(entity *model.Article) error
	Update(entity *model.Article) error
	UpdateIsPublic(entity *model.Article) error
}

type ArticleRepo struct{}

func (repo *ArticleRepo) FindById(id int, userId int) (*model.Article, error) {
	articleCache, err := getArticleCache(id)

	if articleCache != nil {
		return articleCache, nil
	}

	err = createArticleTable()

	query := `
		SELECT *
		FROM articles
		WHERE id = $1 AND (user_id = $2 OR is_public IS TRUE);
	`

	row, err := db.Postgres.Query(query, id, userId)
	defer row.Close()

	if !row.Next() {
		return nil, fmt.Errorf("article not found: %w", err)
	}

	var article model.Article

	err = row.Scan(&article.Id, &article.UserId, &article.Title, &article.Content, &article.IsPublic, &article.CreatedAt, &article.UpdatedAt)

	if err != nil {
		return nil, fmt.Errorf("failing to read data from database: %w", err)
	}

	return &article, nil
}

func (repo *ArticleRepo) FindMany(search string, userId int) ([]model.Article, error) {
	query := `
    	SELECT articles.*
		FROM articles 
		JOIN users ON articles.user_id = users.id 
		WHERE (LOWER(articles.title) LIKE '%' || LOWER($1) || '%' 
		       OR LOWER(users.email) LIKE '%' || LOWER($1) || '%') 
		  	  AND (articles.user_id = $2 OR articles.is_public IS TRUE);
	`

	rows, err := db.Postgres.Query(query, search, userId)
	defer rows.Close()

	if err != nil {
		return nil, fmt.Errorf("articles not found: %w", err)
	}

	var articles []model.Article

	for rows.Next() {
		var article model.Article

		err = rows.Scan(&article.Id, &article.UserId, &article.Title, &article.Content, &article.IsPublic, &article.CreatedAt, &article.UpdatedAt)

		if err != nil {
			return nil, fmt.Errorf("failing to read data from database: %w", err)
		}

		articles = append(articles, article)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("failing to read data from database: %w", err)
	}

	return articles, nil
}

func (repo *ArticleRepo) Create(entity *model.Article) error {
	err := createArticleTable()

	query := `
		INSERT INTO articles (user_id, title, content, is_public, created_at, updated_at) 
		VALUES ($1, $2, $3, $4, $5, $6)
	`

	_, err = db.Postgres.Query(
		query,
		entity.UserId,
		entity.Title,
		entity.Content,
		entity.IsPublic,
		entity.CreatedAt,
		entity.UpdatedAt,
	)

	if err != nil {
		return fmt.Errorf("failed to create article: %w", err)
	}

	setArticleCache(entity)

	return nil
}

func (repo *ArticleRepo) Update(entity *model.Article) error {
	query := `
		UPDATE articles 
		SET title = $2, content = $3, updated_at = $4
		WHERE id = $1
	`

	_, err := db.Postgres.Query(
		query,
		entity.Id,
		entity.Title,
		entity.Content,
		entity.UpdatedAt,
	)

	if err != nil {
		return fmt.Errorf("failed to update article: %w", err)
	}

	setArticleCache(entity)

	return nil
}

func (repo *ArticleRepo) UpdateIsPublic(entity *model.Article) error {
	query := `
		UPDATE articles
		SET is_public = $2
		WHERE id = $1;
	`

	_, err := db.Postgres.Query(
		query,
		entity.Id,
		entity.IsPublic,
	)

	if err != nil {
		return fmt.Errorf("failed to update article.is_public: %w", err)
	}

	return nil
}

func createArticleTable() error {
	query := `
		CREATE TABLE IF NOT EXISTS articles (
    		id SERIAL PRIMARY KEY,
    		user_id INTEGER NOT NULL,											  -- уникальный идентификатор автора
    		title TEXT NOT NULL,                    							  -- эмейл
    		content TEXT NOT NULL,                 								  -- пароль
    		is_public BOOLEAN NOT NULL DEFAULT FALSE,							  -- флаг опубликования
    		created_at TIMESTAMPTZ DEFAULT now(),
		    updated_at TIMESTAMPTZ DEFAULT now(),
		    
			CONSTRAINT fk_articles_user_id_users FOREIGN KEY (user_id) REFERENCES users(id)
		)
	`

	_, err := db.Postgres.Query(query)

	if err != nil {
		return fmt.Errorf("failed to create article table if not exist: %w", err)
	}

	return nil
}

func setArticleCache(article *model.Article) {
	jsonBytes, _ := json.Marshal(article)

	db.Redis.Set(
		context.Background(),
		fmt.Sprintf("article:%d", article.Id),
		jsonBytes,
		3*time.Second,
	)
}

func getArticleCache(articleId int) (*model.Article, error) {
	result, err := db.Redis.Get(
		context.Background(),
		fmt.Sprintf("article:%d", articleId),
	).Result()

	if errors.Is(err, redis.Nil) {
		return nil, err
	}

	var article model.Article

	err = json.Unmarshal([]byte(result), &article)

	if err != nil {
		return nil, err
	}

	return &article, nil
}
