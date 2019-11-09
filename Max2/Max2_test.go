package Max2

import (
	"fmt"
	"testing"
)

func TestMax2(t *testing.T) {
	slc := []int{1, 2, 3, 4, 5}

	m1 := 0
	m2 := 0
	max2(slc, 0, len(slc)-1, &m1, &m2)

	if m1 != 4 {
		fmt.Println(m1)
		t.Fail()
	}

	if m2 != 3 {
		fmt.Println(m2)
		t.Fail()
	}
}

func TestMax2_(t *testing.T) {
	slc := []int{1, 2, 3, 4, 5}

	m1 := 0
	m2 := 0
	max2_(slc, 0, len(slc), &m1, &m2)

	if m1 != 4 {
		fmt.Println(m1)
		t.Fail()
	}

	if m2 != 3 {
		fmt.Println(m2)
		t.Fail()
	}
}

func Test_Max2(t *testing.T) {
	slc := []int{1, 2, 3, 4, 5}

	m1 := 0
	m2 := 0
	_max2(slc, 0, len(slc), &m1, &m2)

	if m1 != 4 {
		fmt.Println(m1)
		t.Fail()
	}

	if m2 != 3 {
		fmt.Println(m2)
		t.Fail()
	}
}
