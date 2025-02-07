package graph

import "fmt"

func Dfs(s, v, e int, adjacencyList [][]int) ([]bool, []int) {
	visited := make([]bool, v)
	path := make([]int, v)
	for i := range path {
		path[i] = defaultPath
	}
	stack := NewStack[int]()
	visited[s] = true
	stack.Push(s)
	for !stack.IsEmpty() {
		node := stack.Top()
		stack.Pop()
		for _, nv := range adjacencyList[node] {
			if visited[nv] == true {
				continue
			}
			visited[nv] = true
			stack.Push(nv)
			path[nv] = node
		}
	}
	return visited, path
}

func DfsRecursive(s, v, e int, adjacencyList [][]int) ([]bool, []int) {
	visited := make([]bool, v)
	path := make([]int, v)
	for i := range path {
		path[i] = defaultPath
	}
	visited[s] = true
	dfsRecursive(s, visited, path, adjacencyList)
	return visited, path
}

func dfsRecursive(node int, visited []bool, path []int, adjacencyList [][]int) {
	for _, nv := range adjacencyList[node] {
		if visited[nv] == true {
			continue
		}
		visited[nv] = true
		path[nv] = node
		dfsRecursive(nv, visited, path, adjacencyList)
	}

}

func TestDfs() {

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

	visited, path := Dfs(s, v, e, aL)
	fmt.Println(visited)
	fmt.Println(path)

	way := Path(path, s, f)
	fmt.Println(way)
}

func TestDfsRecursive() {
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
	visited, path := DfsRecursive(s, v, e, aL)
	fmt.Println(visited)
	fmt.Println(path)

	way := Path(path, s, f)
	fmt.Println(way)
}
