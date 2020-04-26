package tako

import (
	"net/http"
	"testing"
)

func TestHandle_ExpectReturnRoute(t *testing.T) {
	// arrange
	e := New()
	path := "/tests"
	// act
	route := e.Handle(http.MethodGet, path, MockHandler())
	// assert
	if route == nil {
		t.Errorf("Failed create handle route %s", path)
	}
}

func TestGET_ExpectReturnRoute(t *testing.T) {
	// arrange
	e := New()
	// act
	route := e.GET("/tests", MockHandler())
	// assert
	if route == nil || route.Method != http.MethodGet {
		t.Error()
	}
}

func TestGET_ExpectCorrectRouteKey(t *testing.T) {
	// arrange
	e := New()
	// act
	e.GET("/tests", MockHandler())
	// assert
	if _, ok := e.routes["GET/tests"]; !ok {
		t.Error()
	}
}

func TestPOST_ExpectReturnRoute(t *testing.T) {
	// arrange
	e := New()
	// act
	route := e.POST("/tests", MockHandler())
	// assert
	if route == nil || route.Method != http.MethodPost {
		t.Error()
	}
}

func TestPOST_ExpectCorrectRouteKey(t *testing.T) {
	// arrange
	e := New()
	// act
	e.POST("/tests", MockHandler())
	// assert
	if _, ok := e.routes["POST/tests"]; !ok {
		t.Error()
	}
}

func TestDELETE_ExpectReturnRoute(t *testing.T) {
	// arrange
	e := New()
	// act
	route := e.DELETE("/tests", MockHandler())
	// assert
	if route == nil || route.Method != http.MethodDelete {
		t.Error()
	}
}

func TestDELETE_ExpectCorrectRouteKey(t *testing.T) {
	// arrange
	e := New()
	// act
	e.DELETE("/tests", MockHandler())
	// assert
	if _, ok := e.routes["DELETE/tests"]; !ok {
		t.Error()
	}
}

func TestPUT_ExpectReturnRoute(t *testing.T) {
	// arrange
	e := New()
	// act
	route := e.PUT("/tests", MockHandler())
	// assert
	if route == nil || route.Method != http.MethodPut {
		t.Error()
	}
}

func TestPUT_ExpectCorrectRouteKey(t *testing.T) {
	// arrange
	e := New()
	// act
	e.PUT("/tests", MockHandler())
	// assert
	if _, ok := e.routes["PUT/tests"]; !ok {
		t.Error()
	}
}

func TestPATCH_ExpectReturnRoute(t *testing.T) {
	// arrange
	e := New()
	// act
	route := e.PATCH("/tests", MockHandler())
	// assert
	if route == nil || route.Method != http.MethodPatch {
		t.Error()
	}
}

func TestPATCH_ExpectCorrectRouteKey(t *testing.T) {
	// arrange
	e := New()
	// act
	e.PATCH("/tests", MockHandler())
	// assert
	if _, ok := e.routes["PATCH/tests"]; !ok {
		t.Error()
	}
}
