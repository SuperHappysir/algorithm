package MaxSubArray

var slc = make(map[int]int)

// 递归实现
func maxSubArray(nums []int) int {
	len1 := len(nums)
	if len1 == 1 {
		return nums[0]
	}

	var globalMaxSum = nums[0]

	caclMaxSum(&nums, len1-1, &globalMaxSum)

	return globalMaxSum
}

func caclMaxSum(nums *[]int, i int, globalMaxSum *int) int {
	if i == 0 {
		return (*nums)[0]
	}

	var (
		sum int
		ok  bool
	)
	if sum, ok = slc[i]; !ok {
		sum = caclMaxSum(nums, i-1, globalMaxSum)
		if sum > 0 {
			slc[i] = (*nums)[i] + sum
		} else {
			slc[i] = (*nums)[i]
		}
	}

	if slc[i] > *globalMaxSum {
		*globalMaxSum = slc[i]
	}

	return slc[i]
}

// 迭代实现
func maxSubArray2(nums []int) int {
	//  decare
	var (
		globalMaxSum = nums[0]
		sum          = 0
	)

	//  loop
	for _, v := range nums {
		/* 对于当前值v求当前最大sum的场景来说，计算最优解存在2种决策方式： 1.当前最大sum = v本身 2.当前最大sum = v + 子问题的最大sum  */
		/* 状态转移方程: dp[i] = max(nums[i], nums[i] + dp[i-1]) */
		/* sum 是包含当前值子数组的所有和（以当前v来说子问题的解,即dp）  */
		/* 在本次实现里面，我们是用一个变量来存储dp[i-1]，减少空间复杂度 */

		// 如果 sum > 0，则说明 sum 对结果有增益效果（是当前求sum的最优解），则 sum 保留并加上当前遍历数字
		// if sum + v > v {
		if sum > 0 {
			sum = sum + v
			// 如果 sum <= 0，则说明 sum 对结果无增益效果（相加后还不如当前数自身大，最优解为自身），需要舍弃，则 sum 直接更新为当前遍历数字
		} else {
			sum = v
		}

		// 更新目标结果最大值
		if globalMaxSum < sum {
			globalMaxSum = sum
		}
	}
	return globalMaxSum
}
