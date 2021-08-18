package models

import "time"

type User struct {
	ID          uint      `gorm:"primaryKey;autoIncrement"`
	Email       string    `gorm:"unique;not null"`
	DateOfBirth time.Time `gorm:"not null"`
}
