package datasource

import (
	"fmt"
	"testing"
)

func TestGetOrderNo(t *testing.T) {
	type args struct {
		platformId int
		proxyId    int64
		userId     int64
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			"aa",
			args{platformId: 2, proxyId: 174972895133437966, userId: 174972895133437980},
			"234",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetOrderNo(tt.args.platformId, tt.args.proxyId, tt.args.userId); got != tt.want {
				fmt.Println("got:" + got)
				t.Errorf("GetOrderNo() = %v, want %v", got, tt.want)
			}
		})
	}
}
