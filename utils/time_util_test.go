package utils

import (
	"reflect"
	"testing"
	"time"
)

func TestGetJsonTimeWithString(t *testing.T) {
	type args struct {
		date string
	}
	tests := []struct {
		name string
		args args
		want JsonTime
	}{
		// TODO: Add test cases.
		{
			"abc",
			args{
				date:"2018-01-01",
			},
			JsonTime(time.Now()),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetJsonTimeWithString(tt.args.date); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetJsonTimeWithString() = %v, want %v", got, tt.want)
			}
		})
	}
}
