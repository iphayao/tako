package tako

import (
	"errors"
	"log"
	"net/http"
	"strings"
	"sync"
)

type MiddlewareFunc func(HandlerFunc) HandlerFunc

type HandlerFunc func(*Context) error

type PathParam struct {
	prefix string
	value  string
}

type Path struct {
	Path    string
	PathKey string
	Paths   []string
	Params  map[int]string
}

type Route struct {
	Method    string
	Path      Path
	Handlers  []HandlerFunc
	pathParam []string
	engine    *Engine
}

type Engine struct {
	pool       sync.Pool
	routes     map[string]Route
	middleware []HandlerFunc
}

func New() *Engine {
	e := &Engine{}

	e.routes = make(map[string]Route)
	e.middleware = make([]HandlerFunc, 0)
	e.pool.New = func() interface{} {
		return e.allocateContext()
	}

	return e
}

func (e *Engine) Context() *Context {
	return e.pool.Get().(*Context)
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c := e.Context()
	c.Update(r, w)

	if err := e.handleHTTPRequest(c); err != nil {
		w.WriteHeader(http.StatusNotFound)
	}

	e.pool.Put(c)
}

func (e *Engine) Use(middleware HandlerFunc) {
	e.middleware = append(e.middleware, middleware)
}

func (e *Engine) addRoute(method string, path string, h HandlerFunc) *Route {
	route := &Route{
		Path:   parsePath(path),
		Method: method,
		engine: e,
	}

	e.applyMiddleware(route)
	route.Handlers = append(route.Handlers, h)

	e.routes[method+route.PathKey()] = *route
	return route
}

func (e *Engine) applyMiddleware(r *Route) {
	for i := 0; i < len(e.middleware); i++ {
		r.Handlers = append(r.Handlers, e.middleware[i])
	}
}

func (e *Engine) Start(address string) error {
	log.Printf("HTTP server started on %s", address)
	return http.ListenAndServe(address, e)
}

func (e *Engine) handleHTTPRequest(c *Context) error {
	method := c.Request.Method
	path := c.Request.URL.Path

	pKey, paths := parseReqPath(path)
	if route, ok := e.routes[method+pKey]; ok {
		for k, v := range route.Path.Params {
			c.Params[v] = paths[k]
		}

		for i := 0; i < len(route.Handlers); i++ {
			if err := route.Handlers[i](c); err != nil {
				log.Printf("Error %s", err)
			}
		}
	} else {
		return errors.New("route not found")
	}

	return nil
}

func (e *Engine) allocateContext() *Context {
	return &Context{engine: e, KeysMutex: &sync.RWMutex{}}
}

// Setup handles for each Request
func (e *Engine) Handle(method string, path string, h HandlerFunc) *Route {
	return e.handler(method, path, h)
}

func (e *Engine) GET(path string, h HandlerFunc) *Route {
	return e.handler(http.MethodGet, path, h)
}

func (e *Engine) POST(path string, h HandlerFunc) *Route {
	return e.handler(http.MethodPost, path, h)
}

func (e *Engine) DELETE(path string, h HandlerFunc) *Route {
	return e.handler(http.MethodDelete, path, h)
}

func (e *Engine) PUT(path string, h HandlerFunc) *Route {
	return e.handler(http.MethodPut, path, h)
}

func (e *Engine) PATCH(path string, h HandlerFunc) *Route {
	return e.handler(http.MethodPatch, path, h)
}

func (e *Engine) handler(method string, path string, h HandlerFunc) *Route {
	return e.addRoute(method, path, h)
}

func (r *Route) PathKey() string {
	return r.Path.PathKey
}

func parseReqPath(path string) (string, []string) {
	pathKeys := ""
	paths := splitPath(path)

	for i := 0; i < len(paths); i++ {
		if i%2 == 0 {
			pathKeys += "/" + paths[i]
		} else {
			pathKeys += "/???"
		}
	}

	return pathKeys, paths
}

func parsePath(path string) Path {
	pathKeys := ""
	params := make(map[int]string)
	paths := splitPath(path)

	for i := 0; i < len(paths); i++ {
		if strings.HasPrefix(paths[i], ":") {
			pathKeys += "/???"
			params[i] = paths[i][1:]
		} else {
			pathKeys += "/" + paths[i]
		}
	}

	return Path{
		Path:    path,
		PathKey: pathKeys,
		Paths:   paths,
		Params:  params,
	}
}

func splitPath(path string) []string {
	paths := strings.Split(path, "/")[1:]
	if strings.HasSuffix(path, "/") {
		paths = paths[:len(paths)-1]
	}
	return paths
}