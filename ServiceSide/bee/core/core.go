package core

import (
	"bee/global"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

const defaultConfigFile = "config.yaml"

func init() {

	v := viper.New()
	v.SetConfigFile(defaultConfigFile)

	err := v.ReadInConfig()

	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	v.WatchConfig()

	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed: ", e.Name)
		if err := v.Unmarshal(&global.BeeConfig); err != nil {
			fmt.Println(err)
		}
	})
	fmt.Println("core init------->>", &global.BeeConfig)
	if err := v.Unmarshal(&global.BeeConfig); err != nil {
		fmt.Println(err)
	}
	global.BEE_VIP = v
}
