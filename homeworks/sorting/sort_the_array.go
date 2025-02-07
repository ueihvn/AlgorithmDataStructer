package sorting

import (
	"fmt"
	"reflect"
	"sort"
)

func sortTheArray(n int, nums []int) []string {
	dst := make([]int, n)
	copy(dst, nums)
	sort.Ints(nums)
	var count, r, l int
	diff := make([]int, 2)
	for i, v := range nums {
		if v != dst[i] {
			if count == 0 {
				r = i
			}
			count++
			l = i
		}
	}
	if count == 0 {
		diff = []int{1, 1}
	} else {
		diff[0] = r + 1
		diff[1] = l + 1
		for i := 0; i <= l-r; i++ {
			if nums[r+i] != dst[l-i] {
				return []string{"no"}
			}
		}
		sort.Ints(diff)
	}
	return []string{"yes", fmt.Sprintf("%d %d", diff[0], diff[1])}
}

func ConsoleSortTheArray() {
	var n int
	fmt.Scanf("%d", &n)
	nums := make([]int, 0, n)
	var temp int
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &temp)
		nums = append(nums, temp)
	}
	res := sortTheArray(n, nums)
	if len(res) == 1 {
		fmt.Println("no")
	} else {
		for _, v := range res {
			fmt.Println(v)
		}
	}
}

func TestSortTheArray() {
	type args struct {
		n    int
		nums []int
	}

	cases := []struct {
		args args
		want []string
	}{
		{
			args: args{
				n:    3,
				nums: []int{3, 2, 1},
			},
			want: []string{"yes", "1 3"},
		},
		{
			args: args{
				n:    4,
				nums: []int{2, 1, 3, 4},
			},
			want: []string{"yes", "1 2"},
		},
		{
			args: args{
				n:    4,
				nums: []int{3, 1, 2, 4},
			},
			want: []string{"no"},
		},
		{
			args: args{
				n:    2,
				nums: []int{1, 2},
			},
			want: []string{"yes", "1 1"},
		},
	}
	for _, v := range cases {
		if got := sortTheArray(v.args.n, v.args.nums); !reflect.DeepEqual(got, v.want) {
			fmt.Printf("args: %+v - got: %v / want: %v\n", v.args, got, v.want)
		}
	}
}
