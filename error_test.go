package errors

import (
	. "errors"
	"reflect"
	"testing"
)

func TestError(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "1",
			args: args{
				text: "error string",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := Error(tt.args.text)
			if (err != nil) != tt.wantErr {
				t.Errorf("Error() error = '%v', wantErr '%v'", err, tt.wantErr)
			} else {
				if res := New("error string"); !reflect.DeepEqual(err, res) {
					t.Errorf("Error() error = '%v', wantErr '%v'", err, res)
				}
			}
		})
	}
}

func TestErrorf(t *testing.T) {
	type args struct {
		format string
		a      []interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "1",
			args: args{
				format: "error string: %s",
				a:      []interface{}{"reason"},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Errorf(tt.args.format, tt.args.a...); (err != nil) != tt.wantErr {
				t.Errorf("Errorf() error = '%v', wantErr '%v'", err, tt.wantErr)
			} else {
				if res := New("error string: reason"); !reflect.DeepEqual(err, res) {
					t.Errorf("Errorf() error = '%v', wantErr '%v'", err, res)
				}
			}
		})
	}
}
