package model

import "time"

type Product struct {
	ProductId   int    `gorm:" primaryKey;autoIncrement:true" json:"product_id"`
	ProductName string `json:"product_name"`
	Cost        uint32 `json:"cost"`
	Price       uint32 `json:"price"`
	Quantity    uint32 `json:"quantity"`
	Brand       string `json:"brand"`
	Position    string `json:"position"`
	Status      bool   `json:"status"`
	GroupId     uint8  `json:"group_id" gorm:"column:group"`
	Group       Group  `gorm:"foreignKey:group;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"-"`
	Description string `json:"description"`
	CreatedTime time.Time
}

type Group struct {
	GroupId   int    `json:"group_id" gorm:"primaryKey; autoIncrement:true"`
	GroupName string `json:"group_name"`
}
