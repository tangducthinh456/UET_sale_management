package model

import (
	"gorm.io/gorm"
	"time"
)

type Provider struct {
	ProviderID    uint           `json:"provider_id" gorm:"primaryKey;autoIncrement:true"`
	ProviderName  string         `json:"provider_name" form:"provider_name"`
	PhoneNumber   string         `json:"phone_number" form:"phone_number"`
	Address       string         `json:"address" form:"address"`
	Email         string         `json:"email" form:"email"`
	CreatedUserID uint           `json:"created_user_id" form:"created_user_id"`
	CreatedUser   User           `gorm:"foreignKey:CreatedUserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"created_user"`
	CreatedAt     time.Time      `json:"created_at"`
	Deleted       gorm.DeletedAt `json:"-"`
}
