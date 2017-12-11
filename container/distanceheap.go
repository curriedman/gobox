package container

type NodeWithDistance = [2]int
type CloseNodes []NodeWithDistance

func (ns CloseNodes) Len() int {
	return len(ns)
}

func (ns CloseNodes) Less(i, j int) bool {
	return ns[i][1] < ns[j][1]
}

func (ns CloseNodes) Swap(i, j int) {
	ns[i], ns[j] = ns[j], ns[i]
}

func (ns *CloseNodes) Push(x interface{}) {
	*ns = append(*ns, x.(NodeWithDistance))
}

func (ns *CloseNodes) Pop() interface{} {
	old := *ns
	n := len(old)
	v := old[n-1]
	*ns = old[0 : n-1]
	return v
}


