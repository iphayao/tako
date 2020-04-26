package tako

import (
	"bytes"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNewEngine_ExpectEngineNotNull(t *testing.T) {
	// arrange & act
	e := New()
	// assert
	if e == nil {
		t.Error("Failed create new engine")
	}
}

func TestUse_ExpectAddedMiddleware(t *testing.T) {
	// arrange
	e := New()
	// act
	e.Use(MockMiddleware())
	// assert
	if len(e.middleware) == 0 {
		t.Error("Failed add middleware to engine")
	}
}

func TestAddRoute_ExpectAddedRoute(t *testing.T) {
	// arrange
	e := New()
	// act
	e.addRoute(http.MethodGet, "/tests", MockHandler())
	// assert
	if len(e.routes) == 0 {
		t.Error("Failed add route GET method")
	}
}

func TestAddRoute_ExpectCorrectRouteKey(t *testing.T) {
	// arrange
	e := New()
	// act
	e.addRoute(http.MethodGet, "/tests", MockHandler())
	// assert
	if _, ok := e.routes["GET/tests"]; !ok {
		t.Error("Failed route key not found")
	}
}

func TestAddRoute_AfterAddedMiddleware_Expect2Handlers(t *testing.T) {
	// arrange
	e := New()
	e.Use(MockMiddleware())
	// act
	e.GET("/tests", MockHandler())
	// assert
	r, _ := e.routes["GET/tests"]
	if len(r.Handlers) != 2 {
		t.Error("Failed add routed after middleware")
	}
}

func TestHandleHttpRequest_ExpectNotReturnError(t *testing.T) {
	// arrange
	e := New()
	e.GET("/tests", MockHandler())
	c := MockContext(http.MethodGet, "/tests", "")
	// act
	err := e.handleHTTPRequest(&c)
	// assert
	if err != nil {
		t.Error(err)
	}
}

func TestHandleHttpRequestWithId_ExpectNotReturnError(t *testing.T) {
	// arrange
	e := New()
	e.GET("/tests/:id", MockHandler())
	c := MockContext(http.MethodGet, "/tests/1234", "")
	// act
	err := e.handleHTTPRequest(&c)
	// assert
	if err != nil {
		t.Error(err)
	}
}

func TestHandleHttpRequestWithId2_ExpectNotReturnError(t *testing.T) {
	// arrange
	e := New()
	e.GET("/tests/:id", MockHandler())
	c := MockContext(http.MethodGet, "/tests/1234/", "")
	// act
	err := e.handleHTTPRequest(&c)
	// assert
	if err != nil {
		t.Error(err)
	}
}

func TestHandleHttpRequest_ExpectReturnError(t *testing.T) {
	// arrange
	e := New()
	c := MockContext(http.MethodGet, "/tests", "")
	// act
	err := e.handleHTTPRequest(&c)
	// assert
	if err == nil {
		t.Error(err)
	}
}

func TestHandleHttpRequest_ExpectError(t *testing.T) {
	// arrange
	e := New()
	e.GET("/tests", MockHandlerError())
	c := MockContext(http.MethodGet, "/tests", "")
	// act
	err := e.handleHTTPRequest(&c)
	// assert
	if err != nil {
		t.Error(err)
	}
}

func TestServeHTTP_ExpectResponseStatusOK(t *testing.T) {
	// arrange
	e := New()
	e.GET("/tests", MockHandler())
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/tests", bytes.NewReader(nil))
	// act
	e.ServeHTTP(w, r)
	// assert
	if w.Code != http.StatusOK {
		t.Error(w.Code)
	}
}

func TestServeHTTP_ExpectResponseStatusNotFound(t *testing.T) {
	// arrange
	e := New()
	e.GET("/tests", MockHandler())
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, "/tests", bytes.NewReader(nil))
	// act
	e.ServeHTTP(w, r)
	// assert
	if w.Code != http.StatusNotFound {
		t.Error(w.Code)
	}
}


//func TestStart_ExpectNoError(t *testing.T) {
//	// arrange
//	e := New()
//	// act
//	err := e.Start(":1234")
//	// assert
//	if err != nil {
//		t.Error(err)
//	}
//}

func MockContext(method string, path string, body string) Context {
	c := Context{}
	c.Request = httptest.NewRequest(method, path, bytes.NewReader([]byte(body)))
	c.Response = httptest.NewRecorder()
	c.Method = c.Request.Method
	c.Headers = c.Request.Header
	c.Params = make(map[string]string)

	return c
}

func MockHandler() HandlerFunc {
	return func(c *Context) error {
		return nil
	}
}

func MockHandlerError() HandlerFunc {
	return func(c *Context) error {
		return errors.New("mock error")
	}
}

func MockMiddleware() HandlerFunc {
	return func(c *Context) error {
		return nil
	}
}
