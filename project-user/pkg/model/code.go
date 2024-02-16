package model

import common "81jcpd.cn/project-common"

const (
	InvalidRequest common.BusinessCode = 1001 //	无效请求, 缺少必要的参数
	NoLegalMobile  common.BusinessCode = 2001 //	手机号不合法
)
