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

func HandleGETImports(c *gin.Context){
	pageSize, pageToken, filter, er := ParseQueryGET(c)
	if er != nil{
		return
	}
	imps, er := dao.GetDAO().GetImportsByFilter(c, pageSize, pageToken, filter)
	if er != nil{
		ResponseError(c, er ,http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, &forms.ImportGetResponse{Imports: imps})
	return
}

func HandlePOSTImports(c *gin.Context){
	var imp *model.Import
	err := c.Bind(&imp)
	if err != nil{
		ResponseError(c, err, http.StatusBadRequest)
		log.Print("bind json error")
		return
	}

	err = dao.GetDAO().CreateImport(c, imp)
	if err != nil{
		ResponseError(c, err, http.StatusInternalServerError)
		log.Print("create in dao error")
		return
	}
	c.JSON(http.StatusOK, &forms.ImportPostResponse{ID: imp.ImportID})
}

func HandlePUTImport(c *gin.Context){
	var idStr = c.Param("imp")
	idInt, er := strconv.Atoi(idStr)
	if er != nil{
		ResponseError(c, er, http.StatusBadRequest)
		return
	}
	var imp *model.Import
	err := c.Bind(&imp)
	if err != nil{
		ResponseError(c, err, http.StatusBadRequest)
		log.Print("bind json error")
		return
	}

	err = dao.GetDAO().UpdateImport(c, uint(idInt), imp)
	if err != nil{
		ResponseError(c, err, http.StatusInternalServerError)
		log.Print("update imp in dao error")
		return
	}
	c.JSON(http.StatusOK, &forms.ImportPostResponse{ID: imp.ImportID})
}

func HandleDeleteImport(c *gin.Context){
	var idStr = c.Param("imp")
	idInt, er := strconv.Atoi(idStr)
	if er != nil{
		ResponseError(c, er, http.StatusBadRequest)
		return
	}
	er = dao.GetDAO().DeleteImport(c, uint(idInt))
	if er != nil{
		ResponseError(c, er, http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, "Success")
}


