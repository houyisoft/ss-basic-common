package file

import (
	"net/http"
	"testing"
)

func Test_uploadHandler(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		args args
	}{

		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uploadHandler(tt.args.w, tt.args.r)
		})
	}
}
