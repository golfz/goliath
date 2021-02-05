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

	auth1 := ctx1.AuthContext()
	auth2 := ctx1.AuthContext()

	if auth1 != auth2 {
		t.Errorf("AuthContext() want same authContext object, got different")
	}

	ctx2 := NewGoliathContext(r)
	auth3 := ctx2.AuthContext()
	if auth3 == auth1 {
		t.Errorf("AuthContext() want different authContext object, got same")
	}
}
