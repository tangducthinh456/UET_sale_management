package model

import "time"

type Product struct {
	ProductID     uint      `gorm:" primaryKey;autoIncrement:true" json:"product_id"`
	ProductName   string    `json:"product_name" form:"product_name"`
	Cost          uint32    `json:"cost" form:"cost"`
	Price         uint32    `json:"price" form:"price"`
	Quantity      uint32    `json:"quantity" form:"quantity"`
	Brand         string    `json:"brand" form:"brand"`
	Position      string    `json:"position" form:"position"`
	IsOnSale      bool      `json:"is_on_sale"`
	GroupID       uint      `json:"-" form:"group_id"`
	CreatedUserID uint      `json:"-" form:"created_user"`
	CreatedUser   User      `gorm:"foreignKey:CreatedUserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"created_user"`
	Group         Group     `gorm:"foreignKey:GroupID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"group"`
	Description   string    `json:"description" form:"description"`
	CreatedAt     time.Time `json:"created_at"`
	//CreatedUser User
}

type Group struct {
	GroupID   uint       `json:"group_id" gorm:"primaryKey; autoIncrement:true"`
	GroupName string     `json:"group_name" form:"group_name"`
	Products  []*Product `gorm:"foreignKey:GroupID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"products"`
}
