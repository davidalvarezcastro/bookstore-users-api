package users

import (
	"strings"

	errorsutils "github.com/davidalvarezcastro/bookstore-utils-go/rest_errors"
)

const (
	// StatusActive defaults value for User.Status field
	StatusActive = "active"
)

// User stores user database info
type User struct {
	ID          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
	Status      string `json:"status"`
	Password    string `json:"password"`
}

type Users []User

// Validate validates a given user
func (user *User) Validate() *errorsutils.RestErr {
	user.FirstName = strings.TrimSpace(user.FirstName)
	user.LastName = strings.TrimSpace(user.LastName)

	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email == "" {
		return errorsutils.NewBadRequestError("invalid email address")
	}

	user.Password = strings.TrimSpace(user.Password)
	if user.Password == "" {
		return errorsutils.NewBadRequestError("invalid password")
	}
	return nil
}
