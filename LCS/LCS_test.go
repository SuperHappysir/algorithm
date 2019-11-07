package LCS

import (
	"testing"
)

func TestLCSMaxLength(t *testing.T) {
	x := "ABCBDAB"
	y := "BDCABA"

	if MaxLength([]byte(x), []byte(y), len(x), len(y)) != 4 {
		t.Fail()
	}

	z := "ABCBDABCIODKSHAGHSJSKKA"
	a := "HAGHSBDCABAHAGHS"
	if MaxLength([]byte(z), []byte(a), len(z), len(a)) != 10 {
		t.Fail()
	}
}
