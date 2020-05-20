package routex

import (
	"github.com/gin-gonic/gin"
	"github.com/go-web-kits/utils"
)

var NameToRouteMap map[string]Route
var CurrentGroupPath string

func init() {
	NameToRouteMap = make(map[string]Route)
}

func Group(path string, f func() R) []interface{} {
	oldGroupPath := CurrentGroupPath
	CurrentGroupPath += path
	routes := f()
	g := []interface{}{group{path}}
	for _, route := range routes {
		g = append(g, route)
	}
	CurrentGroupPath = oldGroupPath
	return g
}

func API(args ...string) Route {
	if len(args) == 1 {
		return Route{Name: args[0]}
	} else if len(args) == 2 {
		return Route{Name: args[0], Desc: args[1]}
	}

	return Route{}
}

func (api Route) Action(method HTTPMethod, path string, handler gin.HandlerFunc) R {
	api.Method = method
	api.Path = path
	api.GroupPath = CurrentGroupPath
	if api.Name == "" {
		api.Name = utils.GetFuncName(handler)
	}
	NameToRouteMap[api.Name] = api

	return append(R{api}, handler)
}

func (r R) Middleware(middleware ...interface{}) R {
	api, handler := r[0], r[1]
	r = append(R{api}, middleware...)
	return append(r, handler)
}

func (api Route) GET(path string, handler gin.HandlerFunc) R {
	return api.Action(GET, path, handler)
}

func (api Route) POST(path string, handler gin.HandlerFunc) R {
	return api.Action(POST, path, handler)
}

func (api Route) DELETE(path string, handler gin.HandlerFunc) R {
	return api.Action(DELETE, path, handler)
}

func (api Route) PUT(path string, handler gin.HandlerFunc) R {
	return api.Action(PUT, path, handler)
}

func GETx(path string, handler gin.HandlerFunc) R {
	return Route{}.Action(GET, path, handler)
}

func POSTx(path string, handler gin.HandlerFunc) R {
	return Route{}.Action(POST, path, handler)
}

func DELETEx(path string, handler gin.HandlerFunc) R {
	return Route{}.Action(DELETE, path, handler)
}

func PUTx(path string, handler gin.HandlerFunc) R {
	return Route{}.Action(PUT, path, handler)
}
