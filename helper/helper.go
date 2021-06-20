package helper

import (
	"github.com/afistapratama12/projectBackend/user"
	"github.com/gin-gonic/gin"
)

type userResponse struct {
	UserID    string `json:"user_id"`
	Username  string `json:"username"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Photo     string `json:"photo"`
	Email     string `json:"email"`
	Token     string `json:"token"`
}

func APIUserResponse(user user.User, token string) *userResponse {
	return &userResponse{
		UserID:    user.ID,
		Username:  user.Username,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Photo:     user.Photo,
		Email:     user.Email,
		Token:     token,
	}
}

func ResponseVerification(user user.User) gin.H {
	return gin.H{
		"username": user.Username,
		"email":    user.Email,
		"message":  "user email verification successfull",
	}
}

func ResponseRegister(user user.User) gin.H {
	return gin.H{
		"username": user.Username,
		"email":    user.Email,
		"message":  "please verified email before login",
	}
}
