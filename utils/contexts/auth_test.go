package contexts

import (
	"github.com/golfz/goliath/cleanarch/data/output"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func Test_authContext_GetAuthToken(t *testing.T) {
	reqNoAuth := httptest.NewRequest("GET", "/", nil)

	reqSpaceStringAuth := httptest.NewRequest("GET", "/", nil)
	reqSpaceStringAuth.Header.Set("Authorization", "  ")

	reqWithAuth := httptest.NewRequest("GET", "/", nil)
	reqWithAuth.Header.Set("Authorization", "this_is_auth")

	type context struct {
		request *http.Request
	}
	tests := []struct {
		name      string
		ctx       context
		auth      string
		expectErr bool
	}{
		{
			name:      "no authorization, expect error",
			ctx:       context{request: reqNoAuth},
			auth:      "",
			expectErr: true,
		},
		{
			name:      "space string authorization, expect error",
			ctx:       context{request: reqSpaceStringAuth},
			auth:      "",
			expectErr: true,
		},
		{
			name:      "req with authorization, expect no error",
			ctx:       context{request: reqWithAuth},
			auth:      "this_is_auth",
			expectErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := &authContext{
				request: tt.ctx.request,
			}
			auth, err := ctx.GetAuthToken()
			if auth != tt.auth {
				t.Errorf("GetAuthToken() got = %v, want %v", auth, tt.auth)
			}
			if tt.expectErr {
				if err == nil {
					t.Errorf("GetAuthToken() expect error, but err == nil")
				}
			} else {
				if err != nil {
					t.Errorf("GetAuthToken() dont expect error, but got error")
				}
			}
		})
	}
}

func Test_authContext_GetBearerToken(t *testing.T) {
	type fields struct {
		request *http.Request
	}
	type args struct {
		s string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
		want1  output.GoliathError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := &authContext{
				request: tt.fields.request,
			}
			got, got1 := ctx.GetBearerToken(tt.args.s)
			if got != tt.want {
				t.Errorf("GetBearerToken() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("GetBearerToken() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
