package users

import (
	"fmt"

	"github.com/davidalvarezcastro/bookstore-users-api/utils/date"
	"github.com/davidalvarezcastro/bookstore-users-api/utils/errors"
)

var (
	userDB = make(map[int64]*User)
)

// Get returns the user by given userId
func (u *User) Get() *errors.RestErr {
	result := userDB[u.ID]
	if result == nil {
		return errors.NewNotFoundtError(fmt.Sprintf("user %d not found", u.ID))
	}

	u.ID = result.ID
	u.FirstName = result.FirstName
	u.LastName = result.LastName
	u.Email = result.Email
	u.DateCreated = result.LastName

	return nil
}

// Save saves a user in the database
func (u *User) Save() *errors.RestErr {
	current := userDB[u.ID]
	if current != nil {
		if current.Email == u.Email {
			return errors.NewBadRequestError(fmt.Sprintf("email %s already registered", u.Email))
		}

		return errors.NewBadRequestError(fmt.Sprintf("user %d already exists", u.ID))
	}

	u.DateCreated = date.GetNowString()

	userDB[u.ID] = u
	return nil
}
