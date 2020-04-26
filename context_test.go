package tako

import (
	"bytes"
	"io"
	"net/http"
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
	// act
	c.Update(mockHttpRequest(), mockHttpResponse())
	// assert
	if c.Request == nil {
		t.Error("Failed update HTTP request")
	}
	if c.Response == nil {
		t.Error("Failed update HTTP response")
	}
}

func TestUpdate_ExpectContextResponseWasUpdated(t *testing.T) {
	// arrange
	e := New()
	c := e.Context()
	// act
	c.Update(mockHttpRequest(), mockHttpResponse())
	// assert
	if c.Response == nil {
		t.Error("Failed update HTTP response")
	}
}

func TestString_ExpectHttpStatus200(t *testing.T) {
	// arrange
	e := New()
	c := e.Context()
	w := mockHttpResponse()
	c.Update(mockHttpRequest(), w)
	// act
	const TestingMessage = "TESTING MESSAGE"
	err := c.String(http.StatusOK, TestingMessage)
	// assert
	if err != nil && string(w.Data) == TestingMessage {
		t.Error("Failed write string response")
	}
}

func TestJSON_ExpectHttpStatus200(t *testing.T) {
	// arrange
	e := New()
	c := e.Context()
	w := mockHttpResponse()
	c.Update(mockHttpRequest(), w)
	// act
	err := c.JSON(http.StatusOK, TestModel{"TESTING MESSAGE"})
	// assert
	if err != nil || string(w.Data) == "{}" {
		t.Error("Failed write string response")
	}
}

func TestBind_ExpectNoError(t *testing.T) {
	// arrange
	e := New()
	c := e.Context()
	r := mockHttpRequest()
	r.Body = NewReader([]byte("{\n\"message\" : \"TEST_MESSAGE\"\n}"))
	c.Update(r, mockHttpResponse())
	m := &TestModel{}
	// act
	err := c.Bind(m)
	// assert
	if err != nil && m.Message != "TEST_MESSAGE" {
		t.Errorf("Failed write string response %s", err)
	}
}

func mockHttpRequest() *http.Request {
	return &http.Request{}
}

func mockHttpResponse() *TestResponseWriter {
	return &TestResponseWriter{}
}

type TestResponseWriter struct {
	Data []byte
}

func (t *TestResponseWriter) Write(bytes []byte) (int, error) {
	t.Data = bytes
	return len(t.Data), nil
}

func (t *TestResponseWriter) WriteHeader(statusCode int) {
	//return nil
}

func (t *TestResponseWriter) Header() http.Header {
	return http.Header{}
}

type TestReadCloser struct {
	Reader io.Reader
}

func NewReader(b []byte) *TestReadCloser{
	t := &TestReadCloser{}
	t.Reader = bytes.NewReader(b)

	return t
}

func (t TestReadCloser) Read(p []byte) (n int, err error) {
	return t.Reader.Read(p)
}

func (t TestReadCloser) Close() error {
	panic("implement me")
}

type TestModel struct {
	Message string
}

