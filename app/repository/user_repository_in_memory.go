package repository

import (
	"github.com/allgor-data/backend/app/entity"
)

type UserRepositoryInMemory struct {
	users []*entity.User
}

var _ UserRepository = (*UserRepositoryInMemory)(nil)

func NewUserRepositoryInMemory() *UserRepositoryInMemory {
	return &UserRepositoryInMemory{
		users: make([]*entity.User, 0),
	}
}

func (i *UserRepositoryInMemory) CreateUser(user *entity.User) error {
	i.users = append(i.users, user)
	return nil
}
