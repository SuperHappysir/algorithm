package LCS

func MaxLength(sequenceA, sequenceB []byte, m, n int) int {
	// 递归基:
	// 2个序列中至少包含一个空序列
	if m == 0 || n == 0 {
		return 0
	}

	// 减而治之:
	// 如果2个序列最后一位相同
	// 则计算2个新序列（均去除最后一位的新序列）的最长公共子序列长度，然后加上当前长度1
	if sequenceA[m-1] == sequenceB[n-1] {
		return MaxLength(sequenceA, sequenceB, m-1, n-1) + 1
	}

	// 分而治之:
	// 如果2个序列最后一位不相同
	// 则计算LCS( sequenceA[0, n], sequenceB[0, m) )与LCS( sequenceA[0, n), sequenceB[0, m] )中的更长者
	len1 := MaxLength(sequenceA, sequenceB, m, n-1)
	len2 := MaxLength(sequenceA, sequenceB, m-1, n)
	if len1 > len2 {
		return len1
	}
	return len2
}
