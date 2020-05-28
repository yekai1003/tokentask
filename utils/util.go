package utils

const (
	RECODE_OK        = "0"
	RECODE_DBERR     = "4001"
	RECODE_LOGINERR  = "4002"
	RECODE_PARAMERR  = "4003"
	RECODE_SYSERR    = "4004"
	RECODE_ETHERR    = "4105"
	RECODE_NOTEXISTS = "4106"
	RECODE_UNKNOWERR = "4107"
)

var recodeText = map[string]string{
	RECODE_OK:        "成功",
	RECODE_DBERR:     "数据库操作错误",
	RECODE_LOGINERR:  "用户登录失败",
	RECODE_PARAMERR:  "参数错误",
	RECODE_SYSERR:    "系统错误",
	RECODE_ETHERR:    "与bcos交互失败",
	RECODE_NOTEXISTS: "证书不存在",
	RECODE_UNKNOWERR: "未知错误",
}

func RecodeText(code string) string {
	str, ok := recodeText[code]
	if ok {
		return str
	}
	return recodeText[RECODE_UNKNOWERR]
}
