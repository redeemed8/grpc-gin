package login_service_v1

import (
	"81jcpd.cn/project-api/pkg/model"
	communutils "81jcpd.cn/project-common/utils"
	"81jcpd.cn/project-user/pkg/dao"
	"81jcpd.cn/project-user/pkg/repo"
	"81jcpd.cn/project-user/utils"
	"context"
	codes "google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"strconv"
	"time"
)

type LoginService struct {
	UnimplementedLoginServiceServer
	cache repo.Cache
}

const (
	CaptchaPrefix = "Register_"
)

func New() *LoginService {
	return &LoginService{
		cache: dao.Rc,
	}
}

func (ls *LoginService) GetCaptcha(ctx context.Context, request *CaptchaMessage) (*CaptchaResponse, error) {
	//	1.获取参数 手机号
	mobile := request.Mobile
	//	2.校验 手机号格式
	if ok := communutils.VerifyMobile(mobile); !ok {
		return nil, status.Error(codes.Code(model.NoLegalMobile), "手机号格式错误")
	}
	//	生成验证码
	symbol, _ := strconv.Atoi(mobile)
	code := utils.MakeCodeWithNumber(6, symbol)
	//	3.调用短信平台模拟
	go func() {
		time.Sleep(2 * time.Second)
		log.Println("Successfully send captcha to mobile : ", mobile)
		//	4.存入 redis
		err1 := ls.cache.Put(CaptchaPrefix+mobile, code, 5*time.Minute)
		if err1 != nil {
			log.Printf("Failed to save the mobile and captcha to redis : REGISTER_%s : %s , cause by : %v \n", mobile, code, err1)
		}
	}()
	return &CaptchaResponse{Code: code}, nil
}
