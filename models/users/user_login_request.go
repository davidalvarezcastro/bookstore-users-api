package users

// LoginRequest struct stores info needed for loging
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
