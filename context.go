package tako

import (
	"encoding/json"
	"net/http"
	"sync"
)

const HeaderContentType = "Content-Type"

type Context struct {
	Request   *http.Request
	Response  http.ResponseWriter
	Path      string
	Method    string
	Params    map[string]string
	Headers   http.Header
	engine    *Engine
	KeysMutex *sync.RWMutex
}

func (c *Context) JSON(code int, value interface{}) error {
	c.writeContentType("application/json")
	return c.Render(code, value)
}

func (c *Context) String(status int, message string) error {
	return c.JSON(status, message)
}

func (c *Context) Render(code int, value interface{}) error {
	enc := json.NewEncoder(c.Response)
	enc.SetIndent("", "  ")

	c.Status(code)

	return enc.Encode(value)
}

func (c *Context) Status(code int) {
	c.Response.WriteHeader(code)
}

func (c *Context) writeContentType(value string) {
	header := c.Response.Header()

	if header.Get(HeaderContentType) == "" {
		header.Set(HeaderContentType, value)
	}
}

func (c *Context) update(r *http.Request, w http.ResponseWriter) {
	c.Request = r
	c.Response = w
	c.Method = r.Method
	c.Headers = r.Header
	c.Params = make(map[string]string)
}

func (c *Context) reset() {
	c.Request = nil
	c.Response = nil
	c.Method = ""
}