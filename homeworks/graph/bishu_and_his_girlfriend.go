package graph

import "fmt"

func BishuAndHisGirlFriend(n int, edgeList [][]int, q []int) int {
	visited := make([]bool, n+1)
	dist := make([]int, n+1)
	adjacencyList := BuildAdjacencyListFromEdgeList(n+1, n-1, edgeList)
	stack := NewStack[int]()
	start := 1
	visited[start] = true
	stack.Push(start)
	for !stack.IsEmpty() {
		node := stack.Top()
		stack.Pop()
		for _, v := range adjacencyList[node] {
			if visited[v] {
				continue
			}
			visited[v] = true
			dist[v] = dist[node] + 1
			stack.Push(v)
		}
	}
	var min int
	for _, g := range q {
		if dist[min] == 0 {
			min = g
			continue
		}
		if dist[min] > dist[g] {
			min = g
		} else if dist[min] == dist[g] && g < min {
			min = g
		}
	}
	return min
}

func ConsoleBishuAndHisGirlFriend() {
	var n int
	fmt.Scanln(&n)
	edgeList := make([][]int, n-1)
	for i := 0; i < n-1; i++ {
		list := make([]int, 2)
		var u int
		fmt.Scanf("%d", &u)
		list[0] = u
		var v int
		fmt.Scanf("%d", &v)
		list[1] = v
		edgeList[i] = list
	}
	var q int
	fmt.Scanln(&q)
	qs := make([]int, q)
	for i := 0; i < q; i++ {
		var x int
		fmt.Scanln(&x)
		qs[i] = x
	}
	fmt.Println(BishuAndHisGirlFriend(n, edgeList, qs))
}

func TestBishuAndHisGirlFriend() {
	type args struct {
		n        int
		edgeList [][]int
		q        []int
	}

	cases := []struct {
		args args
		want int
	}{
		{
			args: args{
				n: 6,
				edgeList: [][]int{
					{1, 2},
					{1, 3},
					{1, 4},
					{2, 5},
					{2, 6},
				},
				q: []int{5, 6, 3, 4},
			},
			want: 3,
		},
		{
			args: args{
				n: 9,
				edgeList: [][]int{
					{1, 6},
					{6, 9},
					{1, 2},
					{1, 3},
					{2, 4},
					{4, 7},
					{3, 5},
					{5, 8},
				},
				q: []int{7, 8, 9},
			},
			want: 9,
		},
	}
	for _, v := range cases {
		if got := BishuAndHisGirlFriend(v.args.n, v.args.edgeList, v.args.q); got != v.want {
			fmt.Printf("args: %+v - got: %v / want: %v\n", v.args, got, v.want)
		}
	}
}
