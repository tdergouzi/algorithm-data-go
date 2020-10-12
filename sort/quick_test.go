package sort

import (
	"reflect"
	"testing"
)

func TestQuick(t *testing.T) {
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
			name: "TestQuick",
			args: args{
				arr: []int64{7, 2, 1, 4, 4, 6, 5, 9, 8},
			},
			want: []int64{1, 2, 4, 4, 5, 6, 7, 8, 9},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Quick(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Quick() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_partition(t *testing.T) {
	type args struct {
		arr []int64
		p   int
		q   int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "Test_partition",
			args: args{
				arr: []int64{2, 1},
				p:   0,
				q:   1,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := partition(tt.args.arr, tt.args.p, tt.args.q); got != tt.want {
				t.Errorf("partition() = %v, want %v", got, tt.want)
			}
		})
	}
}
