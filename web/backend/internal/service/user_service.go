package service

import (
	"flexgrid/internal/model"
	"flexgrid/internal/repository"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func GetOneUser(repo repository.UserRepository, userId int) (*model.User, error) {
	var user *model.User

	user, err := repo.FindById(userId)

	if err != nil {
		return user, err
	}

	return user, nil
}

func CreateUser(repo repository.UserRepository, email string, password string) (*model.User, error) {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	user := &model.User{}
	user.Email = email
	user.Password = string(hashPassword[:])
	user.CreatedAt = time.Now()
	err = repo.Create(user)

	if err != nil {
		return user, err
	}

	return user, nil
}
