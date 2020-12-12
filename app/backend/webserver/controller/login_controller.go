package controller

import (
	"SaleManagement/dao"
	"SaleManagement/webserver/forms"
	"errors"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func HandleLogin(c *gin.Context) {
	var login forms.LoginForm
	err := c.Bind(&login)
	if err != nil {
		ResponseError(c, err, http.StatusInternalServerError)
		return
	}
	user, err := dao.GetDAO().GetUserByUsername(c, login.Username)
	if err != nil {
		ResponseError(c, err, http.StatusInternalServerError)
		return
	}
	err = bcrypt.CompareHashAndPassword(user.Password, []byte(login.Password))
	if err != nil {
		ResponseError(c, errors.New("Invalid password"), http.StatusUnauthorized)
		return
	}
	var role uint
	if user.Role == "MANAGER" {
		role = 0
	} else if user.Role == "SALE" {
		role = 1
	} else if user.Role == "EMPLOYEE" {
		role = 2
	}
	tokenDetails, err := authModel.CreateToken(user.UserID, role)
	if err != nil {
		ResponseError(c, err, http.StatusInternalServerError)
		return
	}
    //fmt.Println(tokenDetails)
	//If save to redis success return token to the user
	saveErr := authModel.CreateAuth(user.UserID, tokenDetails)
	if saveErr != nil {
		ResponseError(c, saveErr, http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "login success",
		"role":    user.Role,
		"userID":  user.UserID,
		"token":   tokenDetails,
	})
}
