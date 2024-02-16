package main

import (
	srv "81jcpd.cn/project-common"
	_ "81jcpd.cn/project-user/api"
	"81jcpd.cn/project-user/router"
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
	//	grpc 服务注册
	grpc := router.RegisterGrpc()
	//	运行
	srv.Run(r, port, srvName, grpc.Stop)
}
