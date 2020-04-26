package tako

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetContext_ExpectContextNotNull(t *testing.T) {
	// arrange
	e := New()
	// act
	c := e.Context()
	// assert
	if c == nil {
		t.Error("Failed create context from engine")
	}
}

func TestUpdate_ExpectContextRequestWasUpdated(t *testing.T) {
	// arrange
	e := New()
	c := e.Context()
	req := httptest.NewRequest(http.MethodGet, "/tests", bytes.NewReader([]byte("")))
	res := httptest.NewRecorder()
	// act
	c.Update(req, res)
	// assert
	if c.Request == nil {
		t.Error("Failed update HTTP req")
	}
	if c.Response == nil {
		t.Error("Failed update HTTP res")
	}
}

func TestUpdate_ExpectContextResponseWasUpdated(t *testing.T) {
	// arrange
	e := New()
	c := e.Context()
	req := httptest.NewRequest(http.MethodGet, "/tests", bytes.NewReader([]byte("")))
	res := httptest.NewRecorder()
	// act
	c.Update(req, res)
	// assert
	if c.Response == nil {
		t.Error("Failed update HTTP res")
	}
}

func TestString_ExpectHttpStatus200(t *testing.T) {
	// arrange
	e := New()
	c := e.Context()
	req := httptest.NewRequest(http.MethodGet, "/tests", bytes.NewReader([]byte("")))
	res := httptest.NewRecorder()
	c.Update(req, res)
	// act
	const TestingMessage = "TESTING MESSAGE"
	err := c.String(http.StatusOK, TestingMessage)
	// assert
	if err != nil && res.Body.String() == TestingMessage {
		t.Error("Failed write string response")
	}
}

func TestJSON_ExpectHttpStatus200(t *testing.T) {
	// arrange
	e := New()
	c := e.Context()
	req := httptest.NewRequest(http.MethodGet, "/tests", bytes.NewReader([]byte("")))
	res := httptest.NewRecorder()
	c.Update(req, res)
	// act
	err := c.JSON(http.StatusOK, TestModel{"TESTING MESSAGE"})
	// assert
	if err != nil || res.Body.String() == "{}" {
		t.Error("Failed write string response")
	}
}

func TestBind_ExpectNoError(t *testing.T) {
	// arrange
	e := New()
	c := e.Context()
	req := httptest.NewRequest(http.MethodGet, "/tests", bytes.NewReader([]byte("{\n\"message\" : \"TEST_MESSAGE\"\n}")))
	res := httptest.NewRecorder()
	c.Update(req, res)
	m := &TestModel{}
	// act
	err := c.Bind(m)
	// assert
	if err != nil && m.Message != "TEST_MESSAGE" {
		t.Errorf("Failed write string response %s", err)
	}
}

type TestModel struct {
	Message string
}

