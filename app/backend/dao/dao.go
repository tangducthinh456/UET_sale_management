package dao

import (
	"SaleManagement/config"
	"SaleManagement/model"
	"context"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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

	GetProvidersByFilter(ctx context.Context, pageSize int, pageToken int, filter map[string][]string) ([]*model.Provider, error)
	CreateProvider(ctx context.Context, provider *model.Provider) error
	UpdateProvider(ctx context.Context, id uint, provider *model.Provider) error
	DisableProvider(ctx context.Context, id uint) error
	EnableProvider(ctx context.Context, id uint) error

	
	GetCustomersByFilter(ctx context.Context, pageSize int, pageToken int, filter map[string][]string) ([]*model.Customer, error)
	CreateCustomer(ctx context.Context, customer *model.Customer) error
	UpdateCustomer(ctx context.Context, id uint, customer *model.Customer) error
	DisableCustomer(ctx context.Context, id uint) error
	EnableCustomer(ctx context.Context, id uint) error
	
	GetProductsByFilter(ctx context.Context, pageSize int, pageToken int, filter map[string][]string) ([]*model.Product, error)
	CreateProduct(ctx context.Context, product *model.Product) error
	UpdateProduct(ctx context.Context, id uint, product *model.Product) error
	DisableProduct(ctx context.Context, id uint) error
	EnableProduct(ctx context.Context, id uint) error

	GetGroup(ctx context.Context, name string) (*model.Group, error)
	GetGroups(ctx context.Context) ([]*model.Group, error)
	CreateGroup(ctx context.Context, group *model.Group) error

	//
	GetBillsByFilter(ctx context.Context, pageSize int, pageToken int, filter map[string][]string) ([]*model.Bill, error)
	CreateBill(ctx context.Context, bill *model.Bill) error
	UpdateBill(ctx context.Context, id uint, bill *model.Bill) error
	DeleteBill(ctx context.Context, id uint) error
	//
	GetImportsByFilter(ctx context.Context, pageSize int, pageToken int, filter map[string][]string) ([]*model.Import, error)
	CreateImport(ctx context.Context, imp *model.Import) error
	UpdateImport(ctx context.Context, id uint, imp *model.Import) error
	DeleteImport(ctx context.Context, id uint) error
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
	return gDB.WithContext(ctx).Transaction(func(tx *gorm.DB) (er error) {
		user.IsActive = true
		er = tx.WithContext(ctx).Model(&model.User{}).Create(&user).Error
		return er
	})
}

func (c *DAO) UpdateUser(ctx context.Context, id uint, user *model.User) error{
	user.UserID = id
	er := gDB.WithContext(ctx).Model(&model.User{}).Where(&model.User{UserID: id}).Save(&user).Error
    return er
}

func(c *DAO) DisableUser(ctx context.Context, id uint) error {
	er := gDB.WithContext(ctx).Model(&model.User{}).Where(&model.User{UserID: id}).Updates(&model.User{IsActive: false}).Error
	return er
}

func(c *DAO) EnableUser(ctx context.Context, id uint) error {
	er := gDB.WithContext(ctx).Model(&model.User{}).Where(&model.User{UserID: id}).Updates(&model.User{IsActive: true}).Error
	return er
}

func (c *DAO) GetProvidersByFilter(ctx context.Context, pageSize int, pageToken int, filter map[string][]string) ([]*model.Provider, error){
	var prov = []*model.Provider{}
	thisDB := gDB.WithContext(ctx).Model(&model.Provider{})
	for k, v := range filter {
		thisDB = thisDB.Where(fmt.Sprintf("%s %s ?", k, v[0]), v[1])
	}
	offset := pageToken - 1
	thisDB = thisDB.Offset(offset).Limit(pageSize)
	er := thisDB.Find(&prov).Error
	if er != nil {
		return nil, er
	}
	return prov, nil
}

