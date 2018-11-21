package database

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"os"
)

var dbEngine *xorm.Engine

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

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", dbUser, dbPassword, dbHost, 3306, dbDatabase)

	dbEngine, err = xorm.NewEngine("mysql", dsn)

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
