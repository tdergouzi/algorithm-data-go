package sort

import (
	"arithmetic/util"
	"testing"
)

// 单元测试
func TestSelection(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{
			name: "TestSelection",
			args: args{
				util.GenerateRandomNum(1, 100000, 99990),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// if got := Selection(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
			// 	t.Errorf("Selection() = %v, want %v", got, tt.want)
			// }
		})
	}
}

// 基准测试
func BenchmarkSelection(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Selection(util.GenerateRandomNum(1, 100000, 99990))
	}
}
