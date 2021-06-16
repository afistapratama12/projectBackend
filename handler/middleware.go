package handler

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/afistapratama12/projectBackend/auth"
	"github.com/afistapratama12/projectBackend/user"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// middleware for user login
func Middleware(userService user.Service, authService auth.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		//set id and role
		authHeader := c.GetHeader("Authorization")

		if !strings.Contains(authHeader, "Bearer") {
			response := gin.H{"error": "unauthorize user"}

			// tampilan middleware jika ok maka lanjut, jika tidak kita stop dan tampilkan ini
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		var tokenString string

		dataBearer := strings.Split(authHeader, " ")
		if len(dataBearer) == 2 {
			tokenString = dataBearer[1]
		}

		token, err := authService.ValidateToken(tokenString)

		if err != nil {
			response := gin.H{"error": "unauthorize user"}

			// tampilan middleware jika ok maka lanjut, jika tidak kita stop dan tampilkan ini
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		claim, ok := token.Claims.(jwt.MapClaims)

		if !ok || !token.Valid {
			response := gin.H{"error": "unauthorize user"}

			// tampilan middleware jika ok maka lanjut, jika tidak kita stop dan tampilkan ini
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		userID := claim["user_id"].(string)

		fmt.Println(userID)

		user, err := userService.GetByID(userID)

		if err != nil {
			response := gin.H{"error": "error in internal server middleware"}

			// tampilan middleware jika ok maka lanjut, jika tidak kita stop dan tampilkan ini
			c.AbortWithStatusJSON(http.StatusInternalServerError, response)
			return
		}

		c.Set("currentUser", user)
	}
}

// middleware for role admin ro for role user
// func Authorization() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		//get id and role
// 	}
// }
