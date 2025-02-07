package sorting

import (
	"fmt"
	"sort"
)

func eightPointSets(sets [8][2]int) string {
	x := make([]int, 0, 8)
	y := make([]int, 0, 8)
	var countY, countX int
	var maxY, maxX int
	//for _, set := range sets {
	//	if set[0] == 0 {
	//		countX++
	//		if set[1] > maxY {
	//			maxY = set[1]
	//		}
	//	}
	//	if set[1] == 0 {
	//		countY++
	//		if set[0] > maxX {
	//			maxX = set[0]
	//		}
	//	}
	//	x = append(x, set[0])
	//	y = append(y, set[1])
	//}
	if countY < 3 || countX < 3 {
		return "ugly"
	}
	sort.Ints(x)
	sort.Ints(y)
	if x[7] != maxX {
		return "ugly"
	}
	if y[7] != maxY {
		return "ugly"
	}
	return "respectable"
}

func ConsoleEightPointSets() {
	sets := [8][2]int{}
	for i := 0; i < 8; i++ {
		var x int
		fmt.Scanf("%d", &x)
		sets[i][0] = x
		var y int
		fmt.Scanf("%d", &y)
		sets[i][1] = y
	}
	fmt.Println(eightPointSets(sets))
}

func TestEightPointSets() {
	type args struct {
		sets [8][2]int
	}
	cases := []struct {
		args args
		want string
	}{
		{
			args: args{
				sets: [8][2]int{
					{0, 0},
					{0, 1},
					{0, 2},
					{1, 0},
					{1, 2},
					{2, 0},
					{2, 1},
					{2, 2},
				},
			},
			want: "respectable",
		},
		{
			args: args{
				sets: [8][2]int{
					{0, 0},
					{1, 0},
					{2, 0},
					{3, 0},
					{4, 0},
					{5, 0},
					{6, 0},
					{7, 0},
				},
			},
			want: "ugly",
		},
		{
			args: args{
				sets: [8][2]int{
					{1, 1},
					{1, 2},
					{1, 3},
					{2, 1},
					{2, 2},
					{2, 3},
					{3, 1},
					{3, 2},
				},
			},
			want: "ugly",
		},
	}
	for _, v := range cases {
		if got := eightPointSets(v.args.sets); got != v.want {
			fmt.Printf("args: %+v - got: %v / want: %v\n", v.args, got, v.want)
		}
	}
}
