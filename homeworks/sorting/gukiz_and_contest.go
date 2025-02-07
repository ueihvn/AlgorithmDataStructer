package sorting

import (
	"fmt"
	"reflect"
	"sort"
)

func gukizAndContest(n int, nums []int) []int {
	boards := make([]int, n)
	copy(boards, nums)
	sort.Slice(boards, func(i, j int) bool {
		return boards[i] > boards[j]
	})
	result := make([]int, 0, n)
	for i := 0; i < n; i++ {
		var count int
		for j := 0; j < n; j++ {
			if boards[j] > nums[i] {
				count++
			} else {
				break
			}
		}
		count++
		result = append(result, count)
	}
	return result
}

func ConsoleGukizAndContest() {
	var n int
	fmt.Scanln(&n)
	nums := make([]int, 0, n)
	var temp int
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &temp)
		nums = append(nums, temp)
	}
	var res string
	list := gukizAndContest(n, nums)
	for i, v := range list {
		if i == len(list)-1 {
			res += fmt.Sprintf("%d", v)
		} else {
			res += fmt.Sprintf("%d ", v)
		}
	}
	fmt.Println(res)
}

func TestGukizAndContest() {
	type args struct {
		n    int
		nums []int
	}

	cases := []struct {
		args args
		want []int
	}{
		{
			args: args{
				n:    3,
				nums: []int{1, 3, 3},
			},
			want: []int{3, 1, 1},
		},
		{
			args: args{
				n:    1,
				nums: []int{1},
			},
			want: []int{1},
		},
		{
			args: args{
				n:    5,
				nums: []int{3, 5, 3, 4, 5},
			},
			want: []int{4, 1, 4, 3, 1},
		},
		{
			args: args{
				n:    11,
				nums: []int{19, 20, 5, 5, 5, 10, 2000, 1999, 2000, 1999, 99},
			},
			want: []int{7, 6, 9, 9, 9, 8, 1, 3, 1, 3, 5},
		},
	}
	for _, v := range cases {
		if got := gukizAndContest(v.args.n, v.args.nums); !reflect.DeepEqual(got, v.want) {
			fmt.Printf("args: %+v - got: %v / want: %v\n", v.args, got, v.want)
		}
	}
}
