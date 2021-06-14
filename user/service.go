package user

import (
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	GetByID(userID string) (User, error)
	GetByEmail(email string) (User, error)
	Register(input UserRegister, uuid string, avatarPath string) (User, error)
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

	if user.UserID == "" || len(user.UserID) == 0 {
		return user, errors.New("user id not found")
	}

	return user, nil
}
func (s *service) GetByEmail(email string) (User, error) {
	user, err := s.repository.FindByEmail(email)

	if err != nil {
		return user, err
	}

	if user.UserID == "" || len(user.UserID) == 0 {
		return user, errors.New("user id not found")
	}

	return user, nil
}
func (s *service) Register(input UserRegister, uuid string, avatarPath string) (User, error) {

	checkEmailUser, err := s.repository.FindByEmail(input.Email)

	if err != nil {
		return User{}, err
	}

	if checkEmailUser.UserID != "" || len(checkEmailUser.UserID) > 1 {
		return User{}, errors.New("email has been registered")
	}

	genPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)

	var newUser = User{
		UserID:    uuid,
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
