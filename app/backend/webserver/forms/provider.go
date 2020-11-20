package forms

import "SaleManagement/model"

type ProviderGetResponse struct {
	Providers []*model.Provider `json:"providers"`
}

type ProviderPostResponse struct {
	ID uint `json:"id"`
}