func (c *DAO) CreateProvider(ctx context.Context, provider *model.Provider) error{
	return gDB.WithContext(ctx).Transaction(func(tx *gorm.DB) (er error) {
		provider.IsActive = true
		er = tx.WithContext(ctx).Model(&model.Provider{}).Create(&provider).Error
		return er
	})
}

func (c *DAO) UpdateProvider(ctx context.Context, id uint, provider *model.Provider) error{
	provider.ProviderID = id
	er := gDB.WithContext(ctx).Model(&model.Provider{}).Where(&model.Provider{ProviderID: id}).Save(&provider).Error
	return er
}

func(c *DAO) DisableProvider(ctx context.Context, id uint) error {
	er := gDB.WithContext(ctx).Model(&model.Provider{}).Where(&model.Provider{ProviderID: id}).Updates(&model.Provider{IsActive: false}).Error
	return er
}

func(c *DAO) EnableProvider(ctx context.Context, id uint) error {
	er := gDB.WithContext(ctx).Model(&model.Provider{}).Where(&model.Provider{ProviderID: id}).Updates(&model.Provider{IsActive: true}).Error
	return er
}

func (c *DAO) GetCustomersByFilter(ctx context.Context, pageSize int, pageToken int, filter map[string][]string) ([]*model.Customer, error){
	var prov = []*model.Customer{}
	thisDB := gDB.WithContext(ctx).Model(&model.Customer{})
	for k, v := range filter {
		thisDB = thisDB.Where(fmt.Sprintf("%s %s ?", k, v[0]), v[1])
	}
	offset := pageToken - 1
	thisDB = thisDB.Offset(offset).Limit(pageSize)
	er := thisDB.Find(&prov).Error
	if er != nil {
		return nil, er
	}
	return prov, nil
}

func (c *DAO) CreateCustomer(ctx context.Context, customer *model.Customer) error{
	return gDB.WithContext(ctx).Transaction(func(tx *gorm.DB) (er error) {
		customer.IsActive = true
		er = tx.WithContext(ctx).Model(&model.Customer{}).Create(&customer).Error
		return er
	})
}

func (c *DAO) UpdateCustomer(ctx context.Context, id uint, customer *model.Customer) error{
	customer.CustomerID = id
	er := gDB.WithContext(ctx).Model(&model.Customer{}).Where(&model.Customer{CustomerID: id}).Save(&customer).Error
	return er
}

func(c *DAO) DisableCustomer(ctx context.Context, id uint) error {
	er := gDB.WithContext(ctx).Model(&model.Customer{}).Where(&model.Customer{CustomerID: id}).Updates(&model.Customer{IsActive: false}).Error
	return er
}

func(c *DAO) EnableCustomer(ctx context.Context, id uint) error {
	er := gDB.WithContext(ctx).Model(&model.Customer{}).Where(&model.Customer{CustomerID: id}).Updates(&model.Customer{IsActive: true}).Error
	return er
}

func (c *DAO) GetProductsByFilter(ctx context.Context, pageSize int, pageToken int, filter map[string][]string) ([]*model.Product, error){
	var prov = []*model.Product{}
	thisDB := gDB.WithContext(ctx).Model(&model.Product{})
	for k, v := range filter {
		thisDB = thisDB.Where(fmt.Sprintf("%s %s ?", k, v[0]), v[1])
	}
	offset := pageToken - 1
	thisDB = thisDB.Offset(offset).Limit(pageSize)
	er := thisDB.Preload(clause.Associations).Find(&prov).Error
	if er != nil {
		return nil, er
	}
	return prov, nil
}

func (c *DAO) CreateProduct(ctx context.Context, product *model.Product) error{
	return gDB.WithContext(ctx).Transaction(func(tx *gorm.DB) (er error) {
		product.IsOnSale = true
		er = tx.WithContext(ctx).Model(&model.Product{}).Create(&product).Error
		return er
	})
}

