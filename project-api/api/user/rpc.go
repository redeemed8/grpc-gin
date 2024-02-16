package user

import (
	loginServiceV1 "81jcpd.cn/project-user/pkg/service/login.service.v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

var LoginServiceClient loginServiceV1.LoginServiceClient

func InitRpcUserClient() {
	conn, err := grpc.Dial("127.0.0.1:9091", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect , cause by : %v \n", err)
	}

	LoginServiceClient = loginServiceV1.NewLoginServiceClient(conn)

}
