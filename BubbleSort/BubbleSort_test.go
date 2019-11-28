package BubbleSort

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

	BubbleSort(array1)

	array2.Sort()
	for i := range array1 {
		if array1[i] != array2[i] {
			t.Fail()
		}
	}

	fmt.Println(array1)
	fmt.Println(array2)

}

func TestBubbleSort2(t *testing.T) {
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	array1 := make([]int, random.Intn(20))
	for i := range array1 {
		array1[i] = random.Intn(100)
	}

	fmt.Println(array1)

	array2 := make(sort.IntSlice, len(array1))
	copy(array2, array1)

	BubbleSort2(array1)

	array2.Sort()
	for i := range array1 {
		if array1[i] != array2[i] {
			t.Fail()
		}
	}

	fmt.Println(array1)
	fmt.Println(array2)

}

func TestBubbleSort3(t *testing.T) {
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	array1 := make([]int, random.Intn(20))
	for i := range array1 {
		array1[i] = random.Intn(100)
	}

	fmt.Println(array1)

	array2 := make(sort.IntSlice, len(array1))
	copy(array2, array1)

	BubbleSort3(array1)

	array2.Sort()
	for i := range array1 {
		if array1[i] != array2[i] {
			t.Fail()
		}
	}

	fmt.Println(array1)
	fmt.Println(array2)

}
