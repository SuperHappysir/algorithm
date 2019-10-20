package InsertSort

import (
	"fmt"
	"math/rand"
	"sort"
	"testing"
	"time"
)

func TestInsertSort(t *testing.T) {
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	array1 := make([]int, random.Intn(200))
	for i := range array1 {
		array1[i] = random.Intn(100)
	}
	array3 := array1

	array2 := make(sort.IntSlice, len(array1))
	copy(array2, array1)

	InsertSort(array1)
	array2.Sort()
	for i := range array1 {
		if array1[i] != array2[i] {
			t.Fail()
		}
	}

	fmt.Println(array3)
	fmt.Println(array1)

}
