package graph

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
)

// bfs for each query
// update find path using count
func BfsShortestReach(s, e, v int, edgeList [][]int) []int {
	adjacencyList := BuildAdjacencyListFromEdgeList(v+1, e, edgeList)
	dist := make([]int, v+1)
	visited := make([]bool, v+1)
	//
	visited[s] = true
	queue := NewQueue[int]()
	queue.Push(s)
	for !queue.IsEmpty() {
		u := queue.Front()
		queue.Pop()
		nodes := adjacencyList[u]
		for _, node := range nodes {
			if visited[node] == true {
				continue
			}
			visited[node] = true
			queue.Push(node)
			dist[node] += dist[u] + 1
		}
	}
	res := make([]int, 0, v)
	for i := 1; i <= v; i++ {
		if i == s {
			continue
		} else if visited[i] == false {
			res = append(res, -1)
		} else {
			res = append(res, dist[i]*6)
		}
	}
	return res
}

func ConsoleBfsShortestReach() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	q, _ := strconv.Atoi(scanner.Text())
	for j := 0; j < q && scanner.Scan(); j++ {
		ve := scanner.Text()
		ves := strings.Split(ve, " ")
		v, _ := strconv.Atoi(ves[0])
		e, _ := strconv.Atoi(ves[1])
		edgeList := make([][]int, 0, e)
		for i := 0; i < e && scanner.Scan(); i++ {
			uv := scanner.Text()
			uvs := strings.Split(uv, " ")
			from, _ := strconv.Atoi(uvs[0])
			to, _ := strconv.Atoi(uvs[1])
			edgeList = append(edgeList, []int{from, to})
		}
		scanner.Scan()
		s := scanner.Text()
		start, _ := strconv.Atoi(s)
		var result string
		res := BfsShortestReach(start, e, v, edgeList)
		for _, v := range res {
			result += fmt.Sprintf("%d ", v)
		}
		fmt.Println(result)
	}
}

func TestBfsShortestReach() {
	type args struct {
		s        int
		e        int
		v        int
		edgeList [][]int
	}

	cases := []struct {
		args args
		want []int
	}{
		{
			args: args{
				s: 1,
				v: 4,
				e: 2,
				edgeList: [][]int{
					{1, 2},
					{1, 3},
				},
			},
			want: []int{6, 6, -1},
		},
		{
			args: args{
				s: 2,
				v: 3,
				e: 1,
				edgeList: [][]int{
					{2, 3},
				},
			},
			want: []int{-1, 6},
		},
	}

	for _, v := range cases {
		if got := BfsShortestReach(v.args.s, v.args.e, v.args.v, v.args.edgeList); !reflect.DeepEqual(got, v.want) {
			fmt.Printf("args: %+v - got: %v / want: %v\n", v.args, got, v.want)
		}
	}
}
