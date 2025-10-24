package router

import (
	"encoding/json"
	"flexgrid/internal/repository"
	"flexgrid/internal/service"
	"flexgrid/internal/utils"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func NewRouter() *chi.Mux {
	r := chi.NewRouter()

	r.Get("/ping", pong)

	r.Post("/auth/register", register)

	r.Post("/auth/login", login)

	r.Post("/auth/logout", logout)

	r.Get("/auth/me", me)

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
