package repository

import "github.com/allgor-data/backend/app/entity"

type UserRepository interface {
	CreateUser(*entity.User) error
	FindUserByEmail(email string) (*entity.User, error)
}
