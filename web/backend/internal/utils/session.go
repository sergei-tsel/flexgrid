package utils

import (
	"fmt"
	"net/http"

	"github.com/gorilla/sessions"
)

const sessionName = "current_user_id"

var sessionStore = sessions.NewCookieStore()

func Authenticate(w http.ResponseWriter, r *http.Request, token string) error {
	session, err := sessionStore.Get(r, sessionName)

	if err != nil {
		return err
	}

	session.Values[sessionName], err = ParseToken(token)

	err = session.Save(r, w)

	if err != nil {
		return err
	}

	return nil
}

func GetAuthenticatedUserId(r *http.Request) (*int, error) {
	session, err := sessionStore.Get(r, sessionName)

	if err != nil {
		return nil, err
	}

	userId := session.Values[sessionName]

	if userId == nil {
		return nil, fmt.Errorf("user not found in session")
	}

	return userId.(*int), nil
}
