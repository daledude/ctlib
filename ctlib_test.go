package ctlib

import (
	"os"
	"reflect"
	"testing"
)

func TestExecute(t *testing.T) {
	type args struct {
		template      string
		discardStderr bool
	}

	tests := []struct {
		name    string
		args    args
		want    *string
		wantErr bool
	}{
		{
			"renders trivial",
			args{"test", true},
			stringptr("test"),
			false,
		},
		{
			"renders environment variable",
			args{`{{ env "CTLIB_TEST" }}`, true},
			stringptr("test"),
			false,
		},
		{
			"renders consul",
			args{`{{ key "test" }}`, true},
			stringptr("test"),
			false,
		},
	}

	os.Setenv("CTLIB_TEST", "test")
	os.Setenv("CONSUL_HTTP_ADDR", "http://127.0.0.1:8500")

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Execute(tt.args.template, tt.args.discardStderr)
			if (err != nil) != tt.wantErr {
				t.Errorf("Execute() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Execute() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func stringptr(str string) *string {
	return &str
}
