package Uniquify

import "sort"

// 一、先将数组进行排序操作
// 二、逐一对比互异“相邻” 元素，直至末元素
//     跳过不同者，发现不同元素时，将当前不同的元素放到前面已经去重的数据之后
// 三、最后删除后面所有的元素
func uniquify(arr []int) []int {
	// 排序
	sort.Ints(arr)

	var (
		i = 0
		j = 1
	)

	for j < len(arr) {
		// 发现不同元素，将当前不同的元素放到前面已经去重的数据之后
		if arr[i] != arr[j] {
			i++
			arr[i] = arr[j]
		}
		// 遍历下一个元素
		j++
	}

	// 截取[0,i+1)之间的元素作为唯一化后的值
	return arr[:i+1]
}
