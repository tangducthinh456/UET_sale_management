package controller

import (
	"SaleManagement/dao"
	"SaleManagement/model"
	"SaleManagement/webserver/forms"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)


func HandleGETProviders(c *gin.Context){
	pageSize, pageToken, filter, er := ParseQueryGET(c)
	if er != nil{
		return
	}
	providers, er := dao.GetDAO().GetProvidersByFilter(c, pageSize, pageToken, filter)
	if er != nil{
		ResponseError(c, er ,http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, &forms.ProviderGetResponse{Providers: providers})
	return
}

func HandlePOSTProviders(c *gin.Context){
	var provider *model.Provider
	err := c.Bind(&provider)
	if err != nil{
		ResponseError(c, err, http.StatusBadRequest)
		log.Print("bind json error")
		return
	}

	err = dao.GetDAO().CreateProvider(c, provider)
	if err != nil{
		ResponseError(c, err, http.StatusInternalServerError)
		log.Print("create in dao error")
		return
	}
	c.JSON(http.StatusOK, &forms.ProviderPostResponse{ID: provider.ProviderID})
}

func HandlePUTProvider(c *gin.Context){
	var idStr = c.Param("provider")
	idInt, er := strconv.Atoi(idStr)
	if er != nil{
		ResponseError(c, er, http.StatusBadRequest)
		return
	}
	var provider *model.Provider
	err := c.Bind(&provider)
	if err != nil{
		ResponseError(c, err, http.StatusBadRequest)
		log.Print("bind json error")
		return
	}

	err = dao.GetDAO().UpdateProvider(c, uint(idInt), provider)
	if err != nil{
		ResponseError(c, err, http.StatusInternalServerError)
		log.Print("update provider in dao error")
		return
	}
	c.JSON(http.StatusOK, &forms.ProviderPostResponse{ID: provider.ProviderID})
}

func HandleDisableProvider(c *gin.Context){
	var idStr = c.Param("provider")
	idInt, er := strconv.Atoi(idStr)
	if er != nil{
		ResponseError(c, er, http.StatusBadRequest)
		return
	}
	er = dao.GetDAO().DisableProvider(c, uint(idInt))
	if er != nil{
		ResponseError(c, er, http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, "Success")
}

func HandleEnableProvider(c *gin.Context){
	var idStr = c.Param("provider")
	idInt, er := strconv.Atoi(idStr)
	if er != nil{
		ResponseError(c, er, http.StatusBadRequest)
		return
	}
	er = dao.GetDAO().EnableProvider(c, uint(idInt))
	if er != nil{
		ResponseError(c, er, http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, "Success")
}
