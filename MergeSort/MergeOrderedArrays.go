package MergeSort

import "sort"

// 合并2个有序数组
// 时间复杂度O(len(arr1)+len(arr2)) -> O(n)
// 空间复杂度O(len(arr1)+len(arr2)) -> O(n)
func mergeOrderedArrays(arr1, arr2 sort.IntSlice) sort.IntSlice {
	var (
		idx1, idx2 int
		len1, len2 = len(arr1), len(arr2)
		result     = sort.IntSlice{}
	)
	if len1 == 0 {
		return arr2
	}
	if len2 == 0 {
		return arr1
	}

	for idx1 < len1 && idx2 < len2 {
		if arr1[idx1] < arr2[idx2] {
			result = append(result, arr1[idx1])
			idx1++
		} else {
			result = append(result, arr2[idx2])
			idx2++
		}
	}

	if idx1 < len1 {
		result = append(result, arr1[idx1:]...)
	}

	if idx2 < len2 {
		result = append(result, arr1[idx2:]...)
	}

	return result
}
