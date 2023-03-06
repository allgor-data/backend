package repository

import (
	"github.com/allgor-data/backend/app/entity"
)

type InMemoryUserRepository struct {
	users []*entity.User
}

var _ UserRepository = (*InMemoryUserRepository)(nil)

func NewInMemoryUserRepository() *InMemoryUserRepository {
	return &InMemoryUserRepository{
		users: make([]*entity.User, 0),
	}
}

func (i *InMemoryUserRepository) CreateUser(user *entity.User) error {
	i.users = append(i.users, user)
	return nil
}
