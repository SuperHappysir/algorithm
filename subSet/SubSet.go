package SubSet

func SubSet(array []interface{}) [][]interface{} {
	var (
		// 给定集合的长度
		num = len(array)
		// 所有子集个数为2的n次方个,详情查看数学集合部分
		maxCount = 1 << uint(num)
		// 子集数组
		subArray [][]interface{}
		nullArr  []interface{}
	)

	subArray = append(subArray, nullArr)

	// 共有maxCount个子集
	// 将每一种组合转换成bitmap, 使用按位与的方式来确定子集实际元素
	for i := 1; i < maxCount; i++ {
		var tArr = nullArr
		// 获取当前子集的实际组合
		// 这里遍历数组是为了将当前外层循环的可能性转换成实际组合数组
		// 使用按位与的方式来确定当前元素是否在 i对应子集 转换的bitmap中
		for j := 0; j <= num; j++ {
			if (i & (1 << uint(j))) != 0 {
				tArr = append(tArr, array[j])
			}
		}
		subArray = append(subArray, tArr)
	}
	return subArray
}
