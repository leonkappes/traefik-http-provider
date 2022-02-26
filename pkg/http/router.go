package http

import (
	"log"

	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

type Middleware func(ctx *fasthttp.RequestCtx) (int, error)

type Router struct {
	router     *router.Router
	middleware []Middleware
}

func New() *Router {
	r := router.New()
	return &Router{
		router:     r,
		middleware: make([]Middleware, 0),
	}
}

func (router *Router) AddRoute(route *Route) {
	router.router.Handle(route.Method, route.Path, router.handle(route.Handler))
}

func (router *Router) handle(handler func(ctx *fasthttp.RequestCtx)) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		for _, middleware := range router.middleware {
			if status, err := middleware(ctx); err != nil {
				log.Printf("Error in middleware: %d ; %s", status, err)
				return
			}
		}
		handler(ctx)
	}
}

func (router *Router) Use(middleware ...Middleware) {
	router.middleware = append(router.middleware, middleware...)
}

func (router *Router) Handler(ctx *fasthttp.RequestCtx) {
	router.router.Handler(ctx)
}
