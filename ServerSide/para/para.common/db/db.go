package db
import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/go-ini/ini"
	"para/para.common/models"
	"fmt"
	"log"
)

var DB *gorm.DB


func init() {
	var err error

	cfg, err := ini.Load("../para.common/conf/app.ini")

	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
	}

	sec := cfg.Section("mysql")

	port := sec.Key("Port")
	host := sec.Key("Host")
	user := sec.Key("User")
	password := sec.Key("Pwd")
	dbname := sec.Key("DbName")

	// parseTime=True将mysql中的时间类型，自动解析为go结构体中的时间类型
	dbParams := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8&parseTime=True&loc=Local", user, password, host,port, dbname)
	fmt.Println(dbParams)
	DB, err = gorm.Open("mysql", dbParams)
	defer DB.Close()

	DB.AutoMigrate(models.User{})
	if err != nil {
		log.Fatal("error --->>>", err)
	}



	// 全局禁用表名复数
	DB.SingularTable(true)
	// 启用Logger, 显示详细日志
	DB.LogMode(true)

}

