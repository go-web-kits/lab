package response

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	go_redis "github.com/go-redis/redis"
	be "github.com/go-web-kits/lab/business_error"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"gopkg.in/go-playground/validator.v8"
)

func Error(c *gin.Context, realErr error) {
	switch e := errors.Cause(realErr).(type) {
	case be.Renderable:
		e.Render(c)
	case validator.ValidationErrors:
		ParamsError(c, realErr)
	default:
		if gorm.IsRecordNotFoundError(realErr) {
			errRender(c, be.NotFound, realErr)
		} else if go_redis.Nil == realErr {
			errRender(c, be.NotFound, realErr)
		} else {
			c.Set("unknown_error", realErr)
			errRender(c, be.Unknown, realErr, "log")
		}
	}
}

func ParamsError(c *gin.Context, realErr error) {
	errRender(c, be.ParamsError, realErr)
}

func errRender(c *gin.Context, err interface{}, realErr error, log ...interface{}) {
	if code, ok := err.(int); ok {
		err = be.CommonErrors[code]
	}

	e := err.(be.Renderable)
	if len(log) > 0 {
		fmt.Println(realErr)
	}
	if os.Getenv("ENV") == "PRO" {
		e.Render(c)
	} else {
		e.RenderWithMsg(c, realErr.Error())
	}
}
