package contexts

import (
	"net/http"
	"reflect"
	"testing"
)

func TestNewGoliathContext(t *testing.T) {
	type args struct {
		r *http.Request
	}
	tests := []struct {
		name string
		args args
		want *goliathContext
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewGoliathContext(tt.args.r); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewGoliathContext() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_goliathContext_Auth(t *testing.T) {
	type fields struct {
		request     *http.Request
		authContext *authContext
	}
	tests := []struct {
		name   string
		fields fields
		want   *authContext
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := &goliathContext{
				request:     tt.fields.request,
				authContext: tt.fields.authContext,
			}
			if got := ctx.Auth(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Auth() = %v, want %v", got, tt.want)
			}
		})
	}
}
