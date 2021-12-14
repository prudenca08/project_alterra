package models

import (
	"time"

	"gorm.io/gorm"
)

type Session struct {
	gorm.Model
	Idadmin int `json:"idadmin" form:"idadmin"`
	Iddoctor int `json:"iddoctor" form:"iddoctor"`
	Idpatient int `json:"idpatient" form:"idpatient"`
	Idschedule int `json:"idschedule" form:"idschedule"`
	Date time.Time `json:"date" form:"date"`
	Status string `json:"status" form:"status"`

}