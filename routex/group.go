package routex

import (
	"github.com/gin-gonic/gin"
)

type HTTPMethod string

const (
	POST    HTTPMethod = "POST"
	GET     HTTPMethod = "GET"
	DELETE  HTTPMethod = "DELETE"
	PATCH   HTTPMethod = "PATCH"
	PUT     HTTPMethod = "PUT"
	OPTIONS HTTPMethod = "OPTIONS"
	HEAD    HTTPMethod = "HEAD"
	Any     HTTPMethod = "Any"
)

type group struct {
	Path string
}

type Route struct {
	Method    HTTPMethod
	GroupPath string
	Path      string
	Name      string
	Desc      string
}

type GinHandler struct {
	Opt Route
	// BHandler before binding
	BHandler gin.HandlersChain
	// BindParam would use c.ShouldBind()
	// AHandler after binding
	AHandler gin.HandlersChain
}

type R []interface{}
type M = gin.HandlersChain
