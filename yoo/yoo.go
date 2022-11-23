package yoo

import (
	"net/http"
)

type HandlerFunc func(*Context)

type Engine struct {
  // router map[string]HandlerFunc
  router *router
}

func NewEngine() *Engine {
  return &Engine{router: newRouter()}
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request)  {
  c := newContext(w, r)
  engine.router.Handle(c)
}

func (engine *Engine) Run(addr string)  {
  http.ListenAndServe(addr, engine)
}

func (engine *Engine) addRoute(method string, pattern string, handler HandlerFunc)  {
  engine.router.AddRoute(method, pattern, handler)
}

func (engine *Engine) GET(path string, handler HandlerFunc)  {
  engine.addRoute("GET", path, handler)
}

func (engine *Engine) POST(path string, handler HandlerFunc)  {
  engine.addRoute("POST", path, handler)
}


