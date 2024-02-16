package dao

import (
	"github.com/spf13/viper"
	"log"
)

func init() {
	viper.SetConfigName("application") //	设置 要读取的配置文件名称
	viper.AddConfigPath("config")      //	设置 要读取的配置文件的所在路径
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalln("project-api failed to read config application.yaml , cause by : ", err)
	}
	log.Println("project-api init config application.yaml...")
}
