package res

import (
	"github.com/gin-gonic/gin"
	"gvd_server/utils/valid"
	"net/http"
)

type Code int

type Response struct {
	Code Code   `json:"code"`
	Data any    `json:"data"`
	Msg  string `json:"msg"`
}

// ListResponse 列表数据
type ListResponse[T any] struct {
	List  []T `json:"list"`
	Count int `json:"count"`
}

const (
	SUCCESS   Code = 0
	ErrCode   Code = 7 // 系统错误
	ValidCode Code = 9 // 校验错误
)

// OK 返回成功的 数据和消息
func OK(data any, msg string, c *gin.Context) {
	c.JSON(http.StatusOK, Response{Code: SUCCESS, Data: data, Msg: msg})
}

// OKWithMsg 返回成功的消息
func OKWithMsg(msg string, c *gin.Context) {
	OK(map[string]any{}, msg, c)
}

// OKWithData 返回成功的 数据
func OKWithData(data any, c *gin.Context) {
	OK(data, "成功", c)
}

// OKWithList 分页查询使用
func OKWithList[T any](list []T, count int, c *gin.Context) {
	if len(list) == 0 {
		list = []T{}
	}
	OK(ListResponse[T]{
		List:  list,
		Count: count,
	}, "成功", c)
}

// Fail 失败
func Fail(code Code, data any, msg string, c *gin.Context) {
	c.JSON(http.StatusOK, Response{Code: code, Data: data, Msg: msg})
}

// FailWithMsg 错误信息
func FailWithMsg(msg string, c *gin.Context) {
	Fail(ErrCode, map[string]any{}, msg, c)
}

// FailWithValidMsg 参数校验的失败
func FailWithValidMsg(msg string, c *gin.Context) {
	Fail(ValidCode, map[string]any{}, msg, c)
}

// FailWithData 失败，但是返回数据
func FailWithData(data any, c *gin.Context) {
	Fail(ErrCode, data, "系统错误", c)
}

// FailWithError 校验失败，显示校验的错误信息
func FailWithError(err error, c *gin.Context) {
	errorMsg := valid.Error(err)
	Fail(ValidCode, map[string]any{}, errorMsg, c)
}

// FailWithValidError 校验失败，显示校验的错误信息，数据是对应的字段
func FailWithValidError(err error, obj any, c *gin.Context) {
	errorMsg, data := valid.ValidError(err, obj)
	Fail(ValidCode, data, errorMsg, c)
}
