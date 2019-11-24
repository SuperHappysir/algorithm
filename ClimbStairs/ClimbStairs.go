package ClimbStairs

// var mp = make(map[int]int)
func climbStairs(n int) int {
	// 动态规划实现
	// 要爬到第n梯，有2种方案： 1.从第n-1梯爬到第n梯 2.从第n-2梯爬到第n梯
	// 要爬到第n-1梯，有2种方案： 1.从第n-2梯爬到第n-1梯 2.从第n-3梯爬到第n-1梯
	// 不难发现要到第n梯，我们只需要计算到n-1梯的次数和到n-2梯的次数之和
	// f(n) = f(n-1) + f(n - 2)
	// 时间复杂度O(n) 空间复杂度O(1)

	if n <= 2 {
		return n
	}
	var (
		i = 1 // n = 1
		j = 2 // n = 2
	)

	for x := 3; x <= n; x++ {
		j = j + i // 根据f(n) = f(n-1) + f(n - 2)公式计算
		i = j - i // 更新f(n-1)为后续的计算做基础
	}

	return j

	// 递归实现
	// 时间复杂度O(n) 空间复杂度O(n)
	// if n == 0 {
	//     return 0
	// }
	// if n <= 2 {
	//     return n
	// }
	// if _,ok := mp[n-2]; !ok {
	//     mp[n-2] = climbStairs(n - 2)
	// }
	// if _,ok := mp[n-1]; !ok {
	//     mp[n-1] = climbStairs(n - 1)
	// }

	// return mp[n-1] + mp[n-2]
}
