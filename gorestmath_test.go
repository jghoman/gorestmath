package gorestmath

import (
	"bytes"
	"fmt"
	"net/http"
	"net/url"
	"testing"
)

type MockResponseWriter struct {
	writtenBytes  []byte
	writtenHeader int
}

func (*MockResponseWriter) Header() http.Header {
	fmt.Println("hi")
	return nil
}

func (mrw *MockResponseWriter) Write(b []byte) (int, error) {
	mrw.writtenBytes = b
	return 0, nil
}

func (mrw *MockResponseWriter) WriteHeader(header int) {
	mrw.writtenHeader = header
}

func assertByteArrayEquals(expected []byte, actual []byte, t *testing.T) {
	if bytes.Compare(expected, actual) != 0 {
		t.Errorf("Expected %v but got %v. Failing", expected, actual)
	}
}

func assertIntEquals(expected, actual int, t *testing.T) {
	if expected != actual {
		t.Errorf("Expected %v but got %v. Failing", expected, actual)
	}
}

func assertStringEquals(expected, actual string, t *testing.T) {
	if expected != actual {
		t.Errorf("Expected %v but got %v. Failing", expected, actual)
	}
}

func getUrl(s string, t *testing.T) *url.URL {
	url, err := url.Parse(s)
	if err != nil {
		t.Errorf("Somehow couldn't parse the url.")
	}

	return url
}

func TestAdding1and2(t *testing.T) {
	t.Parallel()
	mrw := &MockResponseWriter{}

	url := getUrl("http://www.hello.com/add/1/2", t)

	request := &http.Request{URL: url}
	DoSomeMath(mrw, request)

	assertByteArrayEquals([]byte(`{'result':'3'}`), mrw.writtenBytes, t)
}

func TestBadPath(t *testing.T) {
	t.Parallel()
	mrw := &MockResponseWriter{}
	url := getUrl("http://www.hello.com/xxi/fsa/add/1/2///", t)
	request := &http.Request{URL: url}
	DoSomeMath(mrw, request)

	assertIntEquals(http.StatusBadRequest, mrw.writtenHeader, t)
	assertByteArrayEquals([]byte(ParseError), mrw.writtenBytes, t)
}
