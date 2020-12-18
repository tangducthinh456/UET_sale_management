package controller

import (
	"SaleManagement/webserver/forms"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

const (
	defaultPage     string = "1"
	defaultPageSize string = "10"
)

var operators = [6]string{">=", "<=", ">", "<", "=", "LIKE"}

// ResponseError return response error
func ResponseError(c *gin.Context, er error, statusCode int) {
	response := &forms.ResponseError{
		Message:    er.Error(),
		StatusCode: statusCode,
	}
	c.JSON(statusCode, response)
	c.Abort()
	return
}

func ParseQueryGET(c *gin.Context)(pageSize int, pageToken int, filterMap map[string][]string, er error){
	query := c.Request.URL.Query()

	pageSizeStr := query.Get("page_size")
	if pageSizeStr == "" {
		pageSizeStr = defaultPageSize
	}
	pageSize, er = strconv.Atoi(pageSizeStr)
	if er != nil {
		ResponseError(c, er, http.StatusBadGateway)
		c.Abort()
		return
	}

	pageTokenStr := query.Get("page_token")
	if pageTokenStr == "" {
		pageTokenStr = defaultPage
	}
	pageToken, er = strconv.Atoi(pageTokenStr)
	if er != nil {
		ResponseError(c, er, http.StatusBadRequest)
		c.Abort()
		return
	}

	filter := query.Get("filter")
	conditions := strings.Split(filter, ";")
	filterMap = make(map[string][]string, 0)

	for _, condition := range conditions {
		var filterPair = make([]string, 0)

		for _, operator := range operators {
			if strings.Contains(condition, operator) {
				filterPair = strings.Split(condition, operator)
				filterMap[filterPair[0]] = []string{operator, filterPair[1]}
				break
			}
		}
	}
	//fmt.Println(filterMap)
	return
}
