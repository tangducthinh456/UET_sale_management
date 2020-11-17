package model

import "time"

type Customer struct {
	CustomerID   uint      `json:"customer_id" gorm:"primaryKey;autoIncrement:true;"`
	CustomerName string    `json:"customer_name"`
	DOB          time.Time `json:"dob"`
	PhoneNumber  string    `json:"phone_number"`
	Address      string    `json:"address"`
}
