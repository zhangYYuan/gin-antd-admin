package model

import "github.com/jinzhu/gorm"

type Category struct {
	gorm.Model
	CategoryId   int64  `json:"category_id"`
	CategoryName string `json:"category_name"`
	CategoryNo   int    `json:"category_no"`
}
