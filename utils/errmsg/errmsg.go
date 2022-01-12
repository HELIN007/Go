package errmsg

const (
	SUCCESS = 200
	ERROR   = 500
	//用户模块的错误code=1000...
	ErrorUsernameUsed   = 1001
	ErrorPasswordWrong  = 1002
	ErrorUserNotExist   = 1003
	ErrorTokenExist     = 1004
	ErrorTokenRuntime   = 1005
	ErrorTokenWrong     = 1006
	ErrorTokenTypeWrong = 1007
	ErrorUserNoRight    = 1008
	//文章模块的错误code=2000...
	ErrorArticleExist = 2001
	//分类模块的错误code=3000...
	ErrorCategoryUsed     = 3001
	ErrorCategoryNotExist = 3002
)

var codeMsg = map[int]string{
	SUCCESS:               "OK",
	ERROR:                 "FAIL",
	ErrorUsernameUsed:     "用户名已存在",
	ErrorPasswordWrong:    "密码错误",
	ErrorUserNotExist:     "用户不存在",
	ErrorTokenExist:       "TOKEN不存在",
	ErrorTokenRuntime:     "TOKEN已过期",
	ErrorTokenWrong:       "TOKEN不正确",
	ErrorTokenTypeWrong:   "TOKEN格式错误",
	ErrorArticleExist:     "文章不存在",
	ErrorCategoryUsed:     "分类已存在",
	ErrorCategoryNotExist: "分类不存在",
	ErrorUserNoRight:      "用户无权限",
}

func GetErrMsg(code int) string {
	return codeMsg[code]
}
