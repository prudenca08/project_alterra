package database

import (
	"clinic/config"
	"clinic/models"
)

func GetSchedules() (interface{}, error) {
	var schedules []models.Schedule
	if e := config.DB.Find(&schedules).Error; e!=nil {
		return nil, e
	}
	return schedules, nil
}