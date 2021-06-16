package note

import "time"

type Note struct {
	ID        string `gorm:"primaryKey"`
	UserID    string
	Title     string
	Body      string
	Secret    string
	Type      string
	CreatedAt time.Time
	UpdatedAt time.Time
	Deleted   bool
}
