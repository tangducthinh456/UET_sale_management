package forms

import "SaleManagement/model"

type ProductGetResponse struct{
	Products []*model.Product `json:"products"`
}

type ProductPostResponse struct{
	ID uint `json:"product_id"`
}

