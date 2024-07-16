package router

import (
	"fmt"
	"net/http"
)

type middleware func(http.Handler) http.Handler

type BaseRouter struct {
	Prefix           string
	GlobalMiddleware []middleware
	middleware       []middleware
	Routes           map[string]http.Handler
}

func New(prefix string) *BaseRouter {
	return &BaseRouter{
		Prefix: prefix,
	}
}

func (base *BaseRouter) register(method string, path string, handler http.Handler) {
	route := fmt.Sprintf("%s %s%s", method, base.Prefix, path)

	var mergeHandler = handler
	for i := range base.GlobalMiddleware {
		mergeHandler = base.GlobalMiddleware[i](mergeHandler)
	}

	for i := range base.middleware {
		mergeHandler = base.middleware[i](mergeHandler)
	}

	base.Routes[route] = mergeHandler
}

func (base *BaseRouter) Use(middleware ...middleware) {
	base.middleware = append(base.middleware, middleware...)
}

func (base *BaseRouter) Get(path string, handler http.Handler) {
	base.register("GET", path, handler)
}

func (base *BaseRouter) Post(path string, handler http.Handler) {
	base.register("POST", path, handler)
}

func (base *BaseRouter) Patch(path string, handler http.Handler) {
	base.register("PATCH", path, handler)
}

func (base *BaseRouter) Put(path string, handler http.Handler) {
	base.register("PUT", path, handler)
}

func (base *BaseRouter) Delete(path string, handler http.Handler) {
	base.register("DELETE", path, handler)
}
