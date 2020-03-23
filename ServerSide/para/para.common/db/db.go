package db

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
	"para.common/conf"
)

var DB *gorm.DB


func init() {
	c := conf.GlobalConfig.MysqlAdmin
	source := fmt.Sprintf("%v:%v@(%v)/%v?%v", c.UserName, c.PassWord, c.Path, c.DBName, c.Config)
	db , err := gorm.Open("mysql", source)

	if (err != nil) {
		log.Fatal("数据库连接失败->", err)
	}
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	DB = db
}

