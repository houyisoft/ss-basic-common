package stack

import (
	"fmt"
	"testing"
)

func TestHeapSort(t *testing.T) {
	type args struct {
		s []int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			args:args{
				s:[]int{-1,9, 0, 6, 5, 8, 2, 1, 7, 4, 3},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fmt.Println(tt.args.s[1:])
			HeapSort(tt.args.s)
			fmt.Println(tt.args.s[1:])
		})
	}
}
