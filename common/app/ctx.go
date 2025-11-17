package app

import (
	"errors"

	"github.com/gin-gonic/gin"
)

type M map[string]interface{}
type I []interface{}

func Err(errs ...interface{}) {
	msg := ""
	for i := 0; i < len(errs); i++ {
		err := errs[i]
		switch err.(type) {
		case string:
			msg += err.(string) + ","
		case error:
			msg += err.(error).Error()
		}
	}
	panic(msg)
}

func Ok(c *gin.Context, msg string, data interface{}) {
	res := M{}
	res["code"] = 0
	res["msg"] = msg
	res["data"] = data
	c.JSON(200, res)
}

func NewError(params ...interface{}) error {
	errStr := ""
	for i := 0; i < len(params); i++ {
		p := params[i]
		switch p.(type) {
		case string:
			errStr += p.(string)
		case error:
			if errStr != "" {
				errStr += ":"
			}
			errStr += p.(error).Error()
		default:
			errStr += "发生错误!"
		}
	}
	return errors.New(errStr)
}
