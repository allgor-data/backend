package usecase

import (
	"testing"

	"github.com/allgor-data/backend/app/entity"
	"github.com/allgor-data/backend/app/repository"
	"gotest.tools/v3/assert"
)

func TestCreateUser(t *testing.T) {
	r := repository.NewUserRepositoryInMemory()
	usecase := NewCreateUserUsecase(r)

	input := &CreateUserInputDTO{
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

func TestCreateUser_Validation(t *testing.T) {
	r := repository.NewUserRepositoryInMemory()
	usecase := NewCreateUserUsecase(r)

	input := &CreateUserInputDTO{
		Email:     "invalid_email",
		FirstName: "John",
		LastName:  "Doe",
		Password:  "good_password",
	}

	output, err := usecase.Execute(input)

	assert.Error(t, err, "invalid e-mail")
	assert.Assert(t, output == nil)
}
