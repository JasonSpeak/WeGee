package gee

import (
	"log"
	"net/http"
)

//HandlerFunc defines the request handler used by gee
type HandlerFunc func(*Context)

//RouterGroup defines group routers
type RouterGroup struct {
	prefix      string        //Group route
	middlewares []HandlerFunc //Group middlewares
	parent      *RouterGroup  //parent of current group
	engine      *Engine       //engine of current group
}

//Engine implement the interface of ServeHTTP
type Engine struct {
	*RouterGroup                //inherit all functons of RouterGroup
	router       *router        //router tree of this engine
	groups       []*RouterGroup //All groups of this engine
}

//New is the constructor of gee.Engine
func New() *Engine {
	engine := &Engine{router: newRouter()}
	engine.RouterGroup = &RouterGroup{engine: engine}
	engine.groups = []*RouterGroup{engine.RouterGroup}
	return engine
}

//Group is used to create a new RouterGroup
//All groups share the same engine instance
func (group *RouterGroup) Group(prefix string) *RouterGroup {
	engine := group.engine
	newGroup := &RouterGroup{
		prefix: group.prefix + prefix,
		parent: group,
		engine: engine,
	}

	engine.groups = append(engine.groups, newGroup)
	return newGroup
}

func (group *RouterGroup) addRoute(method string, comp string, handler HandlerFunc) {
	pattern := group.prefix + comp
	log.Printf("Route %4s - %s", method, pattern)
	group.engine.router.addRoute(method, pattern, handler)
}

//GET defines the method to add GET request
func (group *RouterGroup) GET(pattern string, handler HandlerFunc) {
	group.addRoute("GET", pattern, handler)
}

//POST defines the method to add POST request
func (group *RouterGroup) POST(pattern string, handler HandlerFunc) {
	group.addRoute("POST", pattern, handler)
}

//Run defines the method to start a http server
func (engine *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, engine)
}

//ServeHTTP implement the request handler
func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	c := newContext(w, req)
	engine.router.handle(c)
}
