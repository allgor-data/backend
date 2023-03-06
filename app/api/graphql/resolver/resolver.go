package resolver

import "github.com/allgor-data/backend/app/usecase"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	CreateUserUsecase usecase.CreateUserUsecase
}
