package model

type User struct {
	Username    string `json:"username" gorm:"primaryKey;"`
	Password    string `json:"password"`
	Role        uint   `json:"role"`
	PhoneNumber string `json:"phone_number"`
	Email       string `json:"email"`
}
