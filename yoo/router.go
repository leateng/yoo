package yoo

import (
	"log"
	"net/http"
	"strings"
)

type router struct {
	roots    map[string]*node
	handlers map[string]HandlerFunc
}

func newRouter() *router {
	return &router{
		roots:    make(map[string]*node),
		handlers: make(map[string]HandlerFunc),
	}
}

func parsePattern(pattern string) []string {
	splits := strings.Split(pattern, "/")
	parts := make([]string, 0)

	for _, part := range splits {
		if part != "" {
			parts = append(parts, part)

			if part[0] == '*' {
				break
			}
		}
	}

	return parts
}

func (r *router) AddRoute(method string, pattern string, handler HandlerFunc) {
	log.Printf("Route %4s - %s", method, pattern)

  parts := parsePattern(pattern)
  if _, ok := r.roots[method]; !ok {
    r.roots[method] = &node{}
  }
  r.roots[method].insert(pattern, parts, 0)

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
