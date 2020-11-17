package dao

import (
	"SaleManagement/config"
	"SaleManagement/model"
	"context"
	"fmt"
	"log"
	"testing"
	"time"
)

var (
	customer1 = model.Customer{
		//CustomerID:   0,
		CustomerName: "ThinhTd",
		DOB:          time.Now(),
		PhoneNumber:  "0987654321",
		Address:      "Khoi 3 THij a",
	}
	user = model.User{
		//UserID:      0,
		Username:    "huyennt",
		Name:        "Nguyễn Thanh Huyền",
		Password:    "Ilovemath",
		Role:        "MANAGER",
		PhoneNumber: "0987678978",
		Email:       "huyenlovemath@gmail.com",
	}
	group = model.Group{
		GroupID:   0,
		GroupName: "",
		Products:  nil,
	}
	product = model.Product{
		//ProductID:   0,
		ProductName:   "but bi ba bi",
		Cost:          15000,
		Price:         12000,
		Quantity:      12,
		Brand:         "Thien Long",
		Position:      "Ngan 1",
		IsOnSale:      true,
		GroupID:       0,
		CreatedUserID: 0,
		Description:   "but bi thien long",
		CreatedAt:     time.Now(),
	}
	bill1 = model.Bill{
		//BillID:         0,
		CustomerID:     0,
		CustomerDetail: model.Customer{},
		Note:           "",
		CreatedUserID:  0,
		CreatedUser:    model.User{},
		Details:        nil,
		CreatedTime:    time.Time{},
	}
)

func TestMain(m *testing.M) {
	config.Init("")
	er := GetDAO().InitDatabaseConnection()
	if er != nil {
		log.Fatal(er)
	}

	clearDatabase()
	er = GetDAO().MigrateDataModel(context.Background())
	if er != nil {
		log.Fatal(er)
	}
	er = gDB.Model(&model.User{}).Create(&user).Error
	if er != nil {
		log.Fatal(er)
	}
	u := &model.User{}
	gDB.Model(&model.User{}).Where("user_id = ?", 1).Take(&u)
	fmt.Println(u)
}

func createSampleDatabase() {

}

func clearDatabase() {
	gDB.Migrator().DropTable("customers")
	gDB.Migrator().DropTable("providers")
	gDB.Migrator().DropTable("products")
	gDB.Migrator().DropTable("groups")
	gDB.Migrator().DropTable("bills")
	gDB.Migrator().DropTable("bill_lines")
	gDB.Migrator().DropTable("imports")
	gDB.Migrator().DropTable("import_lines")
	gDB.Migrator().DropTable("users")
}
