package services

import (
	"github.com/davidalvarezcastro/bookstore-users-api/models/users"
	"github.com/davidalvarezcastro/bookstore-users-api/utils/crypto"
	"github.com/davidalvarezcastro/bookstore-users-api/utils/date"
	errorsutils "github.com/davidalvarezcastro/bookstore-utils-go/rest_errors"
)

var (
	// UserService is the user service with all the functions defined in userServiceInterface
	UserService userServiceInterface = &userService{}
)

type userService struct {
}

type userServiceInterface interface {
	Get(userID int64) (*users.User, *errorsutils.RestErr)
	Create(user users.User) (*users.User, *errorsutils.RestErr)
	Update(isPartial bool, user users.User) (*users.User, *errorsutils.RestErr)
	Delete(userID int64) *errorsutils.RestErr
	Search(status string) (users.Users, *errorsutils.RestErr)
	LoginUser(users.LoginRequest) (*users.User, *errorsutils.RestErr)
}

// Get service to return an user
func (s *userService) Get(userID int64) (*users.User, *errorsutils.RestErr) {
	result := users.User{ID: userID}
	if err := result.Get(); err != nil {
		return nil, err
	}

	return &result, nil
}

// Create service to create an user
func (s *userService) Create(user users.User) (*users.User, *errorsutils.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}

	user.Status = users.StatusActive
	user.Password = crypto.GetMd5(user.Password)
	user.DateCreated = date.GetNowDBFormat()

	if err := user.Save(); err != nil {
		return nil, err
	}

	return &user, nil
}

// Update service to update user info
func (s *userService) Update(isPartial bool, user users.User) (*users.User, *errorsutils.RestErr) {
	current, err := s.Get(user.ID)
	if err != nil {
		return nil, err
	}

	if isPartial {
		if user.FirstName != "" {
			current.FirstName = user.FirstName
		}

		if user.LastName != "" {
			current.LastName = user.LastName
		}

		if user.Email != "" {
			current.Email = user.Email
		}
	} else {
		current.FirstName = user.FirstName
		current.LastName = user.LastName
		current.Email = user.Email
	}

	if err := current.Update(); err != nil {
		return nil, err
	}

	return current, nil
}

// Delete service to remove an user
func (s *userService) Delete(userID int64) *errorsutils.RestErr {
	user := users.User{ID: userID}
	if err := user.Get(); err != nil {
		return err
	}

	return user.Delete()
}

// Search service to search users
func (s *userService) Search(status string) (users.Users, *errorsutils.RestErr) {
	dao := users.User{}
	return dao.FindByStatus(status)
}

// LoginUser service to login a user
func (s *userService) LoginUser(request users.LoginRequest) (*users.User, *errorsutils.RestErr) {
	dao := &users.User{
		Email:    request.Email,
		Password: crypto.GetMd5(request.Password),
	}

	if err := dao.FindByEmailAndPassword(); err != nil {
		return nil, err
	}

	return dao, nil
}
