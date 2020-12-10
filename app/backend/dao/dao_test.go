package dao

import (
	"SaleManagement/config"
	"SaleManagement/model"
	"context"
	"log"
	"os/exec"
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
		GroupID:       1,
		CreatedUserID: 1,
		Description:   "but bi thien long",
		CreatedAt:     time.Now(),
	}
	bill = model.Bill{
		CustomerID:     1,
		Note:           "bill ban 1 but",
		CreatedUserID:  1,
		Details: []*model.BillLine{
			{
				ProductID: 1,
				Quantity:  10,
				Note:      "This is but bi",
			},
			{
				ProductID: 2,
				Quantity:  13,
				Note:      "This is but bi",
			},
		},
	}
	imp = model.Import{
		//ImportID:      0,
		Note:          "This is import note",
		ProviderID:    1,
		CreatedUserID: 1,
		Details:       []*model.ImportLine{
			{
				ProductID: 1,
				Quantity:  7,
				Note:      "This is import note",
			},
			{
				ProductID: 2,
				Quantity:  5,
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
    genData()
	//er = gDB.Model(&model.Customer{}).Create(&customer).Error
	//if er != nil{
	//	log.Fatal(er)
	//}
}

func genData(){
	//dat, err := ioutil.ReadFile("data.sql")
	//if err != nil{
	//	log.Fatal(err)
	//}
	//err = gDB.Exec(string(dat)).Error
	//if err != nil{
	//	log.Fatal(err)
	//}
	cmd := exec.Command("psql", "postgres", "-d", "tool", "-f", "data.sql")
	err := cmd.Run()
	if err != nil{
		log.Fatal(err)
	}
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
