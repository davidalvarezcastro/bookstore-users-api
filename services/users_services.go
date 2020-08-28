package services

import (
	"github.com/davidalvarezcastro/bookstore-users-api/models/users"
	"github.com/davidalvarezcastro/bookstore-users-api/utils/errors"
)

// Get service to return an user
func Get(userID int64) (*users.User, *errors.RestErr) {
	result := users.User{ID: userID}
	if err := result.Get(); err != nil {
		return nil, err
	}

	return &result, nil
}

// CreateUser service to create an user
func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}

	if err := user.Save(); err != nil {
		return nil, err
	}

	return &user, nil
}
