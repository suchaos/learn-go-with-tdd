package passbyvalue

import (
	"testing"
)

func TestModifyValuesForInt(t *testing.T) {
	a := 10
	t.Log("Before:", a)
	modifyValuesForInt(a)
	t.Log("After:", a)
}

func TestModifyValuesForPointer(t *testing.T) {
	a := 10
	t.Log("Before:", a)
	modifyValuesForPointer(&a)
	t.Log("After:", a)
}

func TestModifyValuesForString(t *testing.T) {
	a := "before"
	t.Log("Before:", a)
	modifyValuesForString(a)
	t.Log("After:", a)
}

func TestModifyValuesForSliceInt(t *testing.T) {
	a := []int{1, 2, 3}
	t.Log("Before:", a)
	modifyValuesForSliceInt(a)
	t.Log("After:", a)

	b := make([]int, 10)
	b[0] = 1
	b[1] = 2
	t.Log("Before:", b)
	modifyValuesForSliceInt2(b)
	t.Log("After:", b)
}
