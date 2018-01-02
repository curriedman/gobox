package heap

import (
	"testing"
	"container/heap"
)


func TestEmptyCloseNodes(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("It doesn't panic")
		}
	}()
	h := &CloseNodes{}
	heap.Init(h)
	heap.Pop(h)
}

func TestCloseNodesPushPop(t *testing.T) {
	h := &CloseNodes{}

	heap.Push(h, ValueWithDistance{3, 0})
	heap.Push(h, ValueWithDistance{1, 1})
	heap.Push(h, ValueWithDistance{5, 2})

	if n, ok := heap.Pop(h).(ValueWithDistance); ok && n.Value != 1 {
		t.Errorf("First pop is not 1 : %d", n.Value)
	}

	if n, ok := heap.Pop(h).(ValueWithDistance); ok && n.Value != 0 {
		t.Errorf("Second pop is not 3 : %d", n.Value)
	}

	if n, ok := heap.Pop(h).(ValueWithDistance); ok && n.Value != 2 {
		t.Errorf("First pop is not 5 : %d", n.Value)
	}
}