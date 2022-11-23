package yoo

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type H map[string]interface{}

type Context struct {
	Writer    http.ResponseWriter
	Req       *http.Request
	Method    string
	Path      string
	StatusCode int
}

func newContext(w http.ResponseWriter, r *http.Request) *Context {
	return &Context{
		Writer: w,
		Req:    r,
		Method: r.Method,
		Path:   r.URL.Path,
	}
}

func (c *Context) PostForm(key string) string {
  return c.Req.FormValue(key)
}

func (c *Context) Query(key string) string {
  return c.Req.URL.Query().Get(key)
}

func (c *Context) Status(status int)  {
  c.StatusCode = status
  c.Writer.WriteHeader(status)
}

func (c *Context) SetHeader(key string, value string)  {
  c.Writer.Header().Set(key, value)
}

func (c *Context) String(status int, format string, values ...interface{})  {
  c.SetHeader("Content-Type", "text/plain")
  c.Status(status)
  c.Writer.Write([]byte(fmt.Sprintf(format, values...)))
}

func (c *Context) Data(status int, data []byte)  {
  c.SetHeader("Content-Type", "text/plain")
  c.Status(status)
  c.Writer.Write(data)
}

func (c *Context) JSON(status int, obj interface{})  {
  c.SetHeader("Content-Type", "application/json")
  c.Status(status)

  encoder := json.NewEncoder(c.Writer)
  if err := encoder.Encode(obj); err != nil {
    http.Error(c.Writer, err.Error(), 500)
  }
}

func (c *Context) HTML(status int, html string)  {
  c.SetHeader("Content-Type", "text/html")
  c.Status(status)
  c.Writer.Write([]byte(html))
}

