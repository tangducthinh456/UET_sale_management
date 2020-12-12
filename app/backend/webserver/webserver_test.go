package webserver

import (
	"SaleManagement/config"
	"SaleManagement/dao"
	"SaleManagement/dao/redis"
	"SaleManagement/webserver/router"
	"context"
	"testing"
)

func TestMain(m *testing.M) {
	config.Init("../config/config.yaml")
	dao.GetDAO().InitDatabaseConnection()
	redis.RedisClientInit()
	dao.GetDAO().MigrateDataModel(context.Background())
	var route = router.NewRouter()
	route.Run(":8080")
}
