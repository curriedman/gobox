package container

import "github.com/recursivecurry/gobox/instance"

type IntHeap []instance.Integer

func (h IntHeap) Len() int {
	return len(h)
}

func (h IntHeap) Less(i, j int) bool {
	return h[i].Less(h[j])
}

func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(instance.Integer))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	v := old[n-1]
	*h = old[0 : n-1]
	return v
}

