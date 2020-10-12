package sort

import (
	"reflect"
	"testing"
)

func TestMerge(t *testing.T) {
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
			name: "TestMerge",
			args: args{
				arr: []int64{9, 8, 7, 6, 5, 4, 3, 2, 1},
			},
			want: []int64{1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Merge(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Merge() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMergeM(t *testing.T) {
	type args struct {
		arr []int64
		s   int
		m   int
		e   int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			name: "TestMergeM",
			args: args{
				arr: []int64{5, 6, 7, 8, 9, 1, 2, 3, 4},
				s:   0,
				m:   5,
				e:   9,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Log(tt.args.arr)
			MergeM(tt.args.arr, tt.args.s, tt.args.m, tt.args.e)
			t.Log(tt.args.arr)
		})
	}
}
