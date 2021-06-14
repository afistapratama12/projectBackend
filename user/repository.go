package user

import (
	"fmt"

	"gorm.io/gorm"
)

type Repository interface {
	Create(user User) (User, error)
	FindByEmail(email string) (User, error)
	FindByUsername(username string) (User, error)
	FindByID(ID string) (User, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Create(user User) (User, error) {
	if err := r.db.Create(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) FindByUsername(username string) (User, error) {
	var user User

	fmt.Println("masuk repository findbyusername")

	if err := r.db.Where("username = ?", username).Find(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (r *repository) FindByEmail(email string) (User, error) {
	var user User

	if err := r.db.Where("email = ?", email).Find(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}
func (r *repository) FindByID(ID string) (User, error) {
	var user User

	if err := r.db.Where("id = ?", ID).Find(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}
