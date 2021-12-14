package database

import (
	"clinic/config"
	"clinic/models"
)

func GetPatients() (interface{}, error) {
	var patients []models.Patient
	if e := config.DB.Find(&patients).Error; e !=nil {
		return nil, e
	}
	return patients,nil
}