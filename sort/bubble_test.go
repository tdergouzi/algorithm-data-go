package sort

import (
	"reflect"
	"testing"
)

func TestBubble(t *testing.T) {
	type args struct {
		arr []int64
	}
	tests := []struct {
		name string
		args args
		want []int64
	}{
		// TODO: Add test cases.
		{
			name: "TestBubble",
			args: args{
				arr: []int64{2, 1, 7, 4, 6, 8, 3, 9, 5},
			},
			want: []int64{1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Bubble(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Bubble() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_bubble(t *testing.T) {
	type args struct {
		arr []int64
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bubble(tt.args.arr)
		})
	}
}
