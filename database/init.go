package database

import (
	"fmt"
	"github.com/InVisionApp/go-logger"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"hagnix-server-go1/database/models"
	"os"
)

var dbEngine *xorm.Engine
var logger = log.NewSimple()

func Init() error {
	var err error
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

	dbEngine, err = xorm.NewEngine("mysql", dsn)

	logger.Info("[SQL] Database connected.")

	logger.Info("[SQL] Registering database entities...")

	dbEngine.SetLogger(dBLogger{})

	//registerEntities()

	return err
}

func GetDBEngine() *xorm.Engine {
	if dbEngine == nil {
		err := Init()
		if err != nil {
			panic(err)
		}
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
