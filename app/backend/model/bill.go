package model

import (
	"gorm.io/gorm"
	"time"
)

type Bill struct {
	BillID        uint           `json:"bill_id" gorm:"primaryKey; autoIncrement:true"`
	CustomerID    uint           `json:"customer_id"`
	Customer      Customer       `gorm:"foreignKey:CustomerID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"customer"`
	Note          string         `json:"note"`
	CreatedUserID uint           `json:"created_user_id"`
	CreatedUser   User           `gorm:"foreignKey:CreatedUserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"created_user"`
	Details       []*BillLine    `json:"details" gorm:"foreignKey:BillID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	CreatedAt     time.Time      `json:"created_at"`
	Deleted       gorm.DeletedAt `json:"-"`
}

type BillLine struct {
	Bill      Bill    `gorm:"foreignKey:BillID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"-"`
	Product   Product `gorm:"foreignKey:ProductID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"product"`
	BillID    uint    `gorm:"primaryKey;" json:"-"`
	ProductID uint    `gorm:"primaryKey;" json:"product_id"`
	Quantity  int     `json:"quantity"`
	Note      string  `json:"note"`
	Deleted       gorm.DeletedAt `json:"-"`
}
