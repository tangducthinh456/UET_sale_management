package forms

import "SaleManagement/model"

type ImportGetResponse struct{
	Imports []*model.Import `json:"imports"`
}

type ImportPostResponse struct{
	ID uint `json:"import_id"`
}

