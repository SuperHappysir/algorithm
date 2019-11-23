package Uniquify

import (
	"reflect"
	"testing"
)

func Test_uniquify(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"数组去重（高效算法）",
			args{[]int{1, 1, 1, 2, 2, 2, 3, 3, 4, 4, 10, 9, 8, 1, 1, 1, 2, 4, 5, 5, 6, 7, 8, 9, 10}},
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := uniquify(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("uniquify() = %v, want %v", got, tt.want)
			}
		})
	}
}
