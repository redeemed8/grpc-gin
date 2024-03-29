package router

import (
	"github.com/gin-gonic/gin"
)

// Router 路由接口
type Router interface {
	Router(r *gin.Engine)
}

// 路由列表
var routers []Router

// InitRouter 初始化路由列表
func InitRouter(r *gin.Engine) {
	for _, ro := range routers {
		ro.Router(r)
	}
}

// Register 注册路由到路由列表
func Register(ro ...Router) {
	routers = append(routers, ro...)
}
