package model

import "time"

type Product struct {
	ProductID     uint   `gorm:" primaryKey;autoIncrement:true" json:"product_id"`
	ProductName   string `json:"product_name"`
	Cost          uint32 `json:"cost"`
	Price         uint32 `json:"price"`
	Quantity      uint32 `json:"quantity"`
	Brand         string `json:"brand"`
	Position      string `json:"position"`
	IsOnSale      bool   `json:"is_on_sale"`
	GroupID       uint   `json:"-"`
	CreatedUserID uint   `json:"-"`
	CreatedUser   User   `gorm:"foreignKey:CreatedUserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"created_user"`
	Group         Group  `gorm:"foreignKey:GroupID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"group"`
	Description   string `json:"description"`
	CreatedAt     time.Time
	//CreatedUser User
}

type Group struct {
	GroupID   uint       `json:"group_id" gorm:"primaryKey; autoIncrement:true"`
	GroupName string     `json:"group_name"`
	Products  []*Product `gorm:"foreignKey:GroupID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"products"`
}
