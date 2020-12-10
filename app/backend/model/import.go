package model

import (
	"gorm.io/gorm"
	"time"
)

type Import struct {
	ImportID      uint           `json:"import_id" gorm:"primaryKey; autoIncrement:true"`
	Note          string         `json:"note"`
	ProviderID    uint           `json:"provider_id"`
	Provider      Provider       `json:"provider" gorm:"foreignKey:ProviderID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	CreatedUserID uint           `json:"created_user_id"`
	CreatedUser   User           `json:"created_user" gorm:"foreignKey:CreatedUserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Details       []*ImportLine  `json:"details" gorm:"foreignKey:ImportID"`
	CreatedAt     time.Time      `json:"created_time"`
	Deleted       gorm.DeletedAt `json:"-"`
}

type ImportLine struct {
	//LineID    uint    `gorm:"primaryKey; autoIncrement:true"`
	Import    Import  `gorm:"foreignKey:ImportID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"-"`
	Product   Product `gorm:"foreignKey:ProductID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"-"`
	ImportID  uint    `gorm:"primaryKey;" json:"-"`
	ProductID uint    `gorm:"primaryKey;" json:"product_id"`
	Quantity  int
	Note      string
	Deleted   gorm.DeletedAt `json:"-"`
}
