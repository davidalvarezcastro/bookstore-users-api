package users

import "encoding/json"

// PublicUser shows (public) information from the user to all the services
type PublicUser struct {
	ID          int64  `json:"id"`
	DateCreated string `json:"date_created"`
	Status      string `json:"status"`
}

// PrivateUser shows (private) information from the user to internal services
type PrivateUser struct {
	ID          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
	Status      string `json:"status"`
}

// Marshall marshalles users
func (users Users) Marshall(isPublic bool) interface{} {
	results := make([]interface{}, len(users))

	for index, user := range users {
		results[index] = user.Marshall(isPublic)
	}

	return results
}

// Marshall marshalles a user depends on the isPublic argument
func (user *User) Marshall(isPublic bool) interface{} {
	if isPublic {
		return PublicUser{
			ID:          user.ID,
			DateCreated: user.DateCreated,
			Status:      user.Status,
		}
	}

	userJson, _ := json.Marshal(user)
	var privateUser PrivateUser
	json.Unmarshal(userJson, &privateUser)
	return privateUser
}
