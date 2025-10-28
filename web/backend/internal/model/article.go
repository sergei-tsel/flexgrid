package model

import "time"

type Article struct {
	Id        int       // уникальный идентификатор
	UserId    int       // уникальный идентификатор автора
	Title     string    // заголовок
	Content   string    // макет контента
	IsPublic  bool      // флаг опубликования
	CreatedAt time.Time // время создания статьи
	UpdatedAt time.Time // время изменения статьи
}
