package SubSet

import (
	"fmt"
	"testing"
)

func TestSubSet(t *testing.T) {
	inputSet := []interface{}{1, 2, 3, 4, 5}

	outputSet := SubSet(inputSet)

	fmt.Println(outputSet)

	expectSet := [][]interface{}{{}, {1}, {2}, {1, 2}, {3}, {1, 3}, {2, 3}, {1, 2, 3}, {4}, {1, 4}, {2, 4}, {1, 2, 4}, {3, 4}, {1, 3, 4}, {2, 3, 4}, {1, 2, 3, 4}, {5}, {1, 5}, {2, 5}, {1, 2, 5}, {3, 5}, {1, 3, 5}, {2, 3, 5}, {1, 2, 3, 5}, {4, 5}, {1, 4, 5}, {2, 4, 5}, {1, 2, 4, 5}, {3, 4, 5}, {1, 3, 4, 5}, {2, 3, 4, 5}, {1, 2, 3, 4, 5}}
	fmt.Println(expectSet)

	if len(expectSet) != len(outputSet) {
		t.Error("结果集合长度与期望集合长度不同")
	}

	for index, arraySlice := range expectSet {

		expectSubSet := expectSet[index]

		for subIndex, value := range arraySlice {
			if value != expectSubSet[subIndex] {
				t.Error("结果子集合与期望子集合不同")
			}
		}
	}
}
