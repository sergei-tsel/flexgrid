package repository

import "flexgrid/internal/model"

type UserRepoMock struct {
	Users map[string]*model.User
	Count int
}

func (repo UserRepoMock) FindById(id int) (*model.User, error) {
	for _, user := range repo.Users {
		if user.Id == id {
			return user, nil
		}
	}

	return nil, nil
}

func (repo UserRepoMock) FindByEmail(email string) (*model.User, error) {
	if repo.Users[email] != nil {
		return repo.Users[email], nil
	}

	return nil, nil
}

func (repo UserRepoMock) Create(entity *model.User) error {
	repo.Count++

	entity.Id = repo.Count

	repo.Users[entity.Email] = entity

	return nil
}
