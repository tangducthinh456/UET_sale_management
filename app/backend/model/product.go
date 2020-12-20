package model

import (
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
	"time"
)

type Product struct {
	ProductID     uint            `gorm:" primaryKey;autoIncrement:true" json:"product_id"`
	ProductName   string          `json:"product_name" form:"product_name"`
	Cost          decimal.Decimal `json:"cost" form:"cost" gorm:"type:decimal"`
	Price         decimal.Decimal `json:"price" form:"price" gorm:"type:decimal"`
	Quantity      uint32          `json:"quantity" form:"quantity"`
	Brand         string          `json:"brand" form:"brand"`
	Position      string          `json:"position" form:"position"`
	GroupID       uint            `json:"group_id" form:"group_id"`
	CreatedUserID uint            `json:"created_user_id" form:"created_user_id"`
	CreatedUser   User            `gorm:"foreignKey:CreatedUserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"created_user"`
	Group         Group           `gorm:"foreignKey:GroupID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"group"`
	Description   string          `json:"description" form:"description"`
	CreatedAt     time.Time       `json:"created_at"`
	Deleted       gorm.DeletedAt  `json:"-"`
}

type Group struct {
	GroupID   uint       `json:"group_id" gorm:"primaryKey; autoIncrement:true"`
	GroupName string     `json:"group_name" form:"group_name"`
	Products  []*Product `gorm:"foreignKey:GroupID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"-"`
}
