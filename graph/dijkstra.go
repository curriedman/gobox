package graph

import (
	"github.com/recursivecurry/gobox/util"
	"github.com/recursivecurry/gobox/container"
	"container/heap"
)

func (g *Graph) Dijkstra(start int) []int {
	var distance = make([]int, len(g.Node))
	for i := 0; i < len(g.Node); i++ {
		distance[i] = util.MaxInt
	}
	var visited map[int]bool = make(map[int]bool)
	var candidates container.CloseNodes = container.CloseNodes{[2]int{start, 0}}

	for len(candidates) > 0 {
		var next container.NodeWithDistance = heap.Pop(&candidates).(container.NodeWithDistance)

		// not implemented

		visited[next[0]] = true
	}
	return distance
}
