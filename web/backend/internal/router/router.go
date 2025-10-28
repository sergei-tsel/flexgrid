package router

import (
	"encoding/json"
	"flexgrid/internal/repository"
	"flexgrid/internal/service"
	"flexgrid/internal/utils"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func NewRouter() *chi.Mux {
	r := chi.NewRouter()

	r.Get("/ping", pong)

	r.Post("/auth/register", register)

	r.Post("/auth/login", login)

	r.Post("/auth/logout", logout)

	r.Get("/auth/me", me)

	r.Post("/article", createArticle)

	r.Post("/article/search", getManyArticles)

	r.Get("/article/{articleId}", getOneArticle)

	r.Post("/article/{articleId}", updateArticle)

	r.Post("/article/{articleId}/publish", publishArticle)

	return r
}

func pong(response http.ResponseWriter, request *http.Request) {
	response.Write([]byte("pong"))
}

func register(w http.ResponseWriter, r *http.Request) {
	isAuth, err := service.CheckAuthentication(w, r)

	if isAuth {
		http.Error(w, "Only guest requests are allowed", http.StatusForbidden)
		return
	}

	var req service.RegisterRequest

	err = json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		http.Error(w, "Invalid request: "+err.Error(), http.StatusBadRequest)
		return
	}

	cookie, err := service.Register(req, &repository.UserRepo{})

	if err != nil {
		http.Error(w, "Invalid request: "+err.Error(), http.StatusBadRequest)
		return
	}

	http.SetCookie(w, cookie)

	utils.RespondJSON(w, nil, http.StatusOK)
}

func login(w http.ResponseWriter, r *http.Request) {
	isAuth, err := service.CheckAuthentication(w, r)

	if isAuth {
		http.Error(w, "Only guest requests are allowed", http.StatusForbidden)
		return
	}

	var req service.LoginRequest

	err = json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		http.Error(w, "Invalid request: "+err.Error(), http.StatusBadRequest)
		return
	}

	cookie, err := service.Login(req, &repository.UserRepo{})

	if err != nil {
		http.Error(w, "Invalid request: "+err.Error(), http.StatusBadRequest)
		return
	}

	http.SetCookie(w, cookie)

	utils.RespondJSON(w, nil, http.StatusOK)
}

func logout(w http.ResponseWriter, r *http.Request) {
	isAuth, _ := service.CheckAuthentication(w, r)

	if !isAuth {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	cookie := service.Logout()

	http.SetCookie(w, cookie)

	utils.RespondJSON(w, nil, http.StatusOK)
}

func me(w http.ResponseWriter, r *http.Request) {
	isAuth, err := service.CheckAuthentication(w, r)

	userId, err := utils.GetAuthenticatedUserId(r)

	if !isAuth || userId == nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	user, err := service.GetOneUser(&repository.UserRepo{}, *userId)

	if err != nil {
		http.Error(w, "Invalid request: "+err.Error(), http.StatusBadRequest)
		return
	}

	utils.RespondJSON(w, user, http.StatusOK)
}

func createArticle(w http.ResponseWriter, r *http.Request) {
	isAuth, err := service.CheckAuthentication(w, r)

	if !isAuth {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	var req service.CreateArticleRequest

	err = json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		http.Error(w, "Invalid request: "+err.Error(), http.StatusBadRequest)
		return
	}

	article, err := service.CreateArticle(req, &repository.ArticleRepo{})

	if err != nil {
		http.Error(w, "Invalid request: "+err.Error(), http.StatusBadRequest)
		return
	}

	utils.RespondJSON(w, article, http.StatusOK)
}

func getManyArticles(w http.ResponseWriter, r *http.Request) {
	isAuth, err := service.CheckAuthentication(w, r)

	if !isAuth {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	var req service.GetManyArticlesRequest

	err = json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		http.Error(w, "Invalid request: "+err.Error(), http.StatusBadRequest)
		return
	}

	authUserId, _ := utils.GetAuthenticatedUserId(r)

	articles, err := service.GetManyArticles(req, &repository.ArticleRepo{}, *authUserId)

	if err != nil {
		http.Error(w, "Articles not found: "+err.Error(), http.StatusNotFound)
		return
	}

	utils.RespondJSON(w, articles, http.StatusOK)
}

func getOneArticle(w http.ResponseWriter, r *http.Request) {
	isAuth, err := service.CheckAuthentication(w, r)

	if !isAuth {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	stringArticleId := chi.URLParam(r, "articleId")
	articleId, err := strconv.Atoi(stringArticleId)

	if err != nil {
		http.Error(w, "Invalid request: "+err.Error(), http.StatusBadRequest)
		return
	}

	authUserId, _ := utils.GetAuthenticatedUserId(r)

	article, err := service.GetOneArticle(&repository.ArticleRepo{}, articleId, *authUserId)

	if err != nil {
		http.Error(w, "Article not found: "+err.Error(), http.StatusNotFound)
		return
	}

	utils.RespondJSON(w, article, http.StatusOK)
}

func updateArticle(w http.ResponseWriter, r *http.Request) {
	isAuth, err := service.CheckAuthentication(w, r)

	if !isAuth {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	stringArticleId := chi.URLParam(r, "articleId")
	articleId, err := strconv.Atoi(stringArticleId)

	if err != nil {
		http.Error(w, "Invalid request: "+err.Error(), http.StatusBadRequest)
		return
	}

	var req service.UpdateArticleRequest

	err = json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		http.Error(w, "Invalid request: "+err.Error(), http.StatusBadRequest)
		return
	}

	authUserId, _ := utils.GetAuthenticatedUserId(r)

	article, err := service.UpdateArticle(req, &repository.ArticleRepo{}, articleId, *authUserId)

	if err != nil {
		http.Error(w, "Article not found: "+err.Error(), http.StatusNotFound)
		return
	}

	utils.RespondJSON(w, article, http.StatusOK)
}

func publishArticle(w http.ResponseWriter, r *http.Request) {
	isAuth, err := service.CheckAuthentication(w, r)

	if !isAuth {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	stringArticleId := chi.URLParam(r, "articleId")
	articleId, err := strconv.Atoi(stringArticleId)

	if err != nil {
		http.Error(w, "Invalid request: "+err.Error(), http.StatusBadRequest)
		return
	}

	authUserId, _ := utils.GetAuthenticatedUserId(r)

	article, err := service.PublishArticle(&repository.ArticleRepo{}, articleId, *authUserId)

	if err != nil {
		http.Error(w, "Article not found: "+err.Error(), http.StatusNotFound)
		return
	}

	utils.RespondJSON(w, article, http.StatusOK)
}
