package forms

import (
	"SaleManagement/model"
	"golang.org/x/crypto/bcrypt"
	)

type UsersGetResponse struct {
	Users []*model.User `json:"users"`
}

type UserPostResponse struct{
	ID   uint  	`json:"id"`
}

type UserDtoRequest struct {
	Username    string `json:"username"`
	Name        string `json:"name"`
	Password    string `json:"password"`
	Role        string `json:"role"`
	PhoneNumber string `json:"phone_number"`
	Email       string `json:"email"`
}

func ConvertUserDtoToDB(user UserDtoRequest) (*model.User, error) {
	hash, er := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	if er != nil{
		return nil, er
	}

	return &model.User{
		//UserID:      0,
		Username:    user.Username,
		Name:        user.Name,
		Password:    string(hash),
		Role:        user.Role,
		PhoneNumber: user.PhoneNumber,
		Email:       user.Email,
		IsActive:    true,
	}, nil
}
