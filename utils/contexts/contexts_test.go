package contexts

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNewGoliathContext(t *testing.T) {
	r := httptest.NewRequest("GET", "/", nil)
	got := NewGoliathContext(nil, r)
	if got == nil {
		t.Errorf("NewGoliathContext() want something, got nil")
	}
}

func Test_GetResponseWriter(t *testing.T) {
	h := http.Header{}
	h.Set("Key-test", "value-test")
	mc := mockWriter{
		header: &h,
	}
	ctx := NewGoliathContext(mc, nil)
	got := ctx.GetResponseWriter()
	if got.Header().Get("Key-test") != "value-test" {
		t.Errorf("GetResponseWriter() expect same w value, got different")
	}
}

func Test_GetRequest(t *testing.T) {
	r := httptest.NewRequest("GET", "/expect", nil)
	ctx := NewGoliathContext(nil, r)

	got := ctx.GetRequest()
	if got != r {
		t.Errorf("GetRequest() expect same r, got different")
	}
}

func Test_goliathContext_AuthContext(t *testing.T) {
	r := httptest.NewRequest("GET", "/", nil)
	ctx1 := NewGoliathContext(nil, r)

	auth1 := ctx1.AuthContext()
	auth2 := ctx1.AuthContext()

	if auth1 != auth2 {
		t.Errorf("AuthContext() want same authContext object, got different")
	}

	ctx2 := NewGoliathContext(nil, r)
	auth3 := ctx2.AuthContext()
	if auth3 == auth1 {
		t.Errorf("AuthContext() want different authContext object, got same")
	}
}

type mockWriter struct {
	header *http.Header
}

func (mc mockWriter) Header() http.Header {
	return *mc.header
}

func (mc mockWriter) Write([]byte) (int, error) {
	return 0, nil
}

func (mc mockWriter) WriteHeader(statusCode int) {
}
