package model

import "time"

type Import struct {
	ImportId          int           `json:"import_id" gorm:"primaryKey; autoIncrement:true"`
	CreatedTime       time.Time     `json:"created_time"`
	Note              string        `json:"note"`
	ProviderId        string        `json:"provider_id" gorm:"column:provider"`
	Provider          Provider      `json:"-" gorm:"foreignKey:provider;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	ProviderDetail    Customer      `json:"provider_detail" gorm:"-"`
	CreatedUser       int           `json:"created_user"`
	CreatedUserDetail User          `json:"created_user_detail" gorm:"-"`
	User              User          `gorm:"foreignKey:created_user;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"-"`
	Details           []*ImportLine `json:"details" gorm:"-"`
}

type ImportLine struct {
	Import    Import  `gorm:"foreignKey:import;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"-"`
	Product   Product `gorm:"foreignKey:product;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"-"`
	ImportID  string  `gorm:"primaryKey; column:import"`
	ProductId string  `gorm:"primaryKey; column:product"`
	Quantity  int
	Note      string
}
