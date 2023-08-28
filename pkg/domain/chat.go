package domain

import "time"

type Conversation struct {
	ID        uint      `gorm:"primarykey;auto_increment"`
	User1     uint      `gorm:"not null"`
	User2     uint      `gorm:"not null"`
	Message   string    `gorm:"not null"`
	CreatedAt time.Time `gorm:"not null"`
}
