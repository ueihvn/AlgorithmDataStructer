package sorting

import (
	"fmt"
	"sort"
)

func devuTheDumbGuy(n, x int, c []int) int {
	sort.Ints(c)
	var res int
	for i := 0; i < n; i++ {
		if x <= 1 {
			res += c[i]
		} else {
			res += c[i] * x
			x--
		}
	}
	return res
}

func ConsoleDevuTheDumbGuy() {
	var n int
	fmt.Scanf("%d", &n)
	var x int
	fmt.Scanf("%d", &x)
	c := make([]int, 0, n)
	var temp int
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &temp)
		c = append(c, temp)
	}
	fmt.Println(devuTheDumbGuy(n, x, c))
}

func TestDevuTheDumbGuy() {
	type args struct {
		n int
		x int
		c []int
	}

	cases := []struct {
		args args
		want int
	}{
		{
			args: args{
				n: 2,
				x: 3,
				c: []int{4, 1},
			},
			want: 11,
		},
		{
			args: args{
				n: 4,
				x: 2,
				c: []int{5, 1, 2, 1},
			},
			want: 10,
		},
		{
			args: args{
				n: 3,
				x: 3,
				c: []int{1, 1, 1},
			},
			want: 6,
		},
	}
	for _, v := range cases {
		if got := devuTheDumbGuy(v.args.n, v.args.x, v.args.c); got != v.want {
			fmt.Printf("args: %+v - got: %d / want: %d\n", v.args, got, v.want)
		}
	}
}
