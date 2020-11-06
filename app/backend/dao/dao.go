package dao

import (
	"SaleManagement/config"
	"SaleManagement/model"
	"context"
	"fmt"
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
)

var gDB *gorm.DB

type dao interface{
    InitDatabaseConnection() error

    MigrateDataModel() error

	GetBillByFilter(ctx context.Context, pageSize int, pageToken int, filter map[string]string) ([]*model.Bill, error)

	GetImportByFilter(ctx context.Context, pageSize int, pageToken int, filter map[string]string) ([]*model.Import, error)
}

func (c *DAO) InitDatabaseConnection()(err error){
	conn := config.GetDBConfig()
	dsn := fmt.Sprintf("user=%s password=%s dbname=%s port=%s host=%s sslmode=disable",
		conn.Username, conn.Password, conn.DBName,
		conn.Port, conn.Host)
	gDB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	return
}

func (c *DAO) MigrateDataModel() error{
	if !gDB.Migrator().HasTable(&model.User{}){
		er := gDB.Migrator().CreateTable(&model.User{})
		if er != nil{
			return er
		}
	}
	if !gDB.Migrator().HasTable(&model.Provider{}){
		er := gDB.Migrator().CreateTable(&model.Provider{})
		if er != nil{
			return er
		}
	}
	if !gDB.Migrator().HasTable(&model.Customer{}){
		er := gDB.Migrator().CreateTable(&model.Customer{})
		if er != nil{
			return er
		}
	}
	if !gDB.Migrator().HasTable(&model.Group{}){
		er := gDB.Migrator().CreateTable(&model.Group{})
		if er != nil{
			return er
		}
	}
	if !gDB.Migrator().HasTable(&model.Product{}){
		er := gDB.Migrator().CreateTable(&model.Product{})
		if er != nil{
			return er
		}
	}
	if !gDB.Migrator().HasTable(&model.Bill{}){
		er := gDB.Migrator().CreateTable(&model.Bill{})
		if er != nil{
			return er
		}
	}
	if !gDB.Migrator().HasTable(&model.BillLine{}){
		er := gDB.Migrator().CreateTable(&model.BillLine{})
		if er != nil{
			return er
		}
	}
	if !gDB.Migrator().HasTable(&model.Import{}){
		er := gDB.Migrator().CreateTable(&model.Import{})
		if er != nil{
			return er
		}
	}
	if !gDB.Migrator().HasTable(&model.ImportLine{}){
		er := gDB.Migrator().CreateTable(&model.ImportLine{})
		if er != nil{
			return er
		}
	}
	return nil
}

func (c* DAO) GetBillByFilter(ctx context.Context, pageSize int, pageToken int, filter map[string]string) ([]*model.Bill, error){
	var bil = []*model.Bill{}
	thisDB := gDB.WithContext(ctx).Table("customers")
	for k, v := range filter{
		thisDB = thisDB.Where(fmt.Sprintf("%s = ?", k), v)
	}
	offset := pageToken - 1
	thisDB = thisDB.Offset(offset).Limit(pageSize)
	er := thisDB.Find(&bil).Error
	if er != nil{
		return nil, er
	}

	for i, v := range bil{
		var lin = []*model.BillLine{}
		er = gDB.WithContext(ctx).Table("bill_lines").Where("bill = ?", v.BillId).Find(&lin).Error
		if er != nil{
			return nil, er
		}
		bil[i].Details = lin
	}
	return bil, nil
}

func (c* DAO) GetImportByFilter(ctx context.Context, pageSize int, pageToken int, filter map[string]string) ([]*model.Import, error){
	var imp = []*model.Import{}
	thisDB := gDB.WithContext(ctx).Table("customers")
	for k, v := range filter{
		thisDB = thisDB.Where(fmt.Sprintf("%s = ?", k), v)
	}
	offset := pageToken - 1
	thisDB = thisDB.Offset(offset).Limit(pageSize)
	er := thisDB.Find(&imp).Error
	if er != nil{
		return nil, er
	}

	for i, v := range imp{
		var lin = []*model.ImportLine{}
		er = gDB.WithContext(ctx).Table("import_lines").Where("import = ?", v.ImportId).Find(&lin).Error
		if er != nil{
			return nil, er
		}
		imp[i].Details = lin
	}
	return imp, nil
}

type DAO struct{}

var d dao = &DAO{}

func GetDAO() dao{
	return d
}