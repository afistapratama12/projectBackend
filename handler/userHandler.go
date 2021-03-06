package handler

import (
	"fmt"
	"net/http"

	"github.com/afistapratama12/projectBackend/auth"
	"github.com/afistapratama12/projectBackend/helper"
	"github.com/afistapratama12/projectBackend/user"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type userHandler struct {
	service     user.Service
	authService auth.Service
}

func NewUserHandler(service user.Service, authService auth.Service) *userHandler {
	return &userHandler{service, authService}
}

func (h *userHandler) LoginUser(c *gin.Context) {
	var input user.UserLogin
	var checkUser user.User
	var err error

	if err = c.ShouldBindJSON(&input); err != nil {
		errResponse := gin.H{"error": err.Error()}
		c.JSON(400, errResponse)
		return
	}

	if len(input.Username) > 0 || input.Email == "" {
		// fmt.Println("masuk sini error")
		checkUser, err = h.service.GetByUsername(input.Username)
	}

	if len(input.Email) > 0 || input.Username == "" {
		checkUser, err = h.service.GetByEmail(input.Email)
	}

	if err != nil {
		errResponse := gin.H{"error": "username / email and password invalid"}
		c.JSON(400, errResponse)
		return
	}

	if checkUser.ID == "" || len(checkUser.ID) <= 1 {
		if err != nil {
			errResponse := gin.H{"error": "username / email and password invalid"}
			c.JSON(400, errResponse)
			return
		}
		return
	}

	if checkUser.VerifiedEmail == false {
		errResponse := gin.H{"error": "please verified email for login"}
		c.JSON(401, errResponse)
		return
	}

	token, err := h.authService.GenerateToken(checkUser.ID)

	if err != nil {
		errResponse := gin.H{"error": err.Error()}
		c.JSON(400, errResponse)
		return
	}

	response := helper.APIUserResponse(checkUser, token)

	c.JSON(200, response)
}

func (h *userHandler) RegisterUser(c *gin.Context) {
	var input user.UserRegister

	c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, 10*1024*1024)

	if err := c.ShouldBind(&input); err != nil {
		errResponse := gin.H{"error": err.Error()}
		c.JSON(400, errResponse)
		return
	}

	generateUUID := uuid.New()

	avatar, err := c.FormFile("photo")

	if err != nil {
		errResponse := gin.H{"error": err.Error()}
		c.JSON(400, errResponse)
		return
	}

	path := fmt.Sprintf("images/avatar-%s=%s", generateUUID.String(), avatar.Filename)

	user, err := h.service.Register(input, generateUUID.String(), path)
	if err != nil {
		errResponse := gin.H{"error": err.Error()}
		c.JSON(500, errResponse)
		return
	}

	token, err := h.authService.GenerateToken(user.ID)

	go h.service.SendEMailConfirmation(user.Email, token)

	response := helper.ResponseRegister(user)

	err = c.SaveUploadedFile(avatar, path)
	if err != nil {
		errResponse := gin.H{"error": err.Error()}
		c.JSON(500, errResponse)
		return
	}

	c.JSON(201, response)
}

func (h *userHandler) RegisterAdmin(c *gin.Context) {
	var input user.UserRegister

	c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, 10*1024*1024)

	if err := c.ShouldBind(&input); err != nil {
		errResponse := gin.H{"error": err.Error()}
		c.JSON(400, errResponse)
		return
	}

	generateUUID := uuid.New()

	avatar, err := c.FormFile("photo")

	if err != nil {
		errResponse := gin.H{"error": err.Error()}
		c.JSON(400, errResponse)
		return
	}

	path := fmt.Sprintf("images/avatar-%s=%s", generateUUID.String(), avatar.Filename)

	user, err := h.service.RegisterAdmin(input, generateUUID.String(), path)
	if err != nil {
		errResponse := gin.H{"error": err.Error()}
		c.JSON(500, errResponse)
		return
	}

	token, err := h.authService.GenerateToken(user.ID)

	go h.service.SendEMailConfirmation(user.Email, token)

	response := helper.ResponseRegister(user)

	err = c.SaveUploadedFile(avatar, path)
	if err != nil {
		errResponse := gin.H{"error": err.Error()}
		c.JSON(500, errResponse)
		return
	}

	c.JSON(201, response)
}

func (h *userHandler) VerificationEmailUser(c *gin.Context) {
	confirmationKey := c.Params.ByName("confirmation_key")

	token, err := h.authService.ValidateToken(confirmationKey)

	if err != nil {
		response := gin.H{"error": "confirmation key email not valid"}
		c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	claim, ok := token.Claims.(jwt.MapClaims)

	if !ok || !token.Valid {
		response := gin.H{"error": "confirmation key email not valid"}
		c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	userID := claim["user_id"].(string)

	fmt.Println(userID)

	user, err := h.service.VerifiedEmailByUserID(userID)

	if err != nil {
		errResponse := gin.H{"error": err.Error()}
		c.JSON(500, errResponse)
		return
	}

	response := helper.ResponseVerification(user)

	c.JSON(200, response)
}
