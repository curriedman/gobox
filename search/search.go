package search

import "sort"

type Interface interface {
	sort.Interface
	Greater(i int, v interface{}) bool
	Equal(i int, v interface{}) bool
}

func LowerBound(s Interface, v interface{}) int {
	return sort.Search(s.Len(), func(i int) bool {
		return s.Greater(i, v) || s.Equal(i, v)
	})
}

func UpperBound(s Interface, v interface{}) int {
	return sort.Search(s.Len(), func(i int) bool {
		return s.Greater(i, v)
	})
}
