package service

import (
	"flexgrid/internal/model"
	"flexgrid/internal/repository"
	"flexgrid/internal/utils"
	"fmt"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

const AuthCookieName = "fg_id"

const cookieLifeTime = 60 * 60 * 24 * 365

type RegisterRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Register(req RegisterRequest, repo repository.UserRepository) (*http.Cookie, error) {
	var user *model.User

	user, err := repo.FindByEmail(req.Email)

	if user == nil {
		user, err = CreateUser(repo, req.Email, req.Password)
	}

	if user == nil {
		return nil, err
	}

	token, err := utils.GenerateToken(user.Id)

	cookie := createLoginCookie(token)

	if err != nil {
		return &cookie, err
	}

	return &cookie, nil
}

func Login(req LoginRequest, repo repository.UserRepository) (*http.Cookie, error) {
	var user *model.User

	user, err := repo.FindByEmail(req.Email)

	if user == nil {
		return nil, fmt.Errorf("user not exists")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))

	if err != nil {
		return nil, err
	}

	token, err := utils.GenerateToken(user.Id)

	cookie := createLoginCookie(token)

	return &cookie, nil
}

func Logout() *http.Cookie {
	cookie := createLogoutCookie()

	return &cookie
}

func CheckAuthentication(w http.ResponseWriter, r *http.Request) (bool, error) {
	cookie, err := r.Cookie(AuthCookieName)

	if cookie == nil {
		return false, err
	}

	err = utils.Authenticate(w, r, cookie.Value)

	if err != nil {
		return true, err
	}

	return true, nil
}

func createLoginCookie(token string) http.Cookie {
	return http.Cookie{
		Name:     AuthCookieName,
		Value:    token,
		Path:     "/",
		MaxAge:   cookieLifeTime,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
		Secure:   true,
	}
}

func createLogoutCookie() http.Cookie {
	return http.Cookie{
		Name:     AuthCookieName,
		Value:    "",
		MaxAge:   -1,
		Path:     "/",
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
		Secure:   true,
	}
}
