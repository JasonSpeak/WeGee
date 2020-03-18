package main

import (
	"gee"
	"net/http"
)

func main() {
	r := gee.New()
	r.GET("/index", func(c *gee.Context) {
		c.HTML(http.StatusOK, "<h1>Index page</h1>\n")
	})

	v1 := r.Group("/v1")
	{
		v1.GET("/", func(c *gee.Context) {
			c.HTML(http.StatusOK, "<h1>Hello fuck you V1</h1>")
		})
		v1.GET("/hello", func(c *gee.Context) {
			// expect /v1/hello?name=geektutu
			c.String(http.StatusOK, "v1: Fuck you %s,you're at %s and kiss my ass\n", c.Query("name"), c.Path)
		})
	}
	v2 := r.Group("/v2")
	{
		v2.GET("/hello/:name", func(c *gee.Context) {
			c.String(http.StatusOK, "V2: Fuck you %s, you are at my ass %s", c.Param("name"), c.Path)
		})
	}

	v2.POST("/login", func(c *gee.Context) {
		c.JSON(http.StatusOK, gee.H{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	})

	r.Run(":9999")
}
