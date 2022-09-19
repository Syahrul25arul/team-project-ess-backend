package database

import (
	"employeeSelfService/config"
	"employeeSelfService/logger"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetClientDb() *gorm.DB {
	db_host := config.DB_HOST
	db_user := config.DB_USER
	db_password := config.DB_PASSWORD
	db_dbname := config.DB_NAME
	db_port := config.DB_PORT

	// check if app run testing
	if config.TESTING == "true" {
		db_dbname = config.DB_NAME_TESTING
	}

	// set dsn
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", db_host, db_user, db_password, db_dbname, db_port)
	// connect to db postgres
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	// if there error connect to db, show error
	if err != nil {
		logger.Fatal("error connect to database : " + err.Error())
	}

	// if connect success give info
	logger.Info("connect to database success")
	return db
}
