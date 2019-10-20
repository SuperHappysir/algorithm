package InsertSort

import (
	"fmt"
	"sort"
)

//Insertion Sort 和打扑克牌时，从牌桌上逐一拿起扑克牌，在手上排序的过程相同。
//举例：
//Input: {5 2 4 6 1 3}。
//首先拿起第一张牌, 手上有 {5}。
//拿起第二张牌 2, 把 2 insert 到手上的牌 {5}, 得到 {2 5}。
//拿起第三张牌 4, 把 4 insert 到手上的牌 {2 5}, 得到 {2 4 5}。
//以此类推。

// 动画演示: https://www.cs.usfca.edu/~galles/visualization/ComparisonSort.html
// 			https://zh.wikipedia.org/wiki/File:Insertion-sort-example-300px.gif
func InsertSort(array sort.IntSlice) sort.IntSlice {
	fmt.Println("输入", array)

	// 遍历未排好序的序列
	for index := 0; index < len(array); index++ {

		var needInsertValue int
		needInsertValue = array[index]

		insertIndex := index
		// 逆向遍历已排好序的序列
		for subIndex := index - 1; subIndex >= 0; subIndex-- {
			fmt.Println(needInsertValue, subIndex, index, array[subIndex])
			if needInsertValue < array[subIndex] {
				insertIndex = subIndex
			} else {
				break
			}
		}
		fmt.Println("需要插入的位置", insertIndex, needInsertValue)
		fmt.Println(insertIndex, needInsertValue)

		fmt.Println(needInsertValue)
		for subIndex := index - 1; subIndex >= insertIndex; subIndex-- {
			fmt.Println(array)
			array[subIndex+1] = array[subIndex]
			fmt.Println(array)
		}
		fmt.Println(needInsertValue)

		array[insertIndex] = needInsertValue

	}

	return array
}
