package util

import (
	"reflect"
	"testing"
)

func TestGenerateRandomNum(t *testing.T) {
	type args struct {
		start int
		end   int
		count int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{
			name: "TestGenerateRandomNum",
			args: args{
				start: 1,
				end:   20,
				count: 10,
			},
			want: []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GenerateRandomNum(tt.args.start, tt.args.end, tt.args.count); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GenerateRandomNum() = %v, want %v", got, tt.want)
			}
		})
	}
}
