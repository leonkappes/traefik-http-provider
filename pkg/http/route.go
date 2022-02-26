package http

import "github.com/valyala/fasthttp"

type Route struct {
	Method  string
	Path    string
	Handler func(ctx *fasthttp.RequestCtx)
}

func NewRoute(method, path string, handler func(ctx *fasthttp.RequestCtx)) *Route {
	return &Route{
		Method:  method,
		Path:    path,
		Handler: handler,
	}
}
