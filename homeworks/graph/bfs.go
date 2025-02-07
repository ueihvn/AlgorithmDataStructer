package graph

import (
	"fmt"
)

// Bfs return visited and path
// vertices start from index 0

func Bfs(s, v, e int, adjacencyList [][]int) ([]bool, []int) {
	visited := make([]bool, v)
	path := make([]int, v)
	for i := range path {
		path[i] = defaultPath
	}
	queue := NewQueue[int]()
	visited[s] = true
	queue.Push(s)
	for !queue.IsEmpty() {
		node := queue.Front()
		queue.Pop()
		for _, nv := range adjacencyList[node] {
			if visited[nv] == true {
				continue
			}
			visited[nv] = true
			queue.Push(nv)
			path[nv] = node
		}
	}
	return visited, path
}

func Path(path []int, s, f int) string {
	// start and dst is the same
	if s == f {
		return ""
	}
	way := make([]int, len(path))
	var m int
	if path[f] == -1 {
		return "no"
	}
	for {
		m++
		way[m] = f
		f = path[f]
		if f == s {
			m++
			way[m] = s
			break
		}
	}
	var res string
	for i := m - 1; i >= 0; i-- {
		res += fmt.Sprintf("%d ", way[i])
	}
	if len(res) > 0 {
		res = res[0 : len(res)-1]
	}
	return res
}

func TestBfs() {
	v := 6
	e := 8

	s := 0
	f := 5
	//
	edgeList := [][]int{
		{0, 1},
		{0, 3},
		{1, 2},
		{1, 3},
		{1, 5},
		{2, 5},
		{3, 4},
		{3, 5},
	}

	aL := BuildAdjacencyListFromEdgeList(v, e, edgeList)

	visited, path := Bfs(s, v, e, aL)
	fmt.Println(visited)
	fmt.Println(path)

	way := Path(path, s, f)
	fmt.Println(way)

}
