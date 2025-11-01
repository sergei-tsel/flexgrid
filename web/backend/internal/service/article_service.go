package service

import (
	"flexgrid/internal/model"
	"flexgrid/internal/repository"
	"time"
)

type CreateArticleRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type UpdateArticleRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type GetManyArticlesRequest struct {
	Search string `json:"search"`
}

func GetOneArticle(articleRepo repository.ArticleRepository, articleId int, authUserId int) (*model.Article, error) {
	var article *model.Article

	article, err := articleRepo.FindById(articleId, authUserId)

	if err != nil {
		return article, err
	}

	return article, nil
}

func GetManyArticles(req GetManyArticlesRequest, articleRepo repository.ArticleRepository, authUserId int) ([]model.Article, error) {
	var articles []model.Article

	articles, err := articleRepo.FindMany(req.Search, authUserId)

	if err != nil {
		return articles, err
	}

	return articles, nil
}

func CreateArticle(req CreateArticleRequest, repo repository.ArticleRepository, authUserId int) (*model.Article, error) {
	article := &model.Article{}
	article.UserId = authUserId
	article.Title = req.Title
	article.Content = req.Content
	article.IsPublic = false
	article.CreatedAt = time.Now()
	article.UpdatedAt = time.Now()
	err := repo.Create(article)

	if err != nil {
		return article, err
	}

	return article, nil
}

func UpdateArticle(req UpdateArticleRequest, repo repository.ArticleRepository, articleId int, authUserId int) (*model.Article, error) {
	article, err := GetOneArticle(repo, articleId, authUserId)

	if article == nil {
		return nil, err
	}

	article.Title = req.Title
	article.Content = req.Content
	article.UpdatedAt = time.Now()
	err = repo.Update(article)

	if err != nil {
		return article, err
	}

	return article, nil
}

func PublishArticle(repo repository.ArticleRepository, articleId int, authUserId int) (*model.Article, error) {
	article, err := GetOneArticle(repo, articleId, authUserId)

	if article == nil {
		return nil, err
	}

	article.IsPublic = true
	err = repo.UpdateIsPublic(article)

	if err != nil {
		return article, err
	}

	return article, nil
}
