package helper

import "github.com/afistapratama12/projectBackend/user"

type userResponse struct {
	ID        string `json:"id"`
	Username  string `json:"username"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Photo     string `json:"photo"`
	Email     string `json:"email"`
	Token     string `json:"token"`
}

func APIUserResponse(user user.User, token string) *userResponse {
	return &userResponse{
		ID:        user.UserID,
		Username:  user.Username,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Photo:     user.Photo,
		Email:     user.Email,
		Token:     token,
	}
}
