package service

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// UserHandler	用户模块
type UserHandler struct {
}

// GetCaptcha	获取验证码
func (*UserHandler) GetCaptcha(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "getCaptcha success")
}
