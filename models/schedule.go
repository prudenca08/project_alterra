package models

import (
	"time"

	"gorm.io/gorm"
)

type Schedule struct {
	gorm.Model
	Idadmin int `json:"idadmin" form:"idadmin"`
	Date time.Time `json:"date" form:"date"`
}