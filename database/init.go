package database

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/ivahaev/go-logger"
	"hagnix-server-go1/database/models"
	"os"
)

var dbEngine *xorm.Engine

func Init() {
	dbHost := os.Getenv("DB_HOST")
	dbDatabase := os.Getenv("DB_DATABASE")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")

	if len(dbHost) < 1 {
		dbHost = "127.0.0.1"
	}

	if len(dbDatabase) < 1 {
		dbDatabase = "rotmgprod"
	}

	if len(dbUser) < 1 {
		dbUser = "root"
	}

	if len(dbPassword) < 1 {
		dbPassword = "root"
	}

	logger.Info("[SQL] Connecting to database...")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", dbUser, dbPassword, dbHost, 3306, dbDatabase)

	dbEnginex, err := xorm.NewEngine("mysql", dsn)

	if err != nil {
		panic(err)
	}

	logger.Info("[SQL] Database connected.")

	logger.Info("[SQL] Registering database entities...")

	dbEnginex.SetLogger(dBLogger{})

	dbEngine = dbEnginex

	//registerEntities()

}

func GetDBEngine() *xorm.Engine {
	if dbEngine == nil {
		Init()
	}

	return dbEngine
}

func registerEntities() {
	dbEngine.Sync(&models.Accounts{})
	dbEngine.Sync(&models.Death{})
	dbEngine.Sync(&models.Giftcodes{})
	dbEngine.Sync(&models.Stats{})
	dbEngine.Sync(&models.Packages{})
	dbEngine.Sync(&models.Classstats{})
	dbEngine.Sync(&models.Guilds{})
	dbEngine.Sync(&models.Vaults{})
}
