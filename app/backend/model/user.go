package model

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	UserID        uint   `json:"id" gorm:"primaryKey; autoIncrement:true"`
	Username      string `json:"username" gorm:"unique; not null"`
	Name          string `json:"name" gorm:"not null"`
	Password      []byte `json:"-" gorm:"not null"`
	Role          string `json:"role" gorm:"not null"`
	PhoneNumber   string `json:"phone_number"`
	Email         string `json:"email"`
	CreatedAt     time.Time `json:"created_at"`
	Deleted       gorm.DeletedAt `json:"-"`
}
