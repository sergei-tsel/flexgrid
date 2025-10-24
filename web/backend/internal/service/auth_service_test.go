package service

import (
	"flexgrid/internal/model"
	"flexgrid/internal/repository"
	"net/http"
	"reflect"
	"testing"
)

func TestRegister(t *testing.T) {
	userRepo := repository.UserRepoMock{
		Users: make(map[string]*model.User),
	}

	type args struct {
		req RegisterRequest
	}

	tests := []struct {
		name    string
		args    args
		want    *http.Cookie
		wantErr bool
	}{
		{
			name: "Регистрация нового пользователя",
			args: args{
				req: RegisterRequest{
					Email:    "пользователь@тест",
					Password: "123123",
				},
			},
			want: &http.Cookie{
				Name: AuthCookieName,
			},
			wantErr: false,
		},
		{
			name: "Вход существующего пользователя",
			args: args{
				req: RegisterRequest{
					Email:    "пользователь@тест",
					Password: "123123",
				},
			},
			want: &http.Cookie{
				Name: AuthCookieName,
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Register(tt.args.req, userRepo)

			if (err != nil) != tt.wantErr {
				t.Errorf("Register() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(&http.Cookie{
				Name: got.Name,
			}, tt.want) {
				t.Errorf("Register() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLogin(t *testing.T) {
	userRepo := repository.UserRepoMock{
		Users: make(map[string]*model.User),
	}

	password := "123123"

	user, _ := CreateUser(
		userRepo,
		"пользователь@тест",
		password,
	)

	type args struct {
		req LoginRequest
	}

	tests := []struct {
		name    string
		args    args
		want    *http.Cookie
		wantErr bool
	}{
		{
			name: "Авторизация пользователя",
			args: args{
				req: LoginRequest{
					Email:    user.Email,
					Password: password,
				},
			},
			want: &http.Cookie{
				Name: AuthCookieName,
			},
			wantErr: false,
		},
		{
			name: "Неверный пароль",
			args: args{
				req: LoginRequest{
					Email:    user.Email,
					Password: "",
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Несуществующий пользователь",
			args: args{
				req: LoginRequest{
					Email:    "",
					Password: "",
				},
			},
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Login(tt.args.req, userRepo)

			if (err != nil) != tt.wantErr {
				t.Errorf("Email() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if tt.want != nil && !reflect.DeepEqual(
				&http.Cookie{
					Name: got.Name,
				}, tt.want) {
				t.Errorf("Email() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLogout(t *testing.T) {
	userRepo := repository.UserRepoMock{
		Users: make(map[string]*model.User),
	}

	CreateUser(
		userRepo,
		"пользователь@тест",
		"123123",
	)

	tests := []struct {
		name string
		want *http.Cookie
	}{
		{
			name: "Выход пользователя",
			want: &http.Cookie{
				Name: AuthCookieName,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Logout(); !reflect.DeepEqual(
				&http.Cookie{
					Name: got.Name,
				}, tt.want) {
				t.Errorf("Logout() = %v, want %v", got, tt.want)
			}
		})
	}
}
