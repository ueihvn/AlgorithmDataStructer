package arrays

import (
	"fmt"
)

func consoleBearAndGame() {
	var n int
	fmt.Scanf("%d", &n)
	ims := make([]int, 0, n)
	var im int
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &im)
		ims = append(ims, im)
	}
	fmt.Println(bearAndGame(n, ims))
}

func bearAndGame(n int, interestingMinutes []int) int {
	consecutiveBoringMinutes := 15
	maxGameMinutes := 90
	var start int
	var latestInterestingMinute int
	for i := 0; i < n; i++ {
		cal := interestingMinutes[i] - start
		if cal > consecutiveBoringMinutes {
			if i == 0 {
				return consecutiveBoringMinutes
			}
			latestInterestingMinute = interestingMinutes[i-1]
			break
		}
		// cal <= consecutiveBoringMinutes
		latestInterestingMinute = interestingMinutes[i]
		start = interestingMinutes[i]
	}
	res := latestInterestingMinute + consecutiveBoringMinutes
	if res > maxGameMinutes {
		return maxGameMinutes
	}
	return res
}

func testBearAndGame() {
	type args struct {
		n                  int
		interestingMinutes []int
	}
	type cases struct {
		args args
		want int
	}
	tests := []cases{
		{
			args: args{
				n:                  3,
				interestingMinutes: []int{7, 20, 88},
			},
			want: 35,
		},
		{
			args: args{
				n:                  9,
				interestingMinutes: []int{16, 20, 30, 40, 50, 60, 70, 80, 90},
			},
			want: 15,
		},
		{
			args: args{
				n:                  9,
				interestingMinutes: []int{15, 20, 30, 40, 50, 60, 70, 80, 90},
			},
			want: 90,
		},
		{
			args: args{
				n:                  5,
				interestingMinutes: []int{14, 29, 44, 59, 73},
			},
			want: 88,
		},
	}
	for _, v := range tests {
		if out := bearAndGame(v.args.n, v.args.interestingMinutes); out != v.want {
			fmt.Printf("args: %+v / got: %d / want: %d\n", v.args, out, v.want)
		}
	}
}
