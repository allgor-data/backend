package dto

import "github.com/allgor-data/backend/app/entity"

type CreateUserInput struct {
	Email     string `json:"email"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Password  string `json:"password"`
}

type CreateUserOutput struct {
	UID       string          `json:"uid"`
	Email     string          `json:"email"`
	FirstName string          `json:"firstName"`
	LastName  string          `json:"lastName"`
	Role      entity.UserRole `json:"role"`
}
