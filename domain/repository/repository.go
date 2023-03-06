package repository

import "github.com/allgor-data/backend/domain/entity"

type UserRepository interface {
	CreateUser(*entity.User) error
}
