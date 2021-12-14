package database

import (
	"clinic/config"
	"clinic/models"
)

func GetRecipes() (interface{}, error) {
	var recipes []models.Recipe
	if e := config.DB.Find(&recipes).Error; e !=nil {
		return nil, e
	}
	return recipes,nil
}