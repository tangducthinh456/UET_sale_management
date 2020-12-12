package forms

import (
	"SaleManagement/model"
	"golang.org/x/crypto/bcrypt"
)



type UsersGetResponse struct {
	Users []*model.User `json:"users"`
}

type UserPostResponse struct {
	ID uint `json:"id"`
}

type UserDtoRequest struct {
	ID          uint   `json:"id" form:"id"`
	Username    string `json:"username" form:"username"`
	Name        string `json:"name" form:"name"`
	Password    string `json:"password" form:"password"`
	Role        string `json:"role" form:"role"`
	PhoneNumber string `json:"phone_number" form:"phone_number"`
	Email       string `json:"email" form:"email"`
}

func ConvertUserDtoToDB(user UserDtoRequest) (*model.User, error) {
	hash, er := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	if er != nil {
		return nil, er
	}

	return &model.User{
		//UserID:      0,
		Username:    user.Username,
		Name:        user.Name,
		Password:    hash,
		Role:        user.Role,
		PhoneNumber: user.PhoneNumber,
		Email:       user.Email,
	}, nil
}



