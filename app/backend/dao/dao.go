package dao

import (
	"SaleManagement/config"
	"SaleManagement/model"
	"context"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var gDB *gorm.DB

type dao interface {
	InitDatabaseConnection() error

	MigrateDataModel(ctx context.Context) error

	
    GetUsersByFilter(ctx context.Context, pageSize int, pageToken int, filter map[string][]string) ([]*model.User, error)
	CreateUser(ctx context.Context, user *model.User) error
	UpdateUser(ctx context.Context, id uint, user *model.User) error
	DisableUser(ctx context.Context, id uint) error
	EnableUser(ctx context.Context, id uint) error

	//GetProviderByFilter(ctx context.Context, pageSize int, pageToken int, filter map[string][]string) ([]*model.Provider, error)
	//CreateProvider(ctx context.Context, provider *model.Provider) error
	//UpdateProvider(ctx context.Context, id uint, provider *model.Provider) error
	//DeleteProvider(ctx context.Context, id uint) error
	//
	//GetCustomerByFilter(ctx context.Context, pageSize int, pageToken int, filter map[string][]string) ([]*model.Customer, error)
	//CreateCustomer(ctx context.Context, customer *model.Customer) error
	//UpdateCustomer(ctx context.Context, id uint, customer *model.Customer) error
	//DeleteCustomer(ctx context.Context, id uint) error
	//
	//GetProductByFilter(ctx context.Context, pageSize int, pageToken int, filter map[string][]string) ([]*model.Product, error)
	//CreateProduct(ctx context.Context, product *model.Product) error
	//UpdateProduct(ctx context.Context, id uint, product *model.Product) error
	//DeleteProduct(ctx context.Context, id uint) error
	//
	//GetBillsByFilter(ctx context.Context, pageSize int, pageToken int, filter map[string][]string) ([]*model.Bill, error)
	//CreateBill(ctx context.Context, bill *model.Bill) error
	//UpdateBill(ctx context.Context, id uint, bill *model.Bill) error
	//DeleteBill(ctx context.Context, id uint) error
	//
	//GetImportsByFilter(ctx context.Context, pageSize int, pageToken int, filter map[string][]string) ([]*model.Import, error)
	//CreateImport(ctx context.Context, imp *model.Import) error
	//UpdateImport(ctx context.Context, id uint, imp *model.Import) error
	//DeleteImport(ctx context.Context, id uint) error
}

func (c *DAO) InitDatabaseConnection() (err error) {
	conn := config.GetDBConfig()
	dsn := fmt.Sprintf("user=%s password=%s dbname=%s port=%s host=%s sslmode=disable",
		conn.Username, conn.Password, conn.DBName,
		conn.Port, conn.Host)
	gDB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	return
}

func (c *DAO) MigrateDataModel(ctx context.Context)(er error) {
	er = gDB.WithContext(ctx).AutoMigrate(&model.Customer{}, &model.User{}, &model.Provider{}, &model.Group{},
	&model.Product{}, &model.Bill{}, &model.BillLine{}, &model.Import{}, &model.ImportLine{})
	return
}

func (c *DAO) GetUsersByFilter(ctx context.Context, pageSize int, pageToken int, filter map[string][]string) ([]*model.User, error){
	var user = []*model.User{}
	thisDB := gDB.WithContext(ctx).Model(&model.User{})
	for k, v := range filter {
		thisDB = thisDB.Where(fmt.Sprintf("%s %s ?", k, v[0]), v[1])
	}
	offset := pageToken - 1
	thisDB = thisDB.Offset(offset).Limit(pageSize)
	er := thisDB.Find(&user).Error
	if er != nil {
		return nil, er
	}
	return user, nil
}

func (c *DAO) CreateUser(ctx context.Context, user *model.User) error{
	er := gDB.WithContext(ctx).Model(&model.User{}).Create(&user).Error
	return er
}

func (c *DAO) UpdateUser(ctx context.Context, id uint, user *model.User) error{
    er := gDB.WithContext(ctx).Model(&model.User{}).Where("user_id = ?", id).Save(&user).Error
    return er
}

func(c *DAO) DisableUser(ctx context.Context, id uint) error {
	er := gDB.WithContext(ctx).Where("user_id = ?", id).Update("is_active", false).Error
	return er
}

func(c *DAO) EnableUser(ctx context.Context, id uint) error {
	er := gDB.WithContext(ctx).Where("user_id = ?", id).Update("is_active", true).Error
	return er
}




func (c *DAO) GetBillsByFilter(ctx context.Context, pageSize int, pageToken int, filter map[string][]string) ([]*model.Bill, error) {
	var bil = []*model.Bill{}
	thisDB := gDB.WithContext(ctx).Model(&model.Bill{})
	for k, v := range filter {
		thisDB = thisDB.Where(fmt.Sprintf("%s %s ?", k, v[0]), v[1])
	}
	offset := pageToken - 1
	thisDB = thisDB.Offset(offset).Limit(pageSize)
	er := thisDB.Preload("bill_lines").Find(&bil).Error
	if er != nil {
		return nil, er
	}

	//for i, v := range bil {
	//	var lin = []*model.BillLine{}
	//	er = gDB.WithContext(ctx).Table("bill_lines").Where("bill = ?", v.BillId).Find(&lin).Error
	//	if er != nil {
	//		return nil, er
	//	}
	//	bil[i].Details = lin
	//}
	return bil, nil
}

func (c *DAO) GetImportsByFilter(ctx context.Context, pageSize int, pageToken int, filter map[string][]string) ([]*model.Import, error) {
	var imp = []*model.Import{}
	thisDB := gDB.WithContext(ctx).Table("customers")
	for k, v := range filter {
		thisDB = thisDB.Where(fmt.Sprintf("%s %s ?", k, v[0]), v[1])
	}
	offset := pageToken - 1
	thisDB = thisDB.Offset(offset).Limit(pageSize)
	er := thisDB.Preload("import_lines").Find(&imp).Error
	if er != nil {
		return nil, er
	}
	//for i, v := range imp {
	//	var lin = []*model.ImportLine{}
	//	er = gDB.WithContext(ctx).Table("import_lines").Where("import = ?", v.ImportID).Find(&lin).Error
	//	if er != nil {
	//		return nil, er
	//	}
	//	imp[i].Details = lin
	//}
	return imp, nil
}

type DAO struct{}

var d dao = &DAO{}

func GetDAO() dao {
	return d
}
