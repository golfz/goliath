package contexts

import (
	"net/http"
	"net/http/httptest"
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
			if tt.expectErr {
				if err == nil {
					t.Errorf("GetAuthToken() expect error, but err == nil")
				}
			} else {
				if auth != tt.auth {
					t.Errorf("GetAuthToken() got = %v, auth %v", auth, tt.auth)
				}
				if err != nil {
					t.Errorf("GetAuthToken() dont expect error, but got error")
				}
			}
		})
	}
}

func Test_authContext_GetBearerToken(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name      string
		args      args
		auth      string
		expectErr bool
	}{
		{
			name:      "empty string, expect error",
			args:      args{s: ""},
			auth:      "",
			expectErr: true,
		},
		{
			name:      "space string, expect error",
			args:      args{s: "  "},
			auth:      "",
			expectErr: true,
		},
		{
			name:      "no bearer, expect error",
			args:      args{s: "this_is_auth"},
			auth:      "",
			expectErr: true,
		},
		{
			name:      "lower-cased bearer, expect error",
			args:      args{s: "bearer this_is_auth"},
			auth:      "",
			expectErr: true,
		},
		{
			name:      "lower-cased bearer, expect error",
			args:      args{s: "Bearer_this_is_auth"},
			auth:      "",
			expectErr: true,
		},
		{
			name:      "not start with bearer, expect error",
			args:      args{s: "this_is_auth Bearer auth"},
			auth:      "",
			expectErr: true,
		},
		{
			name:      "valid format, expect auth without error",
			args:      args{s: "Bearer this_is_auth"},
			auth:      "this_is_auth",
			expectErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := &authContext{}
			auth, err := ctx.GetBearerToken(tt.args.s)
			if tt.expectErr {
				if err == nil {
					t.Errorf("GetBearerToken() expect error, but err == nil")
				}
			} else {
				if auth != tt.auth {
					t.Errorf("GetBearerToken() auth = %v, auth %v", auth, tt.auth)
				}
				if err != nil {
					t.Errorf("GetBearerToken() dont expect error, but got error")
				}
			}
		})
	}
}
