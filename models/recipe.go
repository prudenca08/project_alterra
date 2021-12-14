package models

import "gorm.io/gorm"

type Recipe struct {
	gorm.Model
	Idsession int `json:"idsession" form:"idsession"`
	Title string `json:"title" form:"title"`
	Detail_recipe string `json:"detail_recipe" form:"detail_recipe"`
}