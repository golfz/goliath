package sliceutils

import (
	"reflect"
	"testing"
)

func TestTakeSliceArg(t *testing.T) {
	type args struct {
		arg interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantOut []interface{}
		wantOk  bool
	}{
		{
			name:    "arg is not slice, expect ok = false",
			args:    args{arg: 1},
			wantOut: nil,
			wantOk:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotOut, gotOk := TakeSliceArg(tt.args.arg)
			if !reflect.DeepEqual(gotOut, tt.wantOut) {
				t.Errorf("TakeSliceArg() gotOut = %v, want %v", gotOut, tt.wantOut)
			}
			if gotOk != tt.wantOk {
				t.Errorf("TakeSliceArg() gotOk = %v, want %v", gotOk, tt.wantOk)
			}
		})
	}
}

func Test_takeArg(t *testing.T) {
	type args struct {
		arg  interface{}
		kind reflect.Kind
	}
	tests := []struct {
		name    string
		args    args
		wantVal reflect.Value
		wantOk  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotVal, gotOk := takeArg(tt.args.arg, tt.args.kind)
			if !reflect.DeepEqual(gotVal, tt.wantVal) {
				t.Errorf("takeArg() gotVal = %v, want %v", gotVal, tt.wantVal)
			}
			if gotOk != tt.wantOk {
				t.Errorf("takeArg() gotOk = %v, want %v", gotOk, tt.wantOk)
			}
		})
	}
}
