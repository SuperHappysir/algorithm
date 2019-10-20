package QuickSort

import (
	"sort"
)

// 步骤为：
//
// 1. 挑选基准值：从数列中挑出一个元素，称为“基准”（pivot），
// 2. 分割：重新排序数列，所有比基准值小的元素摆放在基准前面，所有比基准值大的元素摆在基准后面（与基准值相等的数可以到任何一边）。在这个分割结束之后，对基准值的排序就已经完成，
// 3. 递归排序子序列：递归地将小于基准值元素的子序列和大于基准值元素的子序列排序。

func QuickSort(array sort.IntSlice) {
	length := len(array)
	if length <= 1 {
		return
	}
	// 挑选基准值
	pivotValue := array[0]

	// 首/尾分别为小于、大于基准值序列的当前索引
	head, tail := 0, length-1

	for index := 1; index <= tail; {
		if pivotValue > array[index] {
			// 如果基准值大于当前值,代表当前值应该放置到小于基准值的序列中
			// 交换当前值与head位的值
			// array[:head]代表的是小于基准值得序列
			array.Swap(index, head)
			head++
			index++
		} else {
			// 如果基准值小于当前值,代表当前值应该放置到大于基准值的序列中
			// 交换当前值与tail位的值
			// array[tail:]代表的是大于基准值得序列
			array.Swap(index, tail)
			tail--
		}
	}

	// 递归处理前序列/后序列
	QuickSort(array[:head])
	QuickSort(array[head+1:])
}

func QuickSort2(array sort.IntSlice) sort.IntSlice {
	length := len(array)
	if length <= 1 {
		return array
	}

	// 挑选基准值
	pivotValue := array[0]

	// leftArray,rightArray分别为比基准值小/大的子序列
	var leftArray, rightArray sort.IntSlice

	// 遍历数组，按照基准值将序列拆分
	for index := 1; index < length; index++ {
		if pivotValue > array[index] {
			leftArray = append(leftArray, array[index])
		} else {
			rightArray = append(rightArray, array[index])
		}
	}

	// 地柜将子序列继续按照快排排序处理
	rightArray = QuickSort2(rightArray)
	leftArray = QuickSort2(leftArray)
	return append(append(leftArray, pivotValue), rightArray...)
}
