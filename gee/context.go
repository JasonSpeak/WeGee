package gee

import (
	"encoding/json"
	"fmt"
	"net/http"
)

//H defines a nick name for Json data
type H map[string]interface{}

//Context defines a data struct for gee http link
type Context struct {
	//http origin object
	Writer http.ResponseWriter
	Req    *http.Request
	//request info
	Path   string
	Method string
	Params map[string]string
	//response info
	StatusCode int
}

func newContext(w http.ResponseWriter, req *http.Request) *Context {
	return &Context{
		Writer: w,
		Req:    req,
		Path:   req.URL.Path,
		Method: req.Method,
	}
}

//PostForm used to get values from forms
func (c *Context) PostForm(key string) string {
	return c.Req.FormValue(key)
}

//Query used to get values from url query
func (c *Context) Query(key string) string {
	return c.Req.URL.Query().Get(key)
}

//Status used to set Status of current Context
func (c *Context) Status(code int) {
	c.StatusCode = code
	c.Writer.WriteHeader(code)
}

//SetHeader used to set current context's header
func (c *Context) SetHeader(key, value string) {
	c.Writer.Header().Set(key, value)
}

//String used to return  string format of values
func (c *Context) String(code int, format string, values ...interface{}) {
	c.SetHeader("Content-Type", "text/plain")
	c.Status(code)
	c.Writer.Write([]byte(fmt.Sprintf(format, values...)))
}

//JSON used to encode obj to json type
func (c *Context) JSON(code int, obj interface{}) {
	c.SetHeader("Content-Type", "application/json")
	c.Status(code)
	encoder := json.NewEncoder(c.Writer)
	if err := encoder.Encode(obj); err != nil {
		http.Error(c.Writer, err.Error(), 500)
	}
}

//Data used to set data into current context
func (c *Context) Data(code int, data []byte) {
	c.Status(code)
	c.Writer.Write(data)
}

//HTML used to set html into current context
func (c *Context) HTML(code int, html string) {
	c.SetHeader("Content-Type", "text/html")
	c.Status(code)
	c.Writer.Write([]byte(html))
}

//Param used to get route parameter of current context
func (c *Context) Param(key string) string {
	value, _ := c.Params[key]
	return value
}
