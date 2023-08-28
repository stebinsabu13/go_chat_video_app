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

type Connection struct {
	ID              uint `gorm:"primarykey;auto_increment"`
	UserID          uint `gorm:"not null"`
	FriendID        uint `gorm:"not null"`
	RequestAccepted bool `gorm:"default:false"`
}

type Status struct {
	ID     uint   `gorm:"primarykey;auto_increment"`
	Status string `gorm:"not null"`
}

type UserStatus struct {
	ID         uint `gorm:"primarykey;auto_increment"`
	UserID     uint `gorm:"not null"`
	UserStatus uint `gorm:"not null"`
}
