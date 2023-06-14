package model

import "gorm.io/gorm"

var Database *gorm.DB

type UserAttribute struct {
	gorm.Model
	Id_login  string `gorm:"size:255;not null;unique" json:"id_login"`
	Email     string `gorm:"size:255;not null;" json:"e-mail"`
	Firtsname string `gorm:"size:255" json:"firtsname"`
	Midlename string `gorm:"size:255" json:"midlename"`
	Lastname  string `gorm:"size:255" json:"lastname"`
	Job       string `gorm:"size:255" json:"job"`
}
