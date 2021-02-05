package contexts

import (
	"net/http/httptest"
	"testing"
)

func TestNewGoliathContext(t *testing.T) {
	r := httptest.NewRequest("GET", "/", nil)
	got := NewGoliathContext(r)
	if got == nil {
		t.Errorf("NewGoliathContext() want something, got nil")
	}
}

func Test_goliathContext_Auth(t *testing.T) {
	r := httptest.NewRequest("GET", "/", nil)
	ctx1 := NewGoliathContext(r)

	auth1 := ctx1.Auth()
	auth2 := ctx1.Auth()

	if auth1 != auth2 {
		t.Errorf("Auth() want same authContext object, got different")
	}

	ctx2 := NewGoliathContext(r)
	auth3 := ctx2.Auth()
	if auth3 == auth1 {
		t.Errorf("Auth() want different authContext object, got same")
	}
}
