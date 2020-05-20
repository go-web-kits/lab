package routex

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func RegisterRouteX(engine *gin.Engine, routes []interface{}) {
	register(&engine.RouterGroup, routes)
}

func register(r *gin.RouterGroup, routes []interface{}) {
	var grouping *gin.RouterGroup
	if g, ok := routes[0].(group); ok {
		grouping = r.Group(g.Path)
		defer func(g group) {
			if err := recover(); err != nil {
				err = fmt.Errorf("config path: %v. error: %v", g.Path, err)
				return
			}
		}(g)
	} else {
		grouping = r
	}

	for _, route := range routes {
		switch elem := route.(type) {
		case group:
		case R:
			registerHandler(grouping, getRouter(elem))
		case M:
			grouping.Use(elem...)
		case []interface{}:
			register(grouping, elem)
		default:
			panic(fmt.Errorf("unsupport item. type: %t", elem))
		}
	}
}

func registerHandler(r *gin.RouterGroup, h *GinHandler) {
	var handlers gin.HandlersChain
	handlers = append(handlers, h.BHandler...)
	handlers = append(handlers, h.AHandler...)
	http(r, h.Opt.Method, h.Opt.Path, handlers...)
}

func getRouter(h R) *GinHandler {
	handler := &GinHandler{
		Opt: Route{
			Method: GET,
			Path:   "/",
		},
	}

	for _, i := range h {
		switch elem := i.(type) {
		case Route:
			handler.Opt = elem
		case gin.HandlersChain:
			handler.AHandler = append(handler.AHandler, elem...)
		case gin.HandlerFunc:
			handler.BHandler = append(handler.BHandler, elem)
		case func(*gin.Context):
			handler.AHandler = append(handler.AHandler, elem)
		default:
			panic(fmt.Errorf("unsupport elem. type: %t", elem))
		}
	}

	return handler
}

func http(r *gin.RouterGroup, method HTTPMethod, relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes {
	var f func(relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes
	switch method {
	case POST:
		f = r.POST
	case GET:
		f = r.GET
	case DELETE:
		f = r.DELETE
	case PATCH:
		f = r.PATCH
	case PUT:
		f = r.PUT
	case OPTIONS:
		f = r.OPTIONS
	case HEAD:
		f = r.HEAD
	case Any:
		f = r.Any
	default:
		panic("unsupported method " + method)
	}
	return f(relativePath, handlers...)
}
