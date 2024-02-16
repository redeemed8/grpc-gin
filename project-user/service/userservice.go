package service

import (
	common "81jcpd.cn/project-common"
	communutils "81jcpd.cn/project-common/utils"
	"81jcpd.cn/project-user/pkg/dao"
	"81jcpd.cn/project-user/pkg/model"
	"81jcpd.cn/project-user/pkg/model/vo"
	"81jcpd.cn/project-user/pkg/repo"
	"81jcpd.cn/project-user/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

// UserHandler	用户模块
type UserHandler struct {
	cache repo.Cache
}

type CacheType string

const (
	CacheRedis CacheType = "redis_"
	CacheMysql CacheType = "mysql_"
	CacheMongo CacheType = "mongo_"
	Memcahce   CacheType = "memcache_"
)

func New(type_ CacheType) *UserHandler {
	var cache_ repo.Cache
	switch type_ {
	case CacheRedis:
		cache_ = dao.Rc
	case CacheMongo:
		fmt.Println("wait to do...")
	case CacheMysql:
		fmt.Println("wait to do...")
	case Memcahce:
		fmt.Println("wait to do...")
	default:
		cache_ = dao.Rc
	}

	return &UserHandler{
		cache: cache_,
	}
}

const (
	CaptchaPrefix = "Register_"
)

// GetCaptcha	获取验证码
func (h *UserHandler) GetCaptcha(ctx *gin.Context) {
	resp := &common.Resp{}
	//	1.获取参数 手机号
	var mobileVo vo.MobileVo
	if err := ctx.ShouldBind(&mobileVo); err != nil {
		ctx.JSON(http.StatusOK, resp.Fail(model.InvalidRequest, "无效请求,请携带手机号"))
		return
	}
	//	2.校验 手机号格式
	if ok := communutils.VerifyMobile(mobileVo.Mobile); !ok {
		ctx.JSON(http.StatusOK, resp.Fail(model.NoLegalMobile, "手机号格式错误"))
		return
	}
	code := utils.MakeCodeWithNumber(6, 0)
	//	3.调用短信平台模拟
	go func() {
		time.Sleep(2 * time.Second)
		log.Println("Successfully send captcha to mobile : ", mobileVo.Mobile)
		//	4.存入 redis
		err1 := h.cache.Put(CaptchaPrefix+mobileVo.Mobile, code, 5*time.Minute)
		if err1 != nil {
			log.Printf("Failed to save the mobile and captcha to redis : REGISTER_%s : %s , cause by : %v \n", mobileVo.Mobile, code, err1)
		}
	}()
	ctx.JSON(http.StatusOK, resp.Success("验证码已发送: "+code))
}
