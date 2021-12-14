package models

import (
	"time"

	"gorm.io/gorm"
)

type Patient struct {
	gorm.Model
	Name string `json:"name" form:"name"`
	Dob time.Time `json:"dob" form:"dob"`
	Gender string `json:"gender" form:"gender"`
	Phone int `json:"phone" form:"phone"`
	Description string `json:"description" form:"description"`
	Address string `json:"address" form:"address"`
}