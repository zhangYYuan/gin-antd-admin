package conf

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper" //使用fsnotify和viper实现json格式配置文件
	"log"
)
type MysqlAdmin struct {
	UserName string `json:"username"`
	PassWord string `json:"password"`
	Path string `json:"path"`
	DBName string `json:"db_name"`
	Config string `json:"config"`
}

// jwt签名
type JWT struct {
	SigningKey string `json:"signingKey"`
}


type Config struct {
	MysqlAdmin MysqlAdmin `json:"mysqlAdmin"`
	JWT JWT	`json:"jwt"`
}


var GlobalConfig Config
var VTool *viper.Viper
func init() {
 	v := viper.New()
	v.SetConfigName("config")           //  设置配置文件名 (不带后缀)
	v.AddConfigPath("./conf") // 第一个搜索路径
	v.SetConfigType("json")

	err := v.ReadInConfig()
	if err !=nil  {
		log.Fatal("Fail read Config file ", err)
	}

	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		log.Println("Config file changed: ", e.Name)
		if err := v.Unmarshal(&GlobalConfig); err != nil {
			log.Fatal(err)
		}
	})

	if err := v.Unmarshal(&GlobalConfig); err != nil {
		log.Fatal(err)
	}
	VTool = v
}