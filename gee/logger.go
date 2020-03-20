package gee

import (
	"log"
	"time"
)

//Logger is a log Middleware
func Logger() HandlerFunc {
	return func(c *Context) {
		t := time.Now()

		c.Next()

		log.Printf("[%d] %s in %v", c.StatusCode, c.Req.RequestURI, time.Since(t))
	}
}
