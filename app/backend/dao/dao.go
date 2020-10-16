package dao

import (
	"SaleManagement/config"
	"context"
	"fmt"
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
)

var gDB *gorm.DB

type dao interface{
    InitDatabaseConnection(ctx context.Context, id uint32) error
}

func (c *DAO) InitDatabaseConnection(ctx context.Context, id uint32)(err error){
	conn := config.GetDBConfig()
	dsn := fmt.Sprintf("user=%s password=%s dbname=%s port=%s host=%s sslmode=disable",
		conn.Username, conn.Password, conn.DBName,
		conn.Port, conn.Host)
	gDB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	return
}

type DAO struct{}

var d dao = &DAO{}

func (c *DAO) GetDAO() dao{
	return d
}