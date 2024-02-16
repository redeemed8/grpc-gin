package main

import (
	_ "81jcpd.cn/project-api/api"
	_ "81jcpd.cn/project-api/pkg/dao"
	"81jcpd.cn/project-api/router"
	srv "81jcpd.cn/project-common"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

var (
	port    = viper.GetString("server.port")
	srvName = viper.GetString("server.project-name")
)

func main() {
	//	获取默认引擎
	r := gin.Default()
	//	根据路由列表，初始化引擎
	router.InitRouter(r)
	//	运行
	srv.Run(r, port, srvName, nil)
}
