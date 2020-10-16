package model

type Product struct{
	ProductId string `gorm:" primaryKey;" json:"product_id"`
	Id uint32        `gorm:"autoIncrement:true" json:"product_id"`
	ProductName string `json:"product_name"`
	Cost uint32    `json:"cost"`
	Price uint32   `json:"price"`
	Quantity uint32 `json:"quantity"`
	Brand string    `json:"brand"`
	Position string `json:"position"`
	Status bool     `json:"status"`
	GroupId uint8   `json:"group_id" gorm:"column:group"`
	Group Group     `gorm:"foreignKey:group;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"-"`
	Description string `json:"product_name"`
}

type Group struct{
	GroupId string `json:"group_id"`
	GroupName uint8 `json:"group_name"`
}
