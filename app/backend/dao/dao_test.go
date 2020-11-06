package dao

import (
	"SaleManagement/config"
	"log"
	"testing"
)

func TestMain(m *testing.M) {
	config.Init()
	er := GetDAO().InitDatabaseConnection()
	if er != nil {
		log.Fatal(er)
	}
	clearDatabase()
	er = GetDAO().MigrateDataModel()
	if er != nil {
		log.Fatal(er)
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
}
