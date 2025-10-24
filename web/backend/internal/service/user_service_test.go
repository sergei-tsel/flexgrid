package service

import (
	"flexgrid/internal/model"
	"flexgrid/internal/repository"
	"reflect"
	"testing"

	"golang.org/x/crypto/bcrypt"
)

func TestGetOneUser(t *testing.T) {
	userRepo := repository.UserRepoMock{
		Users: make(map[string]*model.User),
	}

	user, _ := CreateUser(
		userRepo,
		"пользователь@тест",
		"123123",
	)

	type args struct {
		userId int
	}

	tests := []struct {
		name    string
		args    args
		want    *model.User
		wantErr bool
	}{
		{
			name: "Получение пользователя",
			args: args{
				userId: user.Id,
			},
			want: &model.User{
				Email:    user.Email,
				Password: user.Password,
			},
			wantErr: false,
		},
		{
			name: "Несуществующий пользователь",
			args: args{
				userId: 0,
			},
			want:    nil,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetOneUser(userRepo, tt.args.userId)

			if (err != nil) != tt.wantErr {
				t.Errorf("GetOneUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if tt.want != nil && !reflect.DeepEqual(&model.User{
				Email:    got.Email,
				Password: got.Password,
			}, &model.User{
				Email:    tt.want.Email,
				Password: tt.want.Password,
			}) {
				t.Errorf("GetOneUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCreateUser(t *testing.T) {
	userRepo := repository.UserRepoMock{
		Users: make(map[string]*model.User),
	}

	type args struct {
		email    string
		password string
	}

	testArgs := args{
		email:    "пользователь@тест",
		password: "123123",
	}

	tests := []struct {
		name    string
		args    args
		want    *model.User
		wantErr bool
	}{
		{
			name: "Создание пользователя",
			args: testArgs,
			want: &model.User{
				Email: testArgs.email,
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CreateUser(userRepo, tt.args.email, tt.args.password)

			if (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if bcrypt.CompareHashAndPassword([]byte(got.Password), []byte(testArgs.password)) == nil && !reflect.DeepEqual(&model.User{
				Email: got.Email,
			}, &model.User{
				Email: tt.want.Email,
			}) {
				t.Errorf("Create() = %v, want %v", got, tt.want)
			}
		})
	}
}
