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

func HandleGETCustomers(c *gin.Context){
	pageSize, pageToken, filter, er := ParseQueryGET(c)
	if er != nil{
		return
	}
	customers, er := dao.GetDAO().GetCustomersByFilter(c, pageSize, pageToken, filter)
	if er != nil{
		ResponseError(c, er ,http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, &forms.CustomerGetResponse{Customers: customers})
	return
}

func HandlePOSTCustomers(c *gin.Context){
	var customer *model.Customer
	err := c.Bind(&customer)
	if err != nil{
		ResponseError(c, err, http.StatusBadRequest)
		log.Print("bind json error")
		return
	}

	err = dao.GetDAO().CreateCustomer(c, customer)
	if err != nil{
		ResponseError(c, err, http.StatusInternalServerError)
		log.Print("create in dao error")
		return
	}
	c.JSON(http.StatusOK, &forms.CustomerPostResponse{ID: customer.CustomerID})
}

func HandlePUTCustomer(c *gin.Context){
	var idStr = c.Param("customer")
	idInt, er := strconv.Atoi(idStr)
	if er != nil{
		ResponseError(c, er, http.StatusBadRequest)
		return
	}
	var customer *model.Customer
	err := c.Bind(&customer)
	if err != nil{
		ResponseError(c, err, http.StatusBadRequest)
		log.Print("bind json error")
		return
	}

	err = dao.GetDAO().UpdateCustomer(c, uint(idInt), customer)
	if err != nil{
		ResponseError(c, err, http.StatusInternalServerError)
		log.Print("update customer in dao error")
		return
	}
	c.JSON(http.StatusOK, &forms.CustomerPostResponse{ID: customer.CustomerID})
}

func HandleDisableCustomer(c *gin.Context){
	var idStr = c.Param("customer")
	idInt, er := strconv.Atoi(idStr)
	if er != nil{
		ResponseError(c, er, http.StatusBadRequest)
		return
	}
	er = dao.GetDAO().DisableCustomer(c, uint(idInt))
	if er != nil{
		ResponseError(c, er, http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, "Success")
}

func HandleEnableCustomer(c *gin.Context){
	var idStr = c.Param("customer")
	idInt, er := strconv.Atoi(idStr)
	if er != nil{
		ResponseError(c, er, http.StatusBadRequest)
		return
	}
	er = dao.GetDAO().EnableCustomer(c, uint(idInt))
	if er != nil{
		ResponseError(c, er, http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, "Success")
}

