package utils

import (
	"fmt"
	"testing"
)

func TestGetIpInfo(t *testing.T) {
	tests := []struct{ ip string }{
		{"122.55.248.179"},
		{"23.144.13.76"},
		{"1.49.191.255"},
	}

	for _, tt := range tests {
		ipInfo := GetIpInfo(tt.ip)
		fmt.Printf("IP信息:%s", ipInfo)
	}
}
