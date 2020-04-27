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
		t.Error()
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
		t.Error()
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
		t.Error(err)
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
		t.Error(err)
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
		t.Error(err)
	}
}

func TestSetStatus_ExpectNoError(t *testing.T) {
	// arrange
	e := New()
	c := e.Context()
	req := httptest.NewRequest(http.MethodGet, "/tests", bytes.NewReader([]byte("")))
	res := httptest.NewRecorder()
	c.Update(req, res)
	// act
	err := c.SetStatus(http.StatusBadRequest)
	// assert
	if err != nil {
		t.Error(err)
	}
}

type TestModel struct {
	Message string
}
