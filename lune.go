package lune

import (
	"log/slog"
	"net/http"
)

// HandlerFunc defines the request handler used by lune
type HandlerFunc func(http.ResponseWriter, *http.Request)

// Engine implement the interface of ServeHTTP
type Engine struct {
	router map[string]HandlerFunc
}

/*
*
通过查看net/http的源码可以发现，Handler是一个接口，需要实现方法 ServeHTTP ，
也就是说，只要传入任何实现了 ServerHTTP 接口的实例，所有的HTTP请求，就都交给了该实例处理了。
*/
func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	routerKey := req.Method + "-" + req.URL.Path
	if handlerFunc, ok := engine.router[routerKey]; ok {
		handlerFunc(w, req)
	} else {
		w.WriteHeader(http.StatusNotFound)
		slog.Error("404 NOT FOUND : %s", req.URL)
	}
}

// New the constructor of lune.Engine
func New() *Engine {
	return &Engine{router: make(map[string]HandlerFunc)}
}

func (engine *Engine) addRoute(method string, pattern string, handler HandlerFunc) {
	routerKey := method + "-" + pattern
	engine.router[routerKey] = handler
}

// GET defines the method to add GET request
func (engine *Engine) GET(pattern string, handler HandlerFunc) {
	engine.addRoute("GET", pattern, handler)
}

// POST defines the method to add POST request
func (engine *Engine) POST(pattern string, handler HandlerFunc) {
	engine.addRoute("POST", pattern, handler)
}

func (engine *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, engine)
}
