package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	ID uint 		`gorm:"primary_key"`
	Phone string 	`gorm:"type:varchar(11);unique"`
	Password string	 `gorm:"size:2"`
}

type JwtToken struct {
	Token string `json:"token"`
}



func (User) TableName() string {
	return "profile"
}