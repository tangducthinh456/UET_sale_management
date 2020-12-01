package dao

import (
	"SaleManagement/config"
	"SaleManagement/model"
	"context"
	"log"
	"testing"
	"time"
)

var (
	customer = model.Customer{
		//CustomerID:   0,
		CustomerName: "ThinhTd",
		DOB:          time.Now(),
		PhoneNumber:  "0987654321",
		Address:      "Khoi 3 THij a",
	}
	user = model.User{
		//UserID:      0,
		Username:    "user",
		Name:        "Nguyễn User",
		Password:    []byte("Ilovemath"),
		Role:        "MANAGER",
		PhoneNumber: "0987678978",
		Email:       "ilovemath@gmail.com",
	}
	provider = model.Provider{
		//ProviderID:   0,
		ProviderName: "Công ty TNHH MTV mtp",
		PhoneNumber:  "123456789",
		Address:      "144, Xuân Thuy",
		Email:        "mtp@tangducthinh.com",
	}
	group = model.Group{
		//GroupID:   ,
		GroupName: "butbibabi",
		//Products:  nil,
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
		GroupID:       1,
		CreatedUserID: 1,
		Description:   "but bi thien long",
		CreatedAt:     time.Now(),
	}
	bill = model.Bill{
		//BillID:         0,
		CustomerID:     1,
		CustomerDetail: model.Customer{},
		Note:           "bill ban 1 but",
		CreatedUserID:  1,
		//CreatedUser:    model.User{},
		Details: []*model.BillLine{
			{
				//Bill:      model.Bill{},
				//Product:   model.Product{},
				//BillID:    0,
				ProductID: 1,
				Quantity:  10,
				Note:      "This is but bi",
			},
		},
		CreatedTime: time.Time{},
	}
	imp = model.Import{
		//ImportID:      0,
		Note:          "This is import note",
		ProviderID:    1,
		//Provider:      model.Provider{},
		CreatedUserID: 1,
		//CreatedUser:   model.User{},
		Details:       []*model.ImportLine{
			{
				//Import:    model.Import{},
				//Product:   model.Product{},
				//ImportID:  0,
				ProductID: 1,
				Quantity:  7,
				Note:      "This is import note",
			},
		},
		CreatedAt:     time.Time{},
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
	er = gDB.Model(&model.Provider{}).Create(&provider).Error
	if er != nil {
		log.Fatal(er)
	}
	er = gDB.Model(&model.Group{}).Create(&group).Error
	if er != nil {
		log.Fatal(er)
	}
	er = gDB.Model(&model.Customer{}).Create(&customer).Error
	if er != nil {
		log.Fatal(er)
	}
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
