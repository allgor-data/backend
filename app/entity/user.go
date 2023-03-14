package entity

import (
	"errors"
	"fmt"
	"net/mail"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserRole = string

const (
	UserRoleAdmin    UserRole = "admin"
	UserRoleOperator UserRole = "operator"
)

type User struct {
	UID          string
	Email        string
	FirstName    string
	LastName     string
	PasswordHash string
	Role         UserRole
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func NewUser() *User {
	timeNow := time.Now().UTC()

	return &User{
		UID:       uuid.New().String(),
		Role:      UserRoleOperator,
		CreatedAt: timeNow,
		UpdatedAt: timeNow,
	}
}

func (u *User) Validate() error {
	err := u.validateUID()
	if err != nil {
		return err
	}

	err = u.validateEmail()
	if err != nil {
		return err
	}

	err = u.validateFirstName()
	if err != nil {
		return err
	}

	err = u.validateLastName()
	if err != nil {
		return err
	}

	err = u.validatePasswordHash()
	if err != nil {
		return err
	}

	err = u.validateRole()
	if err != nil {
		return err
	}

	return nil
}

func (u *User) SetPassword(password string) error {
	const MIN_LENGTH = 8

	if len(password) < MIN_LENGTH {
		return fmt.Errorf("password length must be greater or equal than %d", MIN_LENGTH)
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return err
	}

	u.PasswordHash = string(hash)

	return nil
}

func (u *User) validateUID() error {
	if u.UID == "" {
		return errors.New("uid is not set")
	}

	_, err := uuid.Parse(u.UID)
	if err != nil {
		return errors.New("invalid uid")
	}

	return err
}

func (u *User) validateEmail() error {
	if u.Email == "" {
		return errors.New("e-mail is required")
	}

	_, err := mail.ParseAddress(u.Email)
	if err != nil {
		return errors.New("invalid e-mail")
	}

	return err
}

func (u *User) validateFirstName() error {
	if u.FirstName == "" {
		return errors.New("first name is required")
	}

	return nil
}

func (u *User) validateLastName() error {
	if u.LastName == "" {
		return errors.New("last name is required")
	}

	return nil
}

func (u *User) validatePasswordHash() error {
	if u.PasswordHash == "" {
		return errors.New("password is not set")
	}

	return nil
}

func (u *User) validateRole() error {
	if u.Role == "" {
		return errors.New("role is required")
	}

	if u.Role != UserRoleAdmin && u.Role != UserRoleOperator {
		return errors.New("invalid role")
	}

	return nil
}
