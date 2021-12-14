package models

import (
	"gorm.io/gorm"
)

type Dactor struct {
	gorm.Model
	Idadmin int `json:"idadmin" form:"idadmin"`
	Idsession int `json:"idsession" form:"idsession"`
	Name string `json:"name" form:"name"`
	Email string `json:"email" form:"email"`
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
	Phone string `json:"phone" form:"phone"`
	Nip int `json:"nip" form:"nip"`
	Room string `json:"room" form:"room"`
	Experience string `json:"experience" form:"experience"`
	
}