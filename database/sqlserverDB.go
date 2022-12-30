package database

import (
	"fmt"
	"go-payment/model"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

var Database *gorm.DB

func Init() {
	dsn := "sqlserver://sa:12345@localhost:8080?database=rgb"
	var err error
	Database, err = gorm.Open(sqlserver.Open(dsn), &gorm.Config{})

	Database.AutoMigrate(&model.Customer{})

	if err != nil {
		fmt.Println(err.Error())
	}
}
