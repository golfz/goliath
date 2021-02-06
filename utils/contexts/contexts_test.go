package contexts

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_GetResponseWriter(t *testing.T) {
	h := http.Header{}
	h.Set("Key-test", "value-test")
	w := mockWriter{
		header: &h,
	}
	ctx := GoliathContext{
		Writer: w,
	}
	got := ctx.GetResponseWriter()
	if got.Header().Get("Key-test") != "value-test" {
		t.Errorf("GetResponseWriter() expect same w value, got different")
	}
}

func Test_GetRequest(t *testing.T) {
	r := httptest.NewRequest("GET", "/expect", nil)
	ctx := GoliathContext{
		Request: r,
	}

	got := ctx.GetRequest()
	if got != r {
		t.Errorf("GetRequest() expect same r, got different")
	}
}

func Test_goliathContext_AuthContext(t *testing.T) {
	r := httptest.NewRequest("GET", "/", nil)
	ctx1 := GoliathContext{
		Request: r,
	}

	auth1 := ctx1.GetAuthContext()
	auth2 := ctx1.GetAuthContext()

	if auth1 != auth2 {
		t.Errorf("GetAuthContext() want same authContext object, got different")
	}

	ctx2 := GoliathContext{
		Request: r,
	}
	auth3 := ctx2.GetAuthContext()
	if auth3 == auth1 {
		t.Errorf("GetAuthContext() want different authContext object, got same")
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
