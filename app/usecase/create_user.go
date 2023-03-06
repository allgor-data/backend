package usecase

import (
	"github.com/allgor-data/backend/app/entity"
	"github.com/allgor-data/backend/app/repository"
)

type CreateUserUsecase struct {
	r repository.UserRepository
}

func NewCreateUserUsecase(r repository.UserRepository) *CreateUserUsecase {
	return &CreateUserUsecase{r}
}

type CreateUserInputDTO struct {
	Email     string
	FirstName string
	LastName  string
	Password  string
}

type CreateUserOutputDTO struct {
	UID       string
	Email     string
	FirstName string
	LastName  string
	Role      entity.UserRole
}

func (c *CreateUserUsecase) Execute(input *CreateUserInputDTO) (*CreateUserOutputDTO, error) {
	user := entity.NewUser()

	user.Email = input.Email
	user.FirstName = input.FirstName
	user.LastName = input.LastName
	user.SetPassword(input.Password)

	err := user.Validate()
	if err != nil {
		return nil, err
	}

	err = c.r.CreateUser(user)
	if err != nil {
		return nil, err
	}

	output := &CreateUserOutputDTO{
		UID:       user.UID,
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Role:      user.Role,
	}

	return output, nil
}
