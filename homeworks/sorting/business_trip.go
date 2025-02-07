package sorting

import (
	"fmt"
	"sort"
)

func businessTrip(k int, nums []int) int {
	if k == 0 {
		return 0
	}
	sort.Ints(nums)
	var total, count int
	for i := 11; i >= 0; i-- {
		total += nums[i]
		count++
		if total >= k {
			return count
		}
	}
	return -1
}

func ConsoleBusinessTrip() {
	var k int
	fmt.Scanln(&k)
	nums := make([]int, 0, 12)
	var temp int
	for i := 0; i < 12; i++ {
		fmt.Scanf("%d", &temp)
		nums = append(nums, temp)
	}
	fmt.Println(businessTrip(k, nums))
}

func TestBusinessTrip() {
	type args struct {
		k    int
		nums []int
	}

	cases := []struct {
		args args
		want int
	}{
		{
			args: args{
				k:    5,
				nums: []int{1, 1, 1, 1, 2, 2, 3, 2, 2, 1, 1, 1},
			},
			want: 2,
		},
		{
			args: args{
				k:    0,
				nums: []int{0, 0, 0, 0, 0, 0, 0, 1, 1, 2, 3, 0},
			},
			want: 0,
		},
		{
			args: args{
				k:    11,
				nums: []int{1, 1, 4, 1, 1, 5, 1, 1, 4, 1, 1, 1},
			},
			want: 3,
		},
	}
	for _, v := range cases {
		if got := businessTrip(v.args.k, v.args.nums); got != v.want {
			fmt.Printf("args: %+v - got: %v / want: %v\n", v.args, got, v.want)
		}
	}
}
