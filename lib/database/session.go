package database

import (
	"clinic/config"
	"clinic/models"
)

func GetSessions() (interface{}, error) {
	var sessions []models.Session
	if e := config.DB.Find(&sessions).Error; e!=nil {
		return nil, e
	}
	return sessions, nil
}