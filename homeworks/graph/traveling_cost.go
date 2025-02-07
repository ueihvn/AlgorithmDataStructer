package graph

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
)

func travelingCost(v, e, s, q int, qs []int, edgeList [][3]int) []string {
	graph := make([][]Node, v)
	for _, edge := range edgeList {
		from := Node{
			vertices: edge[0],
			weighted: edge[2],
		}
		to := Node{
			vertices: edge[1],
			weighted: edge[2],
		}
		if el := graph[from.vertices]; el != nil {
			el = append(el, to)
			graph[from.vertices] = el
		} else {
			el = []Node{to}
			graph[from.vertices] = el
		}

		if el := graph[to.vertices]; el != nil {
			el = append(el, from)
			graph[to.vertices] = el
		} else {
			el = []Node{from}
			graph[to.vertices] = el
		}
	}
	_, dist := Dijkstra(s, v, e, graph)
	res := make([]string, 0, len(qs))
	for _, query := range qs {
		queryCost := cost(query, dist)
		res = append(res, queryCost)
	}
	return res
}

func cost(dst int, dist []int) string {
	if dist[dst] == defaultDist {
		return "NO PATH"
	}
	return strconv.Itoa(dist[dst])
}

func ConsoleTravelingCost() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	eText := strings.TrimSpace(scanner.Text())
	e, _ := strconv.Atoi(eText)
	edgeList := make([][3]int, 0, e)
	for i := 0; i < e; i++ {
		scanner.Scan()
		edge := scanner.Text()
		edges := strings.Split(edge, " ")
		u, _ := strconv.Atoi(edges[0])
		v, _ := strconv.Atoi(edges[1])
		w, _ := strconv.Atoi(edges[2])
		edgeList = append(edgeList, [3]int{u, v, w})
	}
	scanner.Scan()
	start := scanner.Text()
	s, _ := strconv.Atoi(start)
	scanner.Scan()
	query := scanner.Text()
	q, _ := strconv.Atoi(query)
	dests := make([]int, 0, q)
	for i := 0; i < q; i++ {
		scanner.Scan()
		dest := scanner.Text()
		dst, _ := strconv.Atoi(dest)
		dests = append(dests, dst)
	}
	costs := travelingCost(501, e, s, q, dests, edgeList)
	for _, c := range costs {
		fmt.Println(c)
	}
}

func TestTravelingCost() {
	type args struct {
		v        int
		e        int
		s        int
		q        int
		qs       []int
		edgeList [][3]int
	}

	cases := []struct {
		args args
		want []string
	}{
		/*
			{
				args: args{
					v:  500,
					e:  7,
					s:  0,
					q:  4,
					qs: []int{1, 4, 5, 7},
					edgeList: [][3]int{
						{0, 1, 4},
						{0, 3, 8},
						{1, 4, 1},
						{1, 2, 2},
						{4, 2, 3},
						{2, 5, 3},
						{3, 4, 2},
					},
				},
				want: []string{"4", "5", "9", "NO PATH"},
			},
		*/
		{
			args: args{
				v:  501,
				e:  20,
				s:  3,
				q:  5,
				qs: []int{8, 6, 2, 4, 5},
				edgeList: [][3]int{
					{10, 11, 11},
					{2, 3, 5},
					{2, 6, 3},
					{2, 1, 4},
					{2, 4, 2},
					{5, 2, 3},
					{4, 6, 2},
					{3, 4, 6},
					{3, 4, 6},
					{2, 3, 5},
					{3, 4, 6},
					{6, 4, 5},
					{1, 2, 4},
					{1, 3, 8},
					{1, 4, 1},
					{1, 4, 4},
					{4, 2, 3},
					{3, 4, 2},
					{2, 4, 1},
					{3, 4, 2},
				},
			},
			want: []string{
				"NO PATH",
				"4",
				"3",
				"2",
				"8",
			},
		},
	}
	for _, v := range cases {
		if got := travelingCost(
			v.args.v,
			v.args.e,
			v.args.s,
			v.args.q,
			v.args.qs,
			v.args.edgeList,
		); !reflect.DeepEqual(got, v.want) {
			fmt.Printf("args: %+v - got: %v / want: %v\n", v.args, got, v.want)
		}
	}
}
