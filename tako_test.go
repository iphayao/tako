package tako

import "testing"

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

func TestAddRoute_GET_ExpectAddedRoute(t *testing.T) {
	// arrange
	e := New()
	// act
	e.GET("/tests", MockHandler())
	// assert
	if len(e.routes) == 0 {
		t.Error("Failed add route GET method")
	}
}

func TestAddRoute_GET_ExpectCorrectRouteKey(t *testing.T) {
	// arrange
	e := New()
	// act
	e.GET("/tests", MockHandler())
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

func MockHandler() HandlerFunc {
	return func(c *Context) error {
		return nil
	}
}

func MockMiddleware() HandlerFunc {
	return func(c *Context) error {
		return nil
	}
}
