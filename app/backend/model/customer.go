package model

import "time"

type Customer struct {
	CustomerID   uint      `json:"customer_id" gorm:"primaryKey;autoIncrement:true;"`
	CustomerName string    `json:"customer_name" form:"name"`
	DOB          time.Time `json:"dob" form:"dob"`
	PhoneNumber  string    `json:"phone_number" form:"phone_number"`
	Address      string    `json:"address" form:"address"`
	IsActive     bool      `json:"is_active"`
}
