package forms

import "SaleManagement/model"

type BillGetResponse struct{
	Bills []*model.Bill `json:"bills"`
}

type BillPostResponse struct{
	ID uint `json:"bill_id"`
}

