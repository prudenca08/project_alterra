package database

import (
	"clinic/config"
	"clinic/models"
)

func GetAdmins() (interface{}, error) {
	var admins []models.Admin
	if e := config.DB.Find(&admins).Error; e != nil {
		return nil, e
	}
	return admins,nil
}


// // input admin
// 	var dataadmin = models.Admin {
// 		Name: "dapo itel",
// 		Email: "email@gmail.com",
// 		Username: "dapo gatel",
// 		Password: "kadal",
// 	}


func CreateAdmin(data models.Admin) (interface{}, error) {
	if e := config.DB.Create(&data).Error; e != nil {
		return nil, e
	}
	return data, nil
}

