package domain

import "time"

type User struct {
	ID           uint      `gorm:"primarykey;auto_increment"`
	CreatedAt    time.Time `gorm:"not null"`
	UpdatedAt    time.Time `gorm:"not null"`
	FirstName    string    `gorm:"not null"`
	LastName     string    `gorm:"not null"`
	Email        string    `gorm:"uniqueIndex;not null"`
	Password     string    `gorm:"not null"`
	MobileNumber string    `gorm:"uniqueIndex;not null"`
}
