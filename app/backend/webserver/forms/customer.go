package forms

import "SaleManagement/model"

type CustomerGetResponse struct{
	Customers []*model.Customer `json:"customers"`
}

type CustomerPostResponse struct{
	ID uint `json:"customer_id"`
}
