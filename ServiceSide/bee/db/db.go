package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

)

var BeeDB *gorm.DB
// https://github.com/angao/gin-xorm-admin/blob/master/db/db.go
func init() {
	var err error
	 BeeDB, err = gorm.Open("mysql", "root:Li763a7689@(119.23.74.242:3306)/bee?charset=utf8&parseTime=True&loc=Local")
	 BeeDB.LogMode(true)
	if err != nil {
		fmt.Println("err -->", err)
	}
	InitTables()
	//defer BeeDB.Close()
}
