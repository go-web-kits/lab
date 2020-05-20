package business_error

import (
	"github.com/gin-gonic/gin"
)

type Common struct {
	Code     int    `json:"code"`
	Message  string `json:"message"`
	HttpCode int    `json:"-"`
}

func (e *Common) Render(c *gin.Context) {
	json := gin.H{
		"result": &e, "data": gin.H{},
	}
	c.Set("response_body", json)
	c.JSON(e.HttpCode, json)
}

func (e *Common) Error() string {
	return e.Message
}

func (e *Common) RenderWithMsg(c *gin.Context, msg string) {
	json := gin.H{
		"result": &Common{e.Code, "[" + e.Message + "] " + msg, e.HttpCode},
		"data":   gin.H{},
	}
	c.Set("response_body", json)
	c.JSON(e.HttpCode, json)
}

func (e *Common) GetCode() int {
	return e.Code
}

func (e *Common) RenderCodePath() []string {
	return []string{"result", "code"}
}

func (e *Common) GetMessage() string {
	return e.Message
}

func (e *Common) GetHttpCode() int {
	return e.HttpCode
}
