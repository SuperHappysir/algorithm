package main

import (
	"fmt"
	"strconv"
)

func main() {
	slic := []interface{}{1, 2, 3}
	fmt.Print(subSet(slic))
}

func subSet(arr []interface{}) [][]interface{} {
	var (
		// 给定集合的长度
		num = len(arr)
		// 所有子集个数为2的n次方,详情查看数学集合部分
		maxCount = 1 << uint(num)
		// 子集数组
		subArray [][]interface{}
		//formatStr string
		nullArr   []interface{}
	)

	subArray = append(subArray, nullArr)

	for i := 1; i < maxCount; i++ {
		var tArr = nullArr
		for j := 0; j <= num; j++ {
			fmt.Println("")
			fmt.Println("源数据")
			fmt.Print(i, j)
			fmt.Println("")
			fmt.Println("二进制数据")
			fmt.Printf("%0"+strconv.Itoa(int(num))+"b", i)
			fmt.Println("")
			fmt.Printf("%0"+strconv.Itoa(int(num))+"b", j)
			fmt.Println("")
			fmt.Println("结果")
			fmt.Printf("%0"+strconv.Itoa(int(num))+"b", i & j)
			fmt.Println("")
			fmt.Println("")
			if (i & (1 << uint(j))) != 0 {
				tArr = append(tArr, arr[j])
			}
		}
		subArray = append(subArray, tArr)
	}
	return subArray
}
