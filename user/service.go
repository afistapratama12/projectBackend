package user

import (
	"errors"
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	GetByID(userID string) (User, error)
	GetByEmail(email string) (User, error)
	Register(input UserRegister, uuid string, avatarPath string) (User, error)
	GetByUsername(username string) (User, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetByID(userID string) (User, error) {
	user, err := s.repository.FindByID(userID)

	if err != nil {
		return user, err
	}

	if user.ID == "" || len(user.ID) == 0 {
		return user, errors.New("user id not found")
	}

	return user, nil
}
func (s *service) GetByEmail(email string) (User, error) {
	user, err := s.repository.FindByEmail(email)

	if err != nil {
		return user, err
	}

	if user.ID == "" || len(user.ID) == 0 {
		return user, errors.New("user id not found")
	}

	return user, nil
}

func (s *service) GetByUsername(username string) (User, error) {
	fmt.Println("masuk service getbyusername")
	user, err := s.repository.FindByUsername(username)

	if err != nil {
		return user, err
	}

	if user.ID == "" || len(user.ID) == 0 {
		return user, errors.New("user id not found")
	}

	return user, nil
}

// func (s *service) Login(input UserLogin) (User, error) {
// 	var checkUser User
// 	var err error

// 	if input.Email == "" && len(input.Username) > 0 {
// 		checkUser, err = s.repository.FindByUsername(input.Username)

// 		if err != nil {
// 			return User{}, err
// 		}
// 	}

// 	if len(input.Email) > 0 {
// 		checkUser, err = s.repository.FindByEmail(input.Email)
// 		if err != nil {
// 			return User{}, err
// 		}
// 	}

// 	if checkUser.ID == "" || len(checkUser.ID) <= 1 {
// 		return User{}, errors.New("username / email and password invalid")
// 	}

// 	s.authService.ValidateToken()

// 	// diambahi code lagi
// }

func (s *service) Register(input UserRegister, uuid string, avatarPath string) (User, error) {

	checkEmailUser, err := s.repository.FindByEmail(input.Email)

	if checkEmailUser.ID != "" || len(checkEmailUser.ID) > 1 {
		return User{}, errors.New("email has been registered")
	}

	checkUsernameUser, err := s.repository.FindByEmail(input.Username)

	if checkUsernameUser.ID != "" || len(checkUsernameUser.ID) > 1 {
		return User{}, errors.New("username has been registered")
	}

	if err != nil {
		return User{}, err
	}

	genPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)

	var newUser = User{
		ID:        uuid,
		FirstName: input.FirstName,
		LastName:  input.LastName,
		Photo:     avatarPath,
		Username:  input.Username,
		Email:     input.Email,
		Password:  string(genPassword),
		Role:      "user",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	user, err := s.repository.Create(newUser)

	if err != nil {
		return user, err
	}

	return user, nil
}
