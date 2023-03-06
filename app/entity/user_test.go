package entity

import (
	"testing"

	"gotest.tools/v3/assert"
)

func makeValidUser() *User {
	user := NewUser()

	user.Email = "johndoe@example.com"
	user.FirstName = "John"
	user.LastName = "Doe"
	user.Role = UserRoleAdmin
	user.PasswordHash = "$2a$12$JhxRcjBAShuAHxey0zsEe.8NR6x.bapNIze6VVeWRIRbLPf.azBXu"

	return user
}

func TestNewUser(t *testing.T) {
	user := NewUser()

	assert.Assert(t, user.UID != "")

	assert.Equal(t, user.Role, UserRoleOperator)

	assert.Equal(t, user.Email, "")
	assert.Equal(t, user.FirstName, "")
	assert.Equal(t, user.LastName, "")
	assert.Equal(t, user.PasswordHash, "")
}

func TestUser_SetPassword(t *testing.T) {
	user := NewUser()

	assert.Equal(t, user.PasswordHash, "")

	err := user.SetPassword("short")

	assert.Error(t, err, "password length must be greater or equal than 8")

	err = user.SetPassword("good_password")

	assert.NilError(t, err)

	assert.Assert(t, user.PasswordHash != "")
}

func TestUser_IsValid_UID(t *testing.T) {
	user := makeValidUser()
	assert.NilError(t, user.Validate())

	user.UID = ""
	assert.Error(t, user.Validate(), "uid is not set")

	user.UID = "invalid_uid"
	assert.Error(t, user.Validate(), "invalid uid")
}

func TestUser_IsValid_Email(t *testing.T) {
	user := makeValidUser()
	assert.NilError(t, user.Validate())

	user.Email = ""
	assert.Error(t, user.Validate(), "e-mail is required")

	user.Email = "invalid_email"
	assert.Error(t, user.Validate(), "invalid e-mail")
}

func TestUser_IsValid_FirstName(t *testing.T) {
	user := makeValidUser()
	assert.NilError(t, user.Validate())

	user.FirstName = ""
	assert.Error(t, user.Validate(), "first name is required")
}

func TestUser_IsValid_LastName(t *testing.T) {
	user := makeValidUser()
	assert.NilError(t, user.Validate())

	user.LastName = ""
	assert.Error(t, user.Validate(), "last name is required")
}

func TestUser_IsValid_Password(t *testing.T) {
	user := makeValidUser()
	assert.NilError(t, user.Validate())

	user.PasswordHash = ""
	assert.Error(t, user.Validate(), "password is not set")
}

func TestUser_IsValid_Role(t *testing.T) {
	user := makeValidUser()
	assert.NilError(t, user.Validate())

	user.Role = ""
	assert.Error(t, user.Validate(), "role is required")

	user.Role = "invalid_role"
	assert.Error(t, user.Validate(), "invalid role")
}