func (c *DAO) UpdateProduct(ctx context.Context, id uint, product *model.Product) error{
	product.ProductID = id
	product.IsOnSale = true
	er := gDB.WithContext(ctx).Model(&model.Product{}).Where(&model.Product{ProductID: id}).Save(&product).Error
	return er
}

func(c *DAO) DisableProduct(ctx context.Context, id uint) error {
	er := gDB.WithContext(ctx).Model(&model.Product{}).Where(&model.Product{ProductID: id}).Update("is_on_sale", false).Error
	return er
}

func(c *DAO) EnableProduct(ctx context.Context, id uint) error {
	er := gDB.WithContext(ctx).Model(&model.Product{}).Where(&model.Product{ProductID: id}).Update("is_on_sale", true).Error
	return er
}

func (c *DAO) GetGroup(ctx context.Context, name string) (*model.Group, error) {
	group := &model.Group{}
	er := gDB.WithContext(ctx).Model(&model.Group{}).Where(&model.Group{GroupName: name}).Find(&group).Error
	if er != nil{
		return nil, er
	}
	return group, nil
}

func (c *DAO) GetGroups(ctx context.Context) ([]*model.Group, error){
	var groups []*model.Group
	er := gDB.WithContext(ctx).Model(&model.Group{}).Find(&groups).Error
	if er != nil{
		return nil, er
	}
	return groups, nil
}

func (c *DAO) CreateGroup(ctx context.Context, group *model.Group) error {
	er := gDB.WithContext(ctx).Model(&model.Group{}).Create(&group).Error
	if er != nil{
		return er
	}
	return nil
}

func (c *DAO) GetBillsByFilter(ctx context.Context, pageSize int, pageToken int, filter map[string][]string) ([]*model.Bill, error){
	var prov = []*model.Bill{}
	thisDB := gDB.WithContext(ctx).Model(&model.Bill{})
	for k, v := range filter {
		thisDB = thisDB.Where(fmt.Sprintf("%s %s ?", k, v[0]), v[1])
	}
	offset := pageToken - 1
	thisDB = thisDB.Offset(offset).Limit(pageSize)
	er := thisDB.Preload(clause.Associations).Find(&prov).Error
	if er != nil {
		return nil, er
	}
	return prov, nil
}

func (c *DAO) CreateBill(ctx context.Context, bill *model.Bill) error{
	return gDB.WithContext(ctx).Transaction(func(tx *gorm.DB) (er error) {
		er = tx.WithContext(ctx).Model(&model.Bill{}).Create(&bill).Error
		if er != nil{
			return er
		}
		for _, v := range bill.Details{
			er = tx.WithContext(ctx).Model(&model.Bill{}).Where(&model.Bill{BillID: v.ProductID}).
				UpdateColumn("quantity", gorm.Expr("quantity - ?", v.Quantity)).Error
			if er != nil{
				return er
			}
		}
		return er
	})
}

func (c *DAO) UpdateBill(ctx context.Context, id uint, bill *model.Bill) error{
	return gDB.WithContext(ctx).Transaction(func(tx *gorm.DB) (er error) {
		billBefore := &model.Bill{}
		er = tx.WithContext(ctx).Model(&model.Bill{}).Where(&model.Bill{BillID: id}).Take(&billBefore).Error
		if er != nil{
			return er
		}
		bill.BillID = id
		er = tx.WithContext(ctx).Model(&model.Bill{}).Where(&model.Bill{BillID: id}).Save(&bill).Error
		if er != nil{
			return er
		}
		for i, v := range bill.Details{
			er = tx.WithContext(ctx).Model(&model.Bill{}).Where(&model.Bill{BillID: v.ProductID}).
				UpdateColumn("quantity", gorm.Expr("quantity - ?", v.Quantity - billBefore.Details[i].Quantity)).Error
			if er != nil{
				return er
			}
		}
		return nil
	})
}

