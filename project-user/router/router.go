package router

import (
	loginServiceV1 "81jcpd.cn/project-user/pkg/service/login.service.v1"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"log"
	"net"
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

type gRPCConfig struct {
	Addr         string
	RegisterFunc func(*grpc.Server)
}

func RegisterGrpc() *grpc.Server {
	c := gRPCConfig{
		Addr: viper.GetString("grpc.addr"),
		RegisterFunc: func(server *grpc.Server) {
			loginServiceV1.RegisterLoginServiceServer(server, loginServiceV1.New())
		},
	}

	s := grpc.NewServer()
	c.RegisterFunc(s)

	lis, err := net.Listen("tcp", c.Addr)
	if err != nil {
		log.Println("Failed to listen : ", c.Addr)
	}
	go func() {
		err = s.Serve(lis)
		if err != nil {
			log.Printf("server %s started error , cause by : %v \n", viper.GetString("grpc.name"), err)
			return
		}
	}()
	return s
}
