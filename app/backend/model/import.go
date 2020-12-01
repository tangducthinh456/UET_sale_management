package model

import "time"

type Import struct {
	ImportID      uint          `json:"import_id" gorm:"primaryKey; autoIncrement:true"`
	Note          string        `json:"note"`
	ProviderID    uint          `json:"-"`
	Provider      Provider      `json:"provider" gorm:"foreignKey:ProviderID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	CreatedUserID uint          `json:"-"`
	CreatedUser   User          `json:"created_user" gorm:"foreignKey:CreatedUserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Details       []*ImportLine `json:"details" gorm:"foreignKey:ImportID"`
	CreatedAt     time.Time     `json:"created_time"`
}

type ImportLine struct {
	//LineID    uint    `gorm:"primaryKey; autoIncrement:true"`
	Import    Import  `gorm:"foreignKey:ImportID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"-"`
	Product   Product `gorm:"foreignKey:ProductID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"-"`
	ImportID  uint    `gorm:"primaryKey;"`
	ProductID uint    `gorm:"primaryKey;"`
	Quantity  int
	Note      string
}
