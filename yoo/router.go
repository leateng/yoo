package yoo

import (
	"log"
	"net/http"
)

type router struct {
  handlers map[string]HandlerFunc
}

func newRouter() *router {
  return &router{ handlers: make(map[string]HandlerFunc)}
}

func (r *router) AddRoute(method string, pattern string, handler HandlerFunc)  {
  log.Printf("Route %4s - %s", method, pattern)

  key := method + "-" + pattern
  r.handlers[key] = handler
}

func (r *router) Handle(c *Context) {
  router_key := c.Req.Method + "-" + c.Req.URL.Path
  if handler, ok := r.handlers[router_key]; ok {
    handler(c)
  } else {
    c.String(http.StatusNotFound, "404 Page not found")
  }
}

