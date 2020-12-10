package controller

import (
	"SaleManagement/dao"
	"log"
	"strconv"

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
    	log.Print("bind json error")
    	return
	}
	userDB, er := forms.ConvertUserDtoToDB(user)
	if er != nil{
		ResponseError(c, err, http.StatusInternalServerError)
		log.Print("convert user dto to db error")
		return
	}
	err = dao.GetDAO().CreateUser(c, userDB)
	if err != nil{
		ResponseError(c, err, http.StatusInternalServerError)
		log.Print("create user in dao error")
		return
	}
	c.JSON(http.StatusOK, &forms.UserPostResponse{ID: userDB.UserID})
}

func HandlePUTUser(c *gin.Context){
	var idStr = c.Param("user")
	idInt, er := strconv.Atoi(idStr)
	if er != nil{
		ResponseError(c, er, http.StatusBadRequest)
		return
	}
	var user forms.UserDtoRequest
	err := c.Bind(&user)
	if err != nil{
		ResponseError(c, err, http.StatusBadRequest)
		log.Print("bind json error")
		return
	}

	userDB, er := forms.ConvertUserDtoToDB(user)
	if er != nil{
		ResponseError(c, err, http.StatusInternalServerError)
		log.Print("convert user dto to db error")
		return
	}
	err = dao.GetDAO().UpdateUser(c, uint(idInt), userDB)
	if err != nil{
		ResponseError(c, err, http.StatusInternalServerError)
		log.Print("update user in dao error")
		return
	}
	c.JSON(http.StatusOK, &forms.UserPostResponse{ID: userDB.UserID})
}

func HandleDisableUser(c *gin.Context){
	var idStr = c.Param("user")
	idInt, er := strconv.Atoi(idStr)
	if er != nil{
		ResponseError(c, er, http.StatusBadRequest)
		return
	}
	er = dao.GetDAO().DeleteUser(c, uint(idInt))
	if er != nil{
		ResponseError(c, er, http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, "Success")
}

//func HandleEnableUser(c *gin.Context){
//	var idStr = c.Param("user")
//	idInt, er := strconv.Atoi(idStr)
//	if er != nil{
//		ResponseError(c, er, http.StatusBadRequest)
//		return
//	}
//	er = dao.GetDAO().EnableUser(c, uint(idInt))
//	if er != nil{
//		ResponseError(c, er, http.StatusInternalServerError)
//		return
//	}
//	c.JSON(http.StatusOK, "Success")
//}