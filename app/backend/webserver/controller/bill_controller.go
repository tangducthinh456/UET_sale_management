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

func HandleGETBills(c *gin.Context){
	pageSize, pageToken, filter, er := ParseQueryGET(c)
	if er != nil{
		return
	}
	bills, er := dao.GetDAO().GetBillsByFilter(c, pageSize, pageToken, filter)
	if er != nil{
		ResponseError(c, er ,http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, &forms.BillGetResponse{Bills: bills})
	return
}

func HandlePOSTBills(c *gin.Context){
	var bill *model.Bill
	err := c.Bind(&bill)
	if err != nil{
		ResponseError(c, err, http.StatusBadRequest)
		log.Print("bind json error")
		return
	}

	err = dao.GetDAO().CreateBill(c, bill)
	if err != nil{
		ResponseError(c, err, http.StatusInternalServerError)
		log.Print("create in dao error")
		return
	}
	c.JSON(http.StatusOK, &forms.BillPostResponse{ID: bill.BillID})
}

func HandlePUTBill(c *gin.Context){
	var idStr = c.Param("bill")
	idInt, er := strconv.Atoi(idStr)
	if er != nil{
		ResponseError(c, er, http.StatusBadRequest)
		return
	}
	var bill *model.Bill
	err := c.Bind(&bill)
	if err != nil{
		ResponseError(c, err, http.StatusBadRequest)
		log.Print("bind json error")
		return
	}

	err = dao.GetDAO().UpdateBill(c, uint(idInt), bill)
	if err != nil{
		ResponseError(c, err, http.StatusInternalServerError)
		log.Print("update bill in dao error")
		return
	}
	c.JSON(http.StatusOK, &forms.BillPostResponse{ID: bill.BillID})
}

func HandleDeleteBill(c *gin.Context){
	var idStr = c.Param("bill")
	idInt, er := strconv.Atoi(idStr)
	if er != nil{
		ResponseError(c, er, http.StatusBadRequest)
		return
	}
	er = dao.GetDAO().DeleteBill(c, uint(idInt))
	if er != nil{
		ResponseError(c, er, http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, "Success")
}



