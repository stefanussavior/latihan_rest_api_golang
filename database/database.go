package database

import (
	"latihan_project_golang/REST_API_FIBER/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Database *gorm.DB
var DATABSE_URI string = "root:root@tcp(localhost:3306)/gorm?charset=ut"

func Connect() error {
	var err error

	Database, err = gorm.Open(mysql.Open(DATABSE_URI), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
	})
	if err != nil {
		panic(err)
	}
	Database.AutoMigrate(&model.Student{})
	return nil
}
