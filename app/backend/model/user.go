package model

import (
	//"SaleManagement/dao"
	//"SaleManagement/webserver/forms"
	//"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"time"
	//"SaleManagement/webserver/forms"

)



type User struct {
	UserID        uint   `json:"id" gorm:"primaryKey; autoIncrement:true"`
	Username      string `json:"username" gorm:"unique; not null"`
	Name          string `json:"name" gorm:"not null"`
	Password      []byte `json:"-" gorm:"not null"`
	Role          string `json:"role" gorm:"not null"`
	PhoneNumber   string `json:"phone_number"`
	Email         string `json:"email"`
	CreatedAt     time.Time `json:"created_at"`
	Deleted       gorm.DeletedAt `json:"-"`
}

//func (u User) Login(form forms.LoginForm) (user User, token Token, err error) {
//	err = dao.GetDAO().G.Error
//
//	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(form.Password))
//
//	if err != nil {
//		return user, token, errors.New("Invalid Password")
//	}
//
//	//Generate JWT token
//	tokenDetails, err := authModel.CreateToken(uint(user.UserID), uint(user.Role))
//	if err != nil {
//		return user, token, errors.New("Can not create token")
//	}
//
//	//If save to redis success return token to the user
//	saveErr := authModel.CreateAuth(uint(user.UserID), tokenDetails)
//	if saveErr == nil {
//		token.AccessToken = tokenDetails.AccessToken
//		token.RefreshToken = tokenDetails.RefreshToken
//	}
//
//	return user, token, nil
//}
