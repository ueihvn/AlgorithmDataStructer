package heap

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
)

type in struct {
	id int
	z  int
	p  int
	l  int
	c  int
	s  int
}

type out struct {
	id int
	z  int
	zC int
}

type outMinHeap []out

func (o *outMinHeap) Len() int {
	return len(*o)
}

func (o *outMinHeap) Less(i, j int) bool {
	access := *o
	if access[i].zC != access[j].zC {
		return access[i].zC < access[j].zC
	}
	return access[i].id < access[j].id
}

func (o *outMinHeap) Swap(i, j int) {
	access := *o
	access[i], access[j] = access[j], access[i]
}

func (o *outMinHeap) Push(x any) {
	*o = append(*o, x.(out))
}

func (o *outMinHeap) Pop() any {
	old := *o
	n := len(old)
	x := old[n-1]
	*o = old[0 : n-1]
	return x
}

func buildOuts(v in) out {
	ot := out{
		id: v.id,
		z:  v.z,
	}
	newZ := 50*v.p + 5*v.l + 10*v.c + 20*v.s
	zC := newZ - ot.z
	ot.zC = zC
	ot.z = newZ
	return ot
}

func royAndTrendingTopics(n int, ins []in) []out {
	const lens = 5
	temp := make([]out, 0, 5)
	minHeap := outMinHeap(temp)
	for i := 0; i < lens; i++ {
		ot := buildOuts(ins[i])
		heap.Push(&minHeap, ot)
	}
	for i := lens; i < n; i++ {
		ot := buildOuts(ins[i])
		if min := minHeap[0]; min.zC < ot.zC || (min.zC == ot.zC && min.id < ot.id) {
			heap.Pop(&minHeap)
			heap.Push(&minHeap, ot)
		}
	}
	res := make([]out, lens)
	for i := 0; i < lens; i++ {
		v := heap.Pop(&minHeap).(out)
		v.zC = 0
		res[lens-1-i] = v
	}
	return res
}

func ReadInputRoyAndTrendingTopics() []in {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	n, err := strconv.Atoi(scanner.Text())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading number of lines: %s\n", err)
		return nil
	}

	var res []in

	// Prepare to read the next n lines of data
	for i := 0; i < n && scanner.Scan(); i++ {
		line := scanner.Text()
		parts := strings.Split(line, " ")
		row := make([]int, len(parts))

		for j, part := range parts {
			row[j], err = strconv.Atoi(part)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error converting string to integer: %s\n", err)
				return nil
			}
		}
		val := in{id: row[0], z: row[1], p: row[2], l: row[3], c: row[4], s: row[5]}
		res = append(res, val)
	}

	// Check for scanning errors
	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading input: %s\n", err)
		return nil
	}
	return res
}

func ConsoleRoyAndTrendingTopics() {
	ins := ReadInputRoyAndTrendingTopics()
	res := royAndTrendingTopics(len(ins), ins)
	for _, v := range res {
		fmt.Printf("%d %d\n", v.id, v.z)
	}
}

func TestRoyAndTrendingTopics() {
	type args struct {
		n   int
		ins []in
	}

	cases := []struct {
		args args
		want []out
	}{
		{
			args: args{
				n: 8,
				ins: []in{
					{
						id: 1003,
						z:  100,
						p:  4,
						l:  0,
						c:  0,
						s:  0,
					},
					{
						id: 1002,
						z:  200,
						p:  6,
						l:  0,
						c:  0,
						s:  0,
					},
					{
						id: 1001,
						z:  300,
						p:  8,
						l:  0,
						c:  0,
						s:  0,
					},
					{
						id: 1004,
						z:  100,
						p:  3,
						l:  0,
						c:  0,
						s:  0,
					},
					{
						id: 1005,
						z:  200,
						p:  3,
						l:  0,
						c:  0,
						s:  0,
					},
					{
						id: 1006,
						z:  300,
						p:  5,
						l:  0,
						c:  0,
						s:  0,
					},
					{
						id: 1007,
						z:  100,
						p:  3,
						l:  0,
						c:  0,
						s:  0,
					},
					{
						id: 999,
						z:  100,
						p:  4,
						l:  0,
						c:  0,
						s:  0,
					},
				},
			},
			want: []out{
				{
					id: 1003,
					z:  200,
				},
				{
					id: 1002,
					z:  300,
				},
				{
					id: 1001,
					z:  400,
				},
				{
					id: 999,
					z:  200,
				},
				{
					id: 1007,
					z:  150,
				},
			},
		},
	}
	for _, v := range cases {
		if got := royAndTrendingTopics(v.args.n, v.args.ins); !reflect.DeepEqual(got, v.want) {
			fmt.Printf("args: %+v - got: %v / want: %v\n", v.args, got, v.want)
		}
	}

}
