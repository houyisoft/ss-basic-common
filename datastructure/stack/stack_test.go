package stack

import (
	"fmt"
	"reflect"
	"testing"
)

func TestStack_Push(t *testing.T) {
	type args struct {
		x interface{}
	}
	tests := []struct {
		name string
		s    *Stack
		args args
	}{
		// TODO: Add test cases.
		{
			args:args{
				x:1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.s.Push(tt.args.x)
			fmt.Println(tt, tt.s.Len())
		})
	}
}

func TestStack_Pop(t *testing.T) {
	tests := []struct {
		name    string
		s       *Stack
		want    interface{}
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.Pop()
			if (err != nil) != tt.wantErr {
				t.Errorf("Stack.Pop() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Stack.Pop() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStack_Len(t *testing.T) {
	tests := []struct {
		name string
		s    *Stack
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Len(); got != tt.want {
				t.Errorf("Stack.Len() = %v, want %v", got, tt.want)
			}
		})
	}
}
