package search

import "testing"

func TestBinary(t *testing.T) {
	type args struct {
		arr []int64
		n   int64
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "TestBinary",
			args: args{
				arr: []int64{0, 1, 2, 3, 3, 3, 4, 5, 6, 7, 8, 9},
				n:   3,
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Binary(tt.args.arr, tt.args.n); got != tt.want {
				t.Errorf("Binary() = %v, want %v", got, tt.want)
			}
		})
	}
}
