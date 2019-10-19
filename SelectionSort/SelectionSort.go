package SelectionSort

import "sort"

// 选择排序（Selection sort）是一种简单直观的排序算法。
// 它的工作原理如下。首先在未排序序列中找到最小元素，存放到排序序列的起始位置，
// 然后，再从剩余未排序元素中继续寻找最小元素，然后放到已排序序列的末尾。
// 以此类推，直到所有元素均排序完毕。

// eg:所谓选择排序，即选择最小的数，放置到最前的位置（相对于当前的循环而言）
// 按照算法拆小步骤的思维，我们拆分第一步，选择第一位应该放置的最小值所在的数组索引，与当前第一位索引值做交换操作
// 对第一位应该存放的最小值进行选择，选择前选择后的顺序如下
// [8] 3 6 9 5 1 2 3
// 1   3 6 9 5 8 2 3

// 关键字: 找最小元素, 交换位置

func SelectionSort(array sort.IntSlice) sort.IntSlice {
	for index := range array {
		minValue := array[index]
		minIndex := index

		// 找到最小值得下标
		for nextIndex := index + 1; nextIndex < len(array); nextIndex++ {
			if minValue > array[nextIndex] {
				minIndex = nextIndex
				minValue = array[nextIndex]
			}
		}

		// 交换
		// array[index], array[minIndex] = array[minIndex], array[index]
		array.Swap(index, minIndex)
	}

	return []int{}
}
