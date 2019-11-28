package MergeSort

import (
	"reflect"
	"sort"
	"testing"
)

func Test_mergeOrderedArrays(t *testing.T) {
	type args struct {
		arr1 sort.IntSlice
		arr2 sort.IntSlice
	}
	tests := []struct {
		name string
		args args
		want sort.IntSlice
	}{
		{
			name: "合并2个有序数组生成新的有序数组",
			args: args{
				arr1: sort.IntSlice{1, 2, 2, 3, 4, 4, 5, 5, 6, 8},
				arr2: sort.IntSlice{2, 2, 3, 4, 4, 5, 6, 6, 7},
			},
			want: sort.IntSlice{1, 2, 2, 2, 2, 3, 3, 4, 4, 4, 4, 5, 5, 5, 6, 6, 6, 7, 8},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeOrderedArrays(tt.args.arr1, tt.args.arr2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeOrderedArrays() = %v, want %v", got, tt.want)
			}
		})
	}
}
