package model

import "time"

type Bill struct {
	BillID         uint        `json:"bill_id" gorm:"primaryKey; autoIncrement:true"`
	CustomerID     uint        `json:"-"`
	CustomerDetail Customer    `gorm:"foreignKey:CustomerID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"customer_detail"`
	Note           string      `json:"note"`
	CreatedUserID  uint        `json:"-"`
	CreatedUser    User        `gorm:"foreignKey:CreatedUserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"created_user"`
	Details        []*BillLine `json:"details" gorm:"foreignKey:BillID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	CreatedTime    time.Time   `json:"created_time"`
}

type BillLine struct {
	Bill      Bill    `gorm:"foreignKey:BillID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"-"`
	Product   Product `gorm:"foreignKey:ProductID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"-"`
	BillID    uint    `gorm:"primaryKey;"`
	ProductID uint    `gorm:"primaryKey; "`
	Quantity  int
	Note      string
}
