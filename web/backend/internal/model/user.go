package model

import "time"

type User struct {
	Id        int       // уникальный идентификатор
	Email     string    // логин
	Password  string    // пароль
	CreatedAt time.Time // дата создания пользователя
}
