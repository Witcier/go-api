package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:""data`
}

const (
	// 请求成功
	HTTP_OK = 20000
	// 缺少参数或参数错误
	ERROR_VALIDATE = 20001
	// 数据库错误
	ERROR_DB_MISTAKE = 30000
	// 验证码过期
	ERROR_CAPTCHA_EXPIRED = 30001
	// 验证码错误
	ERROR_CAPTCHA_MISMATCH = 30002
	// 密码错误
	ERROR_PASSWORD_MISTAKE = 30003
	// 用户不存在
	ERROR_USER_NOT_EXISTS = 30004
	// 未授权未登录
	ERROR_UNAUTHORIZED = 40001
	// 没有权限
	ERROR_PERMISSION_DENY = 40003
	// not found
	ERROR_NOT_FOUND = 40004
	// 内部错误
	ERROR_INTERNAL = 50000
)

func Resp(code int, data interface{}, msg string, c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		code,
		msg,
		data,
	})
}

func Success(c *gin.Context, data interface{}) {
	Resp(HTTP_OK, data, "success", c)
}

func Error(code int, msg string, c *gin.Context) {
	Resp(code, map[string]interface{}{}, msg, c)
	c.Abort()

	return
}

func ValidateFail(c *gin.Context) {
	Error(ERROR_VALIDATE, "validate fail", c)
}

func DbError(c *gin.Context) {
	Error(ERROR_DB_MISTAKE, "Db error", c)
}

func CaptchaExpire(c *gin.Context) {
	Error(ERROR_CAPTCHA_EXPIRED, "captcha expire", c)
}

func CaptchaMistake(c *gin.Context) {
	Error(ERROR_CAPTCHA_MISMATCH, "captcha mistake", c)
}

func PasswordMistake(c *gin.Context) {
	Error(ERROR_PASSWORD_MISTAKE, "password mistake", c)
}

func UserNotExists(c *gin.Context) {
	Error(ERROR_USER_NOT_EXISTS, "user not exists", c)
}

func Unauthorized(c *gin.Context) {
	Error(ERROR_UNAUTHORIZED, "unauthorized", c)
}

func Denied(c *gin.Context) {
	Error(ERROR_PERMISSION_DENY, "denied", c)
}

func NotFound(c *gin.Context) {
	Error(ERROR_NOT_FOUND, "not found", c)
}

func ErrorInternal(c *gin.Context) {
	Error(ERROR_INTERNAL, "error internal", c)
}
