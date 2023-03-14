package usecase

import (
	"testing"

	"github.com/allgor-data/backend/app/dto"
	"github.com/allgor-data/backend/app/entity"
	"github.com/allgor-data/backend/app/repository"
	"gotest.tools/v3/assert"
)

func TestCreateUser(t *testing.T) {
	r := repository.NewUserRepositoryInMemory()
	usecase := NewCreateUserUsecase(r)

	input := &dto.CreateUserInput{
		Email:     "johndoe@example.com",
		FirstName: "John",
		LastName:  "Doe",
		Password:  "good_password",
	}

	output, err := usecase.Execute(input)

	assert.NilError(t, err)

	assert.Assert(t, output.UID != "")

	assert.Equal(t, output.Role, entity.UserRoleOperator)

	assert.Equal(t, output.Email, input.Email)
	assert.Equal(t, output.FirstName, input.FirstName)
	assert.Equal(t, output.LastName, input.LastName)
}

func TestCreateUser_EntityValidation(t *testing.T) {
	r := repository.NewUserRepositoryInMemory()
	usecase := NewCreateUserUsecase(r)

	input := &dto.CreateUserInput{
		Email:     "invalid_email",
		FirstName: "John",
		LastName:  "Doe",
		Password:  "good_password",
	}

	output, err := usecase.Execute(input)

	assert.Error(t, err, "invalid e-mail")
	assert.Assert(t, output == nil)
}

func TestCreateUser_ShortPassword(t *testing.T) {
	r := repository.NewUserRepositoryInMemory()
	usecase := NewCreateUserUsecase(r)

	input := &dto.CreateUserInput{
		Email:     "johndoe@example.com",
		FirstName: "John",
		LastName:  "Doe",
		Password:  "1234567",
	}

	output, err := usecase.Execute(input)

	assert.Error(t, err, "password length must be greater or equal than 8")
	assert.Assert(t, output == nil)
}

func TestCreateUser_AlreadyExists(t *testing.T) {
	r := repository.NewUserRepositoryInMemory()
	usecase := NewCreateUserUsecase(r)

	firstUser := &dto.CreateUserInput{
		Email:     "johndoe@example.com",
		FirstName: "John",
		LastName:  "Doe",
		Password:  "good_password",
	}

	output, err := usecase.Execute(firstUser)

	assert.NilError(t, err)
	assert.Assert(t, output.UID != "")

	secondUser := &dto.CreateUserInput{
		Email:     "johndoe@example.com",
		FirstName: "John",
		LastName:  "Doe",
		Password:  "good_password",
	}

	output, err = usecase.Execute(secondUser)

	assert.Error(t, err, "e-mail already registered")
	assert.Assert(t, output == nil)
}
