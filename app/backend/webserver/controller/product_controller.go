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

func HandleGETProducts(c *gin.Context){
	pageSize, pageToken, filter, er := ParseQueryGET(c)
	if er != nil{
		return
	}
	products, er := dao.GetDAO().GetProductsByFilter(c, pageSize, pageToken, filter)
	if er != nil{
		ResponseError(c, er ,http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, &forms.ProductGetResponse{Products: products})
	return
}

func HandlePOSTProducts(c *gin.Context){
	var product *model.Product
	err := c.Bind(&product)
	if err != nil{
		ResponseError(c, err, http.StatusBadRequest)
		log.Print("bind json error")
		return
	}

	err = dao.GetDAO().CreateProduct(c, product)
	if err != nil{
		ResponseError(c, err, http.StatusInternalServerError)
		log.Print("create in dao error")
		return
	}
	c.JSON(http.StatusOK, &forms.ProductPostResponse{ID: product.ProductID})
}

func HandlePUTProduct(c *gin.Context){
	var idStr = c.Param("product")
	idInt, er := strconv.Atoi(idStr)
	if er != nil{
		ResponseError(c, er, http.StatusBadRequest)
		return
	}
	var product *model.Product
	err := c.Bind(&product)
	if err != nil{
		ResponseError(c, err, http.StatusBadRequest)
		log.Print("bind json error")
		return
	}
	err = dao.GetDAO().UpdateProduct(c, uint(idInt), product)
	if err != nil{
		ResponseError(c, err, http.StatusInternalServerError)
		log.Print("update product in dao error")
		return
	}
	c.JSON(http.StatusOK, &forms.ProductPostResponse{ID: product.ProductID})
}

func HandleDisableProduct(c *gin.Context){
	var idStr = c.Param("product")
	idInt, er := strconv.Atoi(idStr)
	if er != nil{
		ResponseError(c, er, http.StatusBadRequest)
		return
	}
	er = dao.GetDAO().DeleteProduct(c, uint(idInt))
	if er != nil{
		ResponseError(c, er, http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, "Success")
}

//func HandleEnableProduct(c *gin.Context){
//	var idStr = c.Param("product")
//	idInt, er := strconv.Atoi(idStr)
//	if er != nil{
//		ResponseError(c, er, http.StatusBadRequest)
//		return
//	}
//	er = dao.GetDAO().EnableProduct(c, uint(idInt))
//	if er != nil{
//		ResponseError(c, er, http.StatusInternalServerError)
//		return
//	}
//	c.JSON(http.StatusOK, "Success")
//}

func HandleGetGroups(c *gin.Context){
	group, err := dao.GetDAO().GetGroups(c)
	if err != nil{
		ResponseError(c, err, http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, group)
}

func HandlePostGroup(c *gin.Context){
	var group *model.Group
	err := c.Bind(&group)
	if err != nil{
		ResponseError(c, err, http.StatusInternalServerError)
	}
	err = dao.GetDAO().CreateGroup(c, group)
	if err != nil{
		ResponseError(c, err, http.StatusInternalServerError)
	}
	c.JSON(http.StatusOK, "Success")
}



