package controller

import (
	"SaleManagement/dao"
	//"SaleManagement/model"
	"SaleManagement/webserver/forms"
	//"fmt"

	//"encoding/json"
	//"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func HandleGETUsers(c *gin.Context){
    pageSize, pageToken, filter, er := ParseQueryGET(c)
    if er != nil{
    	return
	}
	users, er := dao.GetDAO().GetUsersByFilter(c, pageSize, pageToken, filter)
	if er != nil{
		ResponseError(c, er ,http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, &forms.UsersGetResponse{Users: users})
	return
}

func HandlePOSTUsers(c *gin.Context){
	var user forms.UserDtoRequest
    err := c.Bind(&user)
    if err != nil{
    	ResponseError(c, err, http.StatusBadRequest)
    	return
	}
	userDB, er := forms.ConvertUserDtoToDB(user)
	if er != nil{
		ResponseError(c, err, http.StatusInternalServerError)
		return
	}
	err = dao.GetDAO().CreateUser(c, userDB)
	if err != nil{
		ResponseError(c, err, http.StatusInternalServerError)
	}
	c.JSON(http.StatusOK, &forms.UserPostResponse{ID: userDB.UserID})
}
