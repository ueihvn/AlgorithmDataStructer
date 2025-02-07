package graph

import "math"

const (
	defaultPath = -1
	defaultDist = math.MaxInt
)

type Node struct {
	vertices int
	weighted int
}

func DirectedWeightedAdjacencyMatrix(vertices, edges int, edgesList [][3]int) [][]Node {
	matrix := make([][]Node, vertices)
	for _, edge := range edgesList {
		from := Node{
			vertices: edge[0],
		}
		to := Node{
			vertices: edge[1],
			weighted: edge[2],
		}
		if el := matrix[from.vertices]; el != nil {
			el[to.vertices] = to
		} else {
			el = make([]Node, vertices)
			el[to.vertices] = to
			matrix[from.vertices] = el
		}
	}
	return matrix
}

func DirectedWeightedGraph(vertices, edges int, edgesList [][3]int) [][]Node {
	graph := make([][]Node, vertices)
	for _, edge := range edgesList {
		from := edge[0]
		to := Node{
			vertices: edge[1],
			weighted: edge[2],
		}
		if el := graph[from]; el != nil {
			el = append(el, to)
			graph[from] = el
		} else {
			el = make([]Node, 0, vertices)
			el = append(el, to)
			graph[from] = el
		}
	}
	return graph
}

func BuildAdjacencyListFromEdgeList(v, e int, edgeList [][]int) [][]int {
	adjacencyList := make([][]int, v)
	for i := 0; i < e; i++ {
		// append only 0 or 1 for directed graph
		if al := adjacencyList[edgeList[i][0]]; al != nil {
			adjacencyList[edgeList[i][0]] = append(al, edgeList[i][1])
		} else {
			adjacencyList[edgeList[i][0]] = []int{edgeList[i][1]}
		}
		if al := adjacencyList[edgeList[i][1]]; al != nil {
			adjacencyList[edgeList[i][1]] = append(al, edgeList[i][0])
		} else {
			adjacencyList[edgeList[i][1]] = []int{edgeList[i][0]}
		}
	}
	return adjacencyList
}
