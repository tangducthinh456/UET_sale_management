package model

import (
	"gorm.io/gorm"
	"time"
)

type Customer struct {
	CustomerID    uint           `json:"customer_id" gorm:"primaryKey;autoIncrement:true;"`
	CustomerName  string         `json:"customer_name" form:"name"`
	DOB           time.Time      `json:"-" form:"-" gorm:"type:date"`
	DOBForm       string         `json:"dob" form:"dob" gorm:"-"`
	PhoneNumber   string         `json:"phone_number" form:"phone_number"`
	Email         string         `json:"email" form:"email"`
	Address       string         `json:"address" form:"address"`
	CreatedUserID uint           `json:"created_user_id" form:"created_user_id"`
	CreatedUser   User           `gorm:"foreignKey:CreatedUserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"created_user"`
	CreatedAt     time.Time      `json:"created_at"`
	Deleted       gorm.DeletedAt `json:"-"`
}
