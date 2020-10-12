package sort

import (
	"arithmetic/util"
	"testing"
)

func TestInsertionOrigin(t *testing.T) {
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
			name: "TestInsertionOrigin",
			args: args{
				arr: util.GenerateRandomNum(1, 100000, 99990),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// if got := InsertionOrigin(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
			// 	t.Errorf("InsertionOrigin() = %v, want %v", got, tt.want)
			// }
		})
	}
}

func TestInsertion(t *testing.T) {
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
			name: "TestInsertion",
			args: args{
				arr: util.GenerateRandomNum(1, 100000, 99990),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// if got := Insertion(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
			// 	t.Errorf("Insertion() = %v, want %v", got, tt.want)
			// }
		})
	}
}

func BenchmarkInsertion(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Insertion(util.GenerateRandomNum(1, 100000, 99990))
	}
}

func BenchmarkInsertionOrigin(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = InsertionOrigin(util.GenerateRandomNum(1, 100000, 99990))
	}
}
