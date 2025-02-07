package graph

import (
	"container/heap"
	"fmt"
)

type PriorityQueueNodes []Node

func (p *PriorityQueueNodes) Len() int {
	return len(*p)
}

func (p *PriorityQueueNodes) Less(i, j int) bool {
	return (*p)[i].weighted < (*p)[j].weighted
}

func (p *PriorityQueueNodes) Swap(i, j int) {
	(*p)[i], (*p)[j] = (*p)[j], (*p)[i]
}

func (p *PriorityQueueNodes) Push(x any) {
	*p = append(*p, x.(Node))
}

func (p *PriorityQueueNodes) Pop() any {
	old := *p
	x := len(old) - 1
	out := old[x]
	*p = old[:x]
	return out
}

func Dijkstra(start, vertices, edges int, graph [][]Node) ([]int, []int) {
	dist := make([]int, vertices)
	for i := range dist {
		dist[i] = defaultDist
	}
	path := make([]int, vertices)
	for i := range dist {
		path[i] = defaultPath
	}
	priorityQueue := &PriorityQueueNodes{}
	dist[start] = 0
	startNode := Node{
		vertices: start,
		weighted: 0,
	}
	heap.Push(priorityQueue, startNode)
	for len(*priorityQueue) != 0 {
		u := heap.Pop(priorityQueue).(Node)
		for _, node := range graph[u.vertices] {
			// reduce time complexity
			if dist[u.vertices] < u.weighted {
				continue
			}
			if distFromU := dist[u.vertices] + node.weighted; dist[node.vertices] > distFromU {
				dist[node.vertices] = distFromU
				heap.Push(priorityQueue, Node{vertices: node.vertices, weighted: dist[node.vertices]})
				path[node.vertices] = u.vertices
			}
		}
	}
	return path, dist
}

func TestDijkstra() {
	e, v := 10, 6
	edgesList := [][3]int{
		{0, 1, 1},
		{1, 2, 5},
		{1, 3, 2},
		{1, 5, 7},
		{2, 5, 1},
		{3, 0, 2},
		{3, 2, 1},
		{3, 4, 4},
		{4, 3, 3},
		{5, 4, 1},
	}
	matrix := DirectedWeightedAdjacencyMatrix(v, e, edgesList)
	for _, nodes := range matrix {
		for _, node := range nodes {
			fmt.Printf("%+v ", node)
		}
		fmt.Println()
	}
	graph := DirectedWeightedGraph(v, e, edgesList)
	for _, nodes := range graph {
		for _, node := range nodes {
			fmt.Printf("%+v ", node)
		}
		fmt.Println()
	}
	path, dist := Dijkstra(0, v, e, graph)
	fmt.Println(path)
	fmt.Println(dist)
}
