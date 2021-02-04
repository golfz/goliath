package contexts

import "testing"

func Test_isBearerAuth(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "empty string, expect false", args: args{s: ""}, want: false},
		{name: "space string, expect false", args: args{s: " "}, want: false},
		{name: "no bearer, expect false", args: args{s: "eyJD.bN2U="}, want: false},
		{name: "lower cap bearer, expect false", args: args{s: "bearer eyJD.bN2U="}, want: false},
		{name: "not start with bearer, expect false", args: args{s: "eyJD bearer bN2U="}, want: false},
		{name: "valid patter, expect true", args: args{s: "Bearer eyJD.bN2U="}, want: true},
		{name: "valid patter (bearer in body too), expect true", args: args{s: "Bearer eyJD. Bearer bN2U="}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isBearerAuth(tt.args.s); got != tt.want {
				t.Errorf("isBearerAuth() = %v, auth %v", got, tt.want)
			}
		})
	}
}
