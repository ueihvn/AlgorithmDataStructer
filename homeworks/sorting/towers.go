package sorting

import (
	"fmt"
	"sort"
)

func towers(n int, l []int) (int, int) {
	fre := make([]int, 1000+1)
	sort.Ints(l)
	prev := 0
	maxHeight, count := 1, 1
	fre[l[prev]]++
	for i := 1; i < n; i++ {
		if l[i] == l[prev] {
			count++
			if maxHeight < count {
				maxHeight = count
			}
		} else {
			count = 1
		}
		fre[l[i]]++
		prev++
	}
	var totalTower int
	for _, v := range fre {
		if v > 0 {
			totalTower++
		}
	}
	return maxHeight, totalTower
}

func ConsoleTowers() {
	var n int
	fmt.Scanln(&n)
	l := make([]int, 0, n)
	var temp int
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &temp)
		l = append(l, temp)
	}
	max, total := towers(n, l)
	fmt.Printf("%d %d\n", max, total)
}

func TestTowers() {
	type args struct {
		n int
		l []int
	}

	cases := []struct {
		args  args
		want1 int
		want2 int
	}{
		{
			args: args{
				n: 3,
				l: []int{1, 2, 3},
			},
			want1: 1,
			want2: 3,
		},
		{
			args: args{
				n: 4,
				l: []int{6, 5, 6, 7},
			},
			want1: 2,
			want2: 3,
		},
	}

	for _, v := range cases {
		if got1, got2 := towers(v.args.n, v.args.l); got1 != v.want1 || got2 != v.want2 {
			fmt.Printf("args: %+v - got: %d %d / want: %d %d\n", v.args, got1, got2, v.want1, v.want2)
		}
	}
}
