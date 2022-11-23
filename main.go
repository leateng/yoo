package main

import (
	"net/http"
	"yoo"
)

func main() {
	engine := yoo.NewEngine()

	engine.GET("/", func(c *yoo.Context) {
		c.String(http.StatusOK, "hello yoo")
	})

	engine.GET("/html", func(c *yoo.Context) {
		c.HTML(http.StatusOK, "<h1>hello html</h1>")
	})

	engine.GET("/json", func(c *yoo.Context) {
		c.JSON(http.StatusOK, yoo.H{
			"hello": "world",
			"query": c.Query("query"),
		})
	})

	engine.Run(":3000")
}
