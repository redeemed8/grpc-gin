package main

import (
	srv "81jcpd.cn/project-common"
	_ "81jcpd.cn/project-user/api"
	"81jcpd.cn/project-user/router"
	"github.com/gin-gonic/gin"
)

func main() {
	//	获取默认引擎
	r := gin.Default()
	//	根据路由列表，初始化引擎
	router.InitRouter(r)
	//	运行
	srv.Run(r, ":8081", "project-user")
}
