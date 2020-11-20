package model

type Provider struct {
	ProviderID   uint   `json:"provider_id" gorm:"primaryKey;autoIncrement:true"`
	ProviderName string `json:"provider_name" form:"provider_name"`
	PhoneNumber  string `json:"phone_number" form:"phone_number"`
	Address      string `json:"address" form:"address"`
	Email        string `json:"email" form:"email"`
	IsActive     bool   `json:"is_active" form:"is_active"`
}
