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

func NewError(msg string, err error) error {
	if err == nil {
		return errors.New(msg)
	}
	return errors.New(msg + ":" + err.Error())
}
