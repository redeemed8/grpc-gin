package user

import (
	"81jcpd.cn/project-api/router"
	"github.com/gin-gonic/gin"
	"log"
)

type RouterUser struct {
}

func init() {
	log.Println("init user router...11111")
	ru := &RouterUser{}
	router.Register(ru)
}

func (*RouterUser) Router(r *gin.Engine) {
	//	初始化 grpc的客户端连接
	InitRpcUserClient()
	handler := New()
	r.POST("/project/login/getCaptcha", handler.getCaptcha)
}
