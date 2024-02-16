package user

import (
	"81jcpd.cn/project-user/router"
	"81jcpd.cn/project-user/service"
	"github.com/gin-gonic/gin"
	"log"
)

// init 将当前小模块添加到路由列表中
func init() {
	log.Println("init user router...")
	router.Register(&RouterUser{})
}

// RouterUser 一个小模块
type RouterUser struct {
}

// Router	实现的 Router接口方法
func (*RouterUser) Router(r *gin.Engine) {
	handler := service.New(service.CacheRedis)
	r.POST("/login/getCaptcha", handler.GetCaptcha)
}
