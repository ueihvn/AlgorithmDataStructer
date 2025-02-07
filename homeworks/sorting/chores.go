package sorting

import (
	"fmt"
	"sort"
)

func chores(n, a, b int, h []int) int {
	sort.Ints(h)
	// asc [0;n-1]
	// a num > x
	// b num <= x
	// a = 2, b = 3
	// 1 2 3 6 100
	var res int
	start := h[b-1]
	for i := 0; start < h[b]; i++ {
		if start+i < h[b] {
			res++
		} else {
			break
		}
	}
	return res
}

func ConsoleChores() {
	var n int
	fmt.Scanf("%d", &n)
	var a int
	fmt.Scanf("%d", &a)
	var b int
	fmt.Scanf("%d", &b)
	h := make([]int, 0, n)
	var temp int
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &temp)
		h = append(h, temp)
	}
	fmt.Println(chores(n, a, b, h))

}

func TestChores() {
	type args struct {
		n int
		a int
		b int
		h []int
	}

	cases := []struct {
		args args
		want int
	}{
		{
			args: args{
				n: 5,
				a: 2,
				b: 3,
				h: []int{6, 2, 3, 100, 1},
			},
			want: 3,
		},
		{
			args: args{
				n: 7,
				a: 3,
				b: 4,
				h: []int{1, 1, 9, 1, 1, 1, 1},
			},
			want: 0,
		},
	}
	for _, v := range cases {
		if got := chores(v.args.n, v.args.a, v.args.b, v.args.h); got != v.want {
			fmt.Printf("args: %+v - got: %d / want: %d\n", v.args, got, v.want)
		}
	}
}
