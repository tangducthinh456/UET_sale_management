package model

type Provider struct {
	ProviderID   uint   `json:"provider_id" gorm:"primaryKey;autoIncrement:true"`
	ProviderName string `json:"provider_name"`
	PhoneNumber  string `json:"phone_number"`
	Address      string `json:"address"`
	Email        string `json:"email"`
}
