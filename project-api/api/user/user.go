package user

import (
	"81jcpd.cn/project-api/pkg/model"
	common "81jcpd.cn/project-common"
	loginservicev1 "81jcpd.cn/project-user/pkg/service/login.service.v1"
	"context"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/status"
	"net/http"
	"time"
)

type HandlerUser struct {
}

func New() *HandlerUser {
	return &HandlerUser{}
}

type MobileVo struct {
	Mobile string `json:"mobile"`
}

func (*HandlerUser) getCaptcha(ctx *gin.Context) {
	resp := &common.Resp{}
	var mobileVo MobileVo
	if err := ctx.ShouldBindJSON(&mobileVo); err != nil {
		ctx.JSON(http.StatusOK, resp.Fail(model.InvalidRequest, "无效请求,请携带手机号"))
		return
	}

	c, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	captchaResponse, err := LoginServiceClient.GetCaptcha(c, &loginservicev1.CaptchaMessage{Mobile: mobileVo.Mobile})
	if err != nil {
		fromError, _ := status.FromError(err)
		ctx.JSON(http.StatusOK, resp.Fail(common.BusinessCode(fromError.Code()), fromError.Message()))
		return
	}
	ctx.JSON(http.StatusOK, resp.Success(captchaResponse.Code))
}
