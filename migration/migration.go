package migration

import (
	"time"
)

type User struct {
	ID            string `gorm:"primaryKey"`
	FirstName     string
	LastName      string
	Photo         string // path profile
	Username      string
	Email         string
	VerifiedEmail bool
	Password      string
	Role          string
	CreatedAt     time.Time
	UpdatedAt     time.Time
	Notes         []Note `gorm:"foreignKey:UserID"`
}

type Note struct {
	ID        string `gorm:"primaryKey"`
	UserID    string `gorm:"index"`
	Title     string
	Body      string
	Secret    string
	Type      string
	Deleted   bool
	CreatedAt time.Time
	UpdatedAt time.Time
}
