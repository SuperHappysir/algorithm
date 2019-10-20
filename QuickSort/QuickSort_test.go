package QuickSort

import (
	"fmt"
	"math/rand"
	"sort"
	"testing"
	"time"
)

func TestBubbleSort(t *testing.T) {
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	array1 := make([]int, random.Intn(20))
	for i := range array1 {
		array1[i] = random.Intn(100)
	}

	fmt.Println(array1)

	array2 := make(sort.IntSlice, len(array1))
	copy(array2, array1)

	QuickSort(array1)
	fmt.Println(array1)

	array2.Sort()
	fmt.Println(array2)

	for i := range array1 {
		if array1[i] != array2[i] {
			t.Fail()
		}
	}

}
