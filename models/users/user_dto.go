package users

import (
	"strings"

	"github.com/davidalvarezcastro/bookstore-users-api/utils/errors"
)

// User stores user database info
type User struct {
	ID          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
	// Status      string `json:"status"`
	// Password    string `json:"password"`
}

// Validate validates a given user
func (user *User) Validate() *errors.RestErr {
	user.FirstName = strings.TrimSpace(user.FirstName)
	user.LastName = strings.TrimSpace(user.LastName)

	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email == "" {
		return errors.NewBadRequestError("invalid email address")
	}

	// user.Password = strings.TrimSpace(user.Password)
	// if user.Password == "" {
	// 	return errors.NewBadRequestError("invalid password")
	// }
	return nil
}