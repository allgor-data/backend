// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package generated

import (
	"time"
)

type User struct {
	UID       string    `json:"uid"`
	Email     string    `json:"email"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
