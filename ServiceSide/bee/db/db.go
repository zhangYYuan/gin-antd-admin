package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

var BeeDB *gorm.DB

func init() {
	var db, err = gorm.Open("mysql", "root:Li763a7689@()/bee?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("err -->", err)
	}
	defer db.Close()
}

