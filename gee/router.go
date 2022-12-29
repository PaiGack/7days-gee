package gee

import (
	"log"
	"net/http"
	"reflect"
	"runtime"
)

type router struct {
	router map[string]HandlerFunc
}

func newRouter() *router {
	return &router{router: make(map[string]HandlerFunc)}
}

func (router *router) addRoute(mathod string, pattern string, handler HandlerFunc) {
	log.Printf("Route %4s - %20s %s", mathod, pattern, runtime.FuncForPC(reflect.ValueOf(handler).Pointer()).Name())
	key := mathod + "-" + pattern
	router.router[key] = handler
}

func (router *router) handle(c *Context) {
	key := c.Method + "-" + c.Path
	if handler, ok := router.router[key]; ok {
		log.Printf("%s %s", c.Method, c.Path)

		handler(c)
	} else {
		log.Printf("not found: %s %s", c.Method, c.Path)

		c.String(http.StatusNotFound, "404 NOT FOUND: %q\n", c.Path)
	}
}
