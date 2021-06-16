package user

import (
	"time"

	"github.com/afistapratama12/projectBackend/note"
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
	DeletedAt     time.Time   `gorm:"index"`
	Notes         []note.Note `gorm:"foreignKey:UserID"`
}
