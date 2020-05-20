package business_error

import "github.com/gin-gonic/gin"

type Renderable interface {
	Render(c *gin.Context)
	RenderWithMsg(c *gin.Context, msg string)
	Error() string

	GetCode() int
	GetMessage() string
	GetHttpCode() int
	RenderCodePath() []string
}
