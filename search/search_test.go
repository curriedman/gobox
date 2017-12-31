package search

import (
	"testing"
)

type SortedIntSlice []int

func (s SortedIntSlice) Len() int {
	return len(s)
}

func (s SortedIntSlice) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s *SortedIntSlice) Swap(i, j int) {
	(*s)[i], (*s)[j] = (*s)[j], (*s)[i]
}

func (s SortedIntSlice) Greater(i int, v interface{}) bool {
	if vv, ok := v.(int); ok {
		return s[i] > vv
	}
	return false
}

func (s SortedIntSlice) Equal(i int, v interface{}) bool {
	if vv, ok := v.(int); ok {
		return s[i] == vv
	}
	return false
}

func TestLowerBound(t *testing.T) {
	s := SortedIntSlice([]int{1, 1, 3, 3, 3, 5, 5})

	if 0 != LowerBound(&s, 0) {
		t.Errorf("Wrong LowerBound of 0")
	}

	if 0 != LowerBound(&s, 1) {
		t.Errorf("Wrong LowerBound of 1")
	}

	if 2 != LowerBound(&s, 3) {
		t.Errorf("Wrong LowerBound of 3")
	}

	if 5 != LowerBound(&s, 4) {
		t.Errorf("Wrong LowerBound of 4")
	}

	if 5 != LowerBound(&s, 5) {
		t.Errorf("Wrong LowerBound of 5")
	}

	if 7 != LowerBound(&s, 6) {
		t.Errorf("Wrong LowerBound of 6")
	}
}

func TestUpperBound(t *testing.T) {
	s := SortedIntSlice([]int{1, 1, 3, 3, 3, 5, 5})

	if 0 != UpperBound(&s, 0) {
		t.Errorf("Wrong UpperBound of 0")
	}

	if 2 != UpperBound(&s, 1) {
		t.Errorf("Wrong UpperBound of 1")
	}

	if 5 != UpperBound(&s, 3) {
		t.Errorf("Wrong UpperBound of 3")
	}

	if 5 != UpperBound(&s, 4) {
		t.Errorf("Wrong UpperBound of 4")
	}

	if 7 != UpperBound(&s, 5) {
		t.Errorf("Wrong UpperBound of 5")
	}

	if 7 != UpperBound(&s, 6) {
		t.Errorf("Wrong UpperBound of 6")
	}
}
