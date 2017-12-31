package container

import "github.com/recursivecurry/gobox/instance"

type ValueWithDistance struct {
	Distance instance.Integer
	Value interface{}
}

type CloseNodes []ValueWithDistance

func (ns CloseNodes) Len() int {
	return len(ns)
}

func (ns CloseNodes) Less(i, j int) bool {
	return ns[i].Distance.Less(ns[j].Distance)
}

func (ns CloseNodes) Swap(i, j int) {
	ns[i], ns[j] = ns[j], ns[i]
}

func (ns *CloseNodes) Push(x interface{}) {
	*ns = append(*ns, x.(ValueWithDistance))
}

func (ns *CloseNodes) Pop() interface{} {
	old := *ns
	n := len(old)
	v := old[n-1]
	*ns = old[0 : n-1]
	return v
}


