package graph

type Graph struct {
	Node map[int]bool
	Edge map[int]map[int]int
}