func(c *DAO) DeleteBill(ctx context.Context, id uint) error {
	return gDB.WithContext(ctx).Transaction(func(tx *gorm.DB) (er error) {
		bill := &model.Bill{}
		er = tx.WithContext(ctx).Model(&model.Bill{}).Where(&model.Bill{BillID: id}).Take(&bill).Error
		if er != nil{
			return er
		}
		for _, v := range bill.Details{
			er = tx.WithContext(ctx).Model(&model.Bill{}).Where(&model.Bill{BillID: v.ProductID}).
				UpdateColumn("quantity", gorm.Expr("quantity + ?", v.Quantity)).Error
			if er != nil{
				return er
			}
		}
		er = tx.WithContext(ctx).Model(&model.Bill{}).Where(&model.Bill{BillID: id}).Delete(&model.Bill{}).Error
		return er
	})
}

func (c *DAO) GetImportsByFilter(ctx context.Context, pageSize int, pageToken int, filter map[string][]string) ([]*model.Import, error){
	var prov = []*model.Import{}
	thisDB := gDB.WithContext(ctx).Model(&model.Import{})
	for k, v := range filter {
		thisDB = thisDB.Where(fmt.Sprintf("%s %s ?", k, v[0]), v[1])
	}
	offset := pageToken - 1
	thisDB = thisDB.Offset(offset).Limit(pageSize)
	er := thisDB.Preload(clause.Associations).Find(&prov).Error
	if er != nil {
		return nil, er
	}
	return prov, nil
}

func (c *DAO) CreateImport(ctx context.Context, imp *model.Import) error{
	return gDB.WithContext(ctx).Transaction(func(tx *gorm.DB) (er error) {
		er = tx.WithContext(ctx).Model(&model.Import{}).Create(&imp).Error
		if er != nil{
			return er
		}
		for _, v := range imp.Details{
			er = tx.WithContext(ctx).Model(&model.Import{}).Where(&model.Import{ImportID: v.ProductID}).
				UpdateColumn("quantity", gorm.Expr("quantity + ?", v.Quantity)).Error
			if er != nil{
				return er
			}
		}
		return er
	})
}

func (c *DAO) UpdateImport(ctx context.Context, id uint, imp *model.Import) error{
	return gDB.WithContext(ctx).Transaction(func(tx *gorm.DB) (er error) {
		impBefore := &model.Import{}
		er = tx.WithContext(ctx).Model(&model.Import{}).Where(&model.Import{ImportID: id}).Take(&impBefore).Error
		if er != nil{
			return er
		}
		imp.ImportID = id
		er = tx.WithContext(ctx).Model(&model.Import{}).Where(&model.Import{ImportID: id}).Save(&imp).Error
		if er != nil{
			return er
		}
		for i, v := range imp.Details{
			er = tx.WithContext(ctx).Model(&model.Import{}).Where(&model.Import{ImportID: v.ImportID}).
				UpdateColumn("quantity", gorm.Expr("quantity + ?", v.Quantity - impBefore.Details[i].Quantity)).Error
			if er != nil{
				return er
			}
		}
		return nil
	})
}

func(c *DAO) DeleteImport(ctx context.Context, id uint) error {
	return gDB.WithContext(ctx).Transaction(func(tx *gorm.DB) (er error) {
		imp := &model.Import{}
		er = tx.WithContext(ctx).Model(&model.Import{}).Where(&model.Import{ImportID: id}).Take(&imp).Error
		if er != nil{
			return er
		}
		for _, v := range imp.Details{
			er = tx.WithContext(ctx).Model(&model.Import{}).Where(&model.Import{ImportID: v.ProductID}).
				UpdateColumn("quantity", gorm.Expr("quantity - ?", v.Quantity)).Error
			if er != nil{
				return er
			}
		}
		er = tx.WithContext(ctx).Model(&model.Import{}).Where(&model.Import{ImportID: id}).Delete(&model.Import{}).Error
		return er
	})
}

type DAO struct{}



var d dao = &DAO{}

func GetDAO() dao {
	return d
}
