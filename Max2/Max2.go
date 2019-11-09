package Max2

// 时间复杂度
//  	T(n) = 2 * T(n/2) + 2 = 5n/3 - 2
func max2(slc []int, low, high int, max1, ma2 *int) {
	if low+1 == high {
		*max1 = low
		*ma2 = low + 1
		if slc[*max1] < slc[*ma2] {
			swap(max1, ma2)
		}
		return
	}

	// 二分 , 分别递归查找左右2个序列的最大和第二大的值-
	var mid = (low + high) / 2
	var max1L, max2L, max1R, max2R int
	max2(slc, low, mid, &max1L, &max2L)
	max2(slc, mid, high, &max1R, &max2R)

	// 1. 左序列最大值和有序列最大值之中较大者就是我们想要的最大值max1
	// 2. 再用步骤1中获胜的序列（左序列或者右序列）的次大值与失败序列的最大值做比对，2个值更大者为我们要找的次大值ma2
	if slc[max1L] > slc[max1R] {
		*max1 = max1L
		if slc[max2L] > slc[max1R] {
			*ma2 = max2L
		} else {
			*ma2 = max1R
		}
	} else {
		*max1 = max1R
		if slc[max2R] > slc[max1L] {
			*ma2 = max2R
		} else {
			*ma2 = max1L
		}
	}
}

// 时间复杂度
// 最好情况:循环中只进行了1次比较，其余全走else
//      1 + (n - 2) * 1 = n - 1
// 最坏情况:每次循环都需要执行到最里层,然后执行交换
//      1 + (n - 2) * 2 = 2n - 3
func max2_(slc []int, low, high int, max1, max2 *int) {
	if high < 1 {
		return
	}

	// 比较第1个数和第二个数
	*max1 = low
	*max2 = low + 1
	if slc[*max1] < slc[*max2] {
		swap(max1, max2)
	}

	// 依次找出第一大和第二大的数
	for i := low + 2; i < high; i++ {
		// 将当前数与第二大的数做比较，当期数更大，则更新第二大的数
		if slc[*max2] < slc[i] {
			*max2 = i
			// 比较第二大的数与第一大的数
			if slc[*max1] < slc[*max2] {
				swap(max1, max2)
			}
		}
	}
}

func swap(max1 *int, max2 *int) {
	tmp := *max1
	*max1 = *max2
	*max2 = tmp
}

// 从数组区间A[low, high)中找出最大的2个整数m1, m2
// 时间复杂度 2n - 3 <O(n)>
//           循环1：n - 1
//           循环2：max1 - low - 1
//           循环3：high - max1 - 1
func _max2(slc []int, low, high int, max1, max2 *int) {
	*max1 = low
	*max2 = low
	for i := low + 1; i < high; i++ {
		if slc[*max1] < slc[i] { // 扫描slc[ low, high ), 找出最大值的index
			*max1 = i
		}
	}

	for i := low + 1; i < *max1; i++ {
		if slc[*max2] < slc[i] { // 扫描slc[low, max1)
			*max2 = i
		}
	}

	for i := *max1 + 1; i < high; i++ {
		if slc[*max2] < slc[i] { // 扫描slc( max1, high), 找出max2的index
			*max2 = i
		}
	}
}
