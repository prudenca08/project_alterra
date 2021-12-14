package config

import (
	"clinic/models"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
func InitDB(){
	config := map[string]string{
		"DB_Username": "root",
		"DB_Password": "@D4ff44lv1",
		"DB_Host":"127.0.0.1",
		"DB_Port":"3306",
		"DB_Name":"db_rs",
	}
	// fmt.Sprintf("asdadasd")
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
	  config["DB_Username"], config["DB_Password"], config["DB_Host"], config["DB_Port"],config["DB_Name"])
	var e error
	DB, e = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if e != nil {
		panic(e)
	}
	InitMigrate()
}
func InitMigrate(){
	DB.AutoMigrate(&models.Admin{})
	DB.AutoMigrate(&models.Dactor{})
	DB.AutoMigrate(&models.Patient{})
	DB.AutoMigrate(&models.Recipe{})
	DB.AutoMigrate(&models.Schedule{})
	DB.AutoMigrate(&models.Session{})
}