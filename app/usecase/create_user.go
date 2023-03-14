package usecase

import (
	"github.com/allgor-data/backend/app/dto"
	"github.com/allgor-data/backend/app/entity"
	"github.com/allgor-data/backend/app/repository"
)

type CreateUserUsecase struct {
	r repository.UserRepository
}

func NewCreateUserUsecase(r repository.UserRepository) *CreateUserUsecase {
	return &CreateUserUsecase{r}
}

func (c *CreateUserUsecase) Execute(input *dto.CreateUserInput) (*dto.CreateUserOutput, error) {
	user := entity.NewUser()

	user.Email = input.Email
	user.FirstName = input.FirstName
	user.LastName = input.LastName

	err := user.SetPassword(input.Password)
	if err != nil {
		return nil, err
	}

	err = user.Validate()
	if err != nil {
		return nil, err
	}

	err = c.r.CreateUser(user)
	if err != nil {
		return nil, err
	}

	output := &dto.CreateUserOutput{
		UID:       user.UID,
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Role:      user.Role,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	return output, nil
}
