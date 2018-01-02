package heap

import "github.com/recursivecurry/gobox/typeclass/ord"

type OrdHeap []ord.Interface

func (h OrdHeap) Len() int {
	return len(h)
}

func (h OrdHeap) Less(i, j int) bool {
	return h[i].Less(h[j])
}

func (h OrdHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *OrdHeap) Push(x interface{}) {
	*h = append(*h, x.(ord.Interface))
}

func (h *OrdHeap) Pop() interface{} {
	old := *h
	n := len(old)
	v := old[n-1]
	*h = old[0 : n-1]
	return v
}
