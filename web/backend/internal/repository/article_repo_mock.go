package repository

import (
	"flexgrid/internal/model"
	"strings"
)

type ArticleRepoMock struct {
	Articles map[int]*model.Article
	Count    int
}

var ArticleAuthorMock UserRepoMock

func (repo ArticleRepoMock) FindById(id int, userId int) (*model.Article, error) {
	if repo.Articles[id] != nil && (repo.Articles[id].UserId == userId || repo.Articles[id].IsPublic) {
		return repo.Articles[id], nil
	}

	return nil, nil
}

func (repo ArticleRepoMock) FindMany(search string, userId int) ([]model.Article, error) {
	var articles []model.Article

	var usersMap = map[int]*model.User{}

	for _, user := range ArticleAuthorMock.Users {
		usersMap[user.Id] = user
	}

	for _, article := range repo.Articles {
		if article.UserId != userId && !article.IsPublic {
			continue
		}

		if strings.Contains(article.Title, search) || strings.Contains(usersMap[article.UserId].Email, search) {
			articles = append(articles, *article)
		}
	}

	return articles, nil
}

func (repo ArticleRepoMock) Create(entity *model.Article) error {
	repo.Count++

	entity.Id = repo.Count

	repo.Articles[entity.Id] = entity

	return nil
}

func (repo ArticleRepoMock) Update(entity *model.Article) error {
	repo.Articles[entity.Id].Title = entity.Title
	repo.Articles[entity.Id].Content = entity.Content
	repo.Articles[entity.Id].UpdatedAt = entity.UpdatedAt

	return nil
}

func (repo ArticleRepoMock) UpdateIsPublic(entity *model.Article) error {
	repo.Articles[entity.Id].IsPublic = entity.IsPublic

	return nil
}
