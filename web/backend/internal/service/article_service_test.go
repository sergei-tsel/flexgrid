package service

import (
	"flexgrid/internal/model"
	"flexgrid/internal/repository"
	"fmt"
	"reflect"
	"testing"
)

func TestGetOneArticle(t *testing.T) {
	articleRepo := repository.ArticleRepoMock{
		Articles: make(map[int]*model.Article),
	}

	req := CreateArticleRequest{
		UserId:  1,
		Title:   "Тестовая статья",
		Content: "Тестовый макет контента",
	}

	article, _ := CreateArticle(
		req,
		articleRepo,
	)

	type args struct {
		articleId  int
		authUserId int
	}

	tests := []struct {
		name    string
		args    args
		want    *model.Article
		wantErr bool
	}{
		{
			name: "Получение статьи",
			args: args{
				articleId:  article.Id,
				authUserId: req.UserId,
			},
			want: &model.Article{
				Title:   article.Title,
				Content: article.Content,
			},
			wantErr: false,
		},
		{
			name: "Недоступная статья",
			args: args{
				articleId:  article.Id,
				authUserId: 0,
			},
			want:    nil,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetOneArticle(articleRepo, tt.args.articleId, tt.args.authUserId)

			if (err != nil) != tt.wantErr {
				t.Errorf("GetOneArticle() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if tt.want == nil && got != nil {
				t.Errorf("GetOneArticle() = %v, want %v", got, tt.want)
			}

			if tt.want != nil && !reflect.DeepEqual(&model.Article{
				Title:   got.Title,
				Content: got.Content,
			}, &model.Article{
				Title:   tt.want.Title,
				Content: tt.want.Content,
			}) {
				t.Errorf("GetOneArticle() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetManyArticles(t *testing.T) {
	repository.ArticleAuthorMock = repository.UserRepoMock{
		Users: map[string]*model.User{},
		Count: 0,
	}

	articleRepo := repository.ArticleRepoMock{
		Articles: make(map[int]*model.Article),
		Count:    0,
	}

	for i := 0; i < 2; i++ {
		user, _ := CreateUser(
			repository.ArticleAuthorMock,
			"пользователь@тест"+fmt.Sprintf("%d", i),
			"123123",
		)

		repository.ArticleAuthorMock.Count++

		article, _ := CreateArticle(
			CreateArticleRequest{
				UserId:  user.Id,
				Title:   "Тестовая статья",
				Content: "Тестовый макет контента",
			},
			articleRepo,
		)

		articleRepo.Count++

		if i == 1 {
			PublishArticle(articleRepo, article.Id, user.Id)
		}
	}

	type args struct {
		req        GetManyArticlesRequest
		authUserId int
	}

	tests := []struct {
		name    string
		args    args
		want    []model.Article
		wantErr bool
	}{
		{
			name: "Получение своих статей по заголовку",
			args: args{
				req: GetManyArticlesRequest{
					Search: articleRepo.Articles[1].Title[:3],
				},
				authUserId: articleRepo.Articles[1].UserId,
			},
			want: []model.Article{
				*articleRepo.Articles[1],
				*articleRepo.Articles[2],
			},
			wantErr: false,
		},
		{
			name: "Получение своих статей по эмейлу автора",
			args: args{
				req: GetManyArticlesRequest{
					Search: "@те",
				},
				authUserId: articleRepo.Articles[1].UserId,
			},
			want: []model.Article{
				*articleRepo.Articles[1],
				*articleRepo.Articles[2],
			},
			wantErr: false,
		},
		{
			name: "Получение опубликованных статей по заголовку",
			args: args{
				req: GetManyArticlesRequest{
					Search: articleRepo.Articles[2].Title[:3],
				},
				authUserId: articleRepo.Articles[2].UserId,
			},
			want: []model.Article{
				*articleRepo.Articles[1],
				*articleRepo.Articles[2],
			},
			wantErr: false,
		},
		{
			name: "Получение опубликованных статей по эмейлу автора",
			args: args{
				req: GetManyArticlesRequest{
					Search: "@те",
				},
				authUserId: articleRepo.Articles[2].UserId,
			},
			want: []model.Article{
				*articleRepo.Articles[1],
				*articleRepo.Articles[2],
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetManyArticles(tt.args.req, articleRepo, articleRepo.Articles[1].UserId)

			if (err != nil) != tt.wantErr {
				t.Errorf("GetManyArticles() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetManyArticles() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCreateArticle(t *testing.T) {
	articleRepo := repository.ArticleRepoMock{
		Articles: make(map[int]*model.Article),
	}

	type args struct {
		req CreateArticleRequest
	}

	testArgs := args{
		req: CreateArticleRequest{
			UserId:  1,
			Title:   "Тестовая статья",
			Content: "Тестовый макет контента",
		},
	}

	tests := []struct {
		name    string
		args    args
		want    *model.Article
		wantErr bool
	}{
		{
			name: "Создание статьи",
			args: testArgs,
			want: &model.Article{
				Title:   testArgs.req.Title,
				Content: testArgs.req.Content,
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CreateArticle(tt.args.req, articleRepo)

			if (err != nil) != tt.wantErr {
				t.Errorf("CreateArticle() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(&model.Article{
				Title:   got.Title,
				Content: got.Content,
			}, &model.Article{
				Title:   tt.want.Title,
				Content: tt.want.Content,
			}) {
				t.Errorf("CreateArticle() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpdateArticle(t *testing.T) {
	articleRepo := repository.ArticleRepoMock{
		Articles: make(map[int]*model.Article),
	}

	req := CreateArticleRequest{
		UserId:  1,
		Title:   "Тестовая статья",
		Content: "Тестовый макет контента",
	}

	article, _ := CreateArticle(
		req,
		articleRepo,
	)

	type args struct {
		req        UpdateArticleRequest
		articleId  int
		authUserId int
	}

	testArgs := args{
		req: UpdateArticleRequest{
			Title:   "Тестовая статья изменена",
			Content: "Тестовый макет контента изменён",
		},
		articleId:  article.Id,
		authUserId: article.UserId,
	}

	tests := []struct {
		name    string
		args    args
		want    *model.Article
		wantErr bool
	}{
		{
			name: "Изменение статьи",
			args: testArgs,
			want: &model.Article{
				Title:   testArgs.req.Title,
				Content: testArgs.req.Content,
			},
			wantErr: false,
		},
		{
			name: "Несуществующая статья",
			args: args{
				req: UpdateArticleRequest{
					Title:   "",
					Content: "",
				},
				articleId: 0,
			},
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := UpdateArticle(tt.args.req, articleRepo, tt.args.articleId, tt.args.authUserId)

			if (err != nil) != tt.wantErr {
				t.Errorf("UpdateArticle() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if tt.want != nil && !reflect.DeepEqual(&model.Article{
				Title:   got.Title,
				Content: got.Content,
			}, &model.Article{
				Title:   tt.want.Title,
				Content: tt.want.Content,
			}) {
				t.Errorf("UpdateArticle() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPublishArticle(t *testing.T) {
	articleRepo := repository.ArticleRepoMock{
		Articles: make(map[int]*model.Article),
	}

	req := CreateArticleRequest{
		UserId:  1,
		Title:   "Тестовая статья",
		Content: "Тестовый макет контента",
	}

	article, _ := CreateArticle(
		req,
		articleRepo,
	)

	type args struct {
		articleId  int
		authUserId int
	}

	tests := []struct {
		name    string
		args    args
		want    *model.Article
		wantErr bool
	}{
		{
			name: "Публикация статьи",
			args: args{
				articleId:  article.Id,
				authUserId: article.UserId,
			},
			want: &model.Article{
				IsPublic: true,
			},
			wantErr: false,
		},
		{
			name: "Несуществующая статья",
			args: args{
				articleId: 0,
			},
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := PublishArticle(articleRepo, tt.args.articleId, article.UserId)

			if (err != nil) != tt.wantErr {
				t.Errorf("PublishArticle() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if tt.want != nil && !reflect.DeepEqual(&model.Article{
				IsPublic: got.IsPublic,
			}, &model.Article{
				IsPublic: tt.want.IsPublic,
			}) {
				t.Errorf("PublishArticle() = %v, want %v", got, tt.want)
			}
		})
	}
}
