package database

import (
	"clinic/config"
	"clinic/models"
)

func GetDoctors() (interface{}, error) {
	var doctors []models.Dactor
	if e := config.DB.Find(&doctors).Error; e != nil {
		return nil, e
	}
	return doctors, nil
}