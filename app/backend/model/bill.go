package model

import "time"

type Bill struct {
	BillId            int         `json:"bill_id" gorm:"primaryKey; autoIncrement:true"`
	CreatedTime       time.Time   `json:"created_time"`
	CustomerId        string      `json:"customer_id" gorm:"column:customer"`
	CustomerDetail    Customer    `json:"customer_detail" gorm:"-"`
	Customer          Customer    `gorm:"foreignKey:customer;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"-"`
	Note              string      `json:"note"`
	CreatedUser       int         `json:"created_user"`
	User              User        `gorm:"foreignKey:created_user;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"-"`
	CreatedUserDetail User        `json:"created_user_detail" gorm:"-"`
	Details           []*BillLine `json:"details" gorm:"-"`
}

type BillLine struct {
	Bill      Bill    `gorm:"foreignKey:bill;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"-"`
	Product   Product `gorm:"foreignKey:product;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"-"`
	BillID    int     `gorm:"primaryKey; column:bill"`
	ProductId string  `gorm:"primaryKey; column:product"`
	Quantity  int
	Note      string
}
