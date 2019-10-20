package InsertSort

import "sort"

//Insertion Sort 和打扑克牌时，从牌桌上逐一拿起扑克牌，在手上排序的过程相同。
//举例：
//Input: {5 2 4 6 1 3}。
//首先拿起第一张牌, 手上有 {5}。
//拿起第二张牌 2, 把 2 insert 到手上的牌 {5}, 得到 {2 5}。
//拿起第三张牌 4, 把 4 insert 到手上的牌 {2 5}, 得到 {2 4 5}。
//以此类推。

// 将未排好序的序列 循环与 排好序的序列(内层循环遍历完过后，当前外层循环index之前的位置都是拍好序的) 循环做对比
// 外层循环是遍历未排好序的序列，内层循环是遍历排好序的序列，内层循环需要逆向遍历
// 将外层循环当前遍历的值插入到拍好序的序列中正确的位置
// 动画演示: https://www.cs.usfca.edu/~galles/visualization/ComparisonSort.html
// 			https://zh.wikipedia.org/wiki/File:Insertion-sort-example-300px.gif
func InsertSort(array sort.IntSlice) sort.IntSlice {

	// 遍历未排好序的序列
	for index := 1; index < len(array); index++ {

		// 逆向遍历已排好序的序列

		// 我们需要将外层循环遍历的数值【index位上的值】插入到已排好序序列的正确顺序位置
		// 因此我们需要将当前值之前已经排序好的序列做【倒序遍历】，
		// 两两做对比，然后依次交换位置，直到插入到了应该插入的位置
		for subIndex := index - 1; subIndex >= 0; subIndex-- {
			if array[subIndex+1] < array[subIndex] {
				array.Swap(subIndex+1, subIndex)
			} else {
				break
			}
		}
	}

	return array
}
