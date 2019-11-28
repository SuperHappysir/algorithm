package BubbleSort

import "sort"

// 它重复地走访过要排序的数列，一次比较两个元素，如果他们的顺序错误就把他们交换过来。
// 走访数列的工作是重复地进行直到没有再需要交换，也就是说该数列已经排序完成。
// 这个算法的名字由来是因为越小的元素会经由交换慢慢“浮”到数列的顶端。

// 冒泡排序算法的运作如下：
// 比较相邻的元素。如果第一个比第二个大，就交换他们两个。
// 对每一对相邻元素作同样的工作，从开始第一对到结尾的最后一对。这步做完后，最后的元素会是最大的数。
// 针对所有的元素重复以上的步骤，除了最后一个。
// 持续每次对越来越少的元素重复上面的步骤，直到没有任何一对数字需要比较。

func BubbleSort(array sort.IntSlice) sort.IntSlice {

	length := len(array)

	/*
	   依次比较每一对相邻元素，如有必要，交换之
	   若整趟扫描都没有进行交换，则排序完成，否则再做一趟扫描交换
	*/
	sorted := false
	for ; !sorted; length-- { // 逐趟扫描交换,直至完全有序
		sorted = true                                        // 我们假设默认全局有序
		for subIndex := 0; subIndex < length-1; subIndex++ { // 自左向右，逐渐检查A[0,n)内各相邻元素
			// 如果前一位的数比后一位的数大（逆序）
			if array[subIndex] > array[subIndex+1] {
				array.Swap(subIndex, subIndex+1) // 就交换他们两个
				sorted = false                   // 清楚全局有序标志
			}
		}
	}

	// // 循环冒泡倒数第index位
	// // 第一次需要冒泡到length位，需要对比的最后位置为length-1,length
	// for index := 0; index < length-1; index++ {
	// 	// 两两对比，将最大的数冒泡到倒数第length - index - 1位（length - index - 1是为了排除已冒泡完成的位）
	// 	for subIndex := 0; subIndex < length-index-1; subIndex++ {
	// 		// 如果前一位的数比后一位的数大，就交换他们两个
	// 		if array[subIndex] > array[subIndex+1] {
	// 			array.Swap(subIndex, subIndex+1)
	// 		}
	// 	}
	// }

	return array
}

func BubbleSort2(array sort.IntSlice) sort.IntSlice {

	hi := len(array)
	// while(!bubble(array, 9, hi--))
	for !bubble(array, 0, hi) { // 逐趟做扫描交换，直至全序
		hi--
	}
	return array
}

func bubble(array sort.IntSlice, lo, hi int) bool {
	sorted := true // 全局有序标志（用于做冒泡排序的优化，当对未排序序列做完一趟扫描后发现全部有序，那么剩下的就不用排序了，已经达到了全局有序的目的）
	// 同其他语言: while (++lo < hi) {}
	for lo++; lo < hi; lo++ {
		if array[lo-1] > array[lo] { // 若为逆序，进行交换
			sorted = false
			//  swap
			array[lo-1], array[lo] = array[lo], array[lo-1]
		}
	}

	return sorted
}
