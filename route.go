package tako

import "net/http"

type Router struct {
	engine *Engine
}

type Route struct {
	Method    string
	Path      Path
	Handlers  []HandlerFunc
	pathParam []string
	engine    *Engine
}

// Setup handles for each Request
func (r *Router) Handle(method string, path string, h HandlerFunc) *Route {
	return r.handler(method, path, h)
}

func (r *Router) GET(path string, h HandlerFunc) *Route {
	return r.handler(http.MethodGet, path, h)
}

func (r *Router) POST(path string, h HandlerFunc) *Route {
	return r.handler(http.MethodPost, path, h)
}

func (r *Router) DELETE(path string, h HandlerFunc) *Route {
	return r.handler(http.MethodDelete, path, h)
}

func (r *Router) PUT(path string, h HandlerFunc) *Route {
	return r.handler(http.MethodPut, path, h)
}

func (r *Router) PATCH(path string, h HandlerFunc) *Route {
	return r.handler(http.MethodPatch, path, h)
}

func (r *Router) handler(method string, path string, h HandlerFunc) *Route {
	return r.engine.addRoute(method, path, h)
}
