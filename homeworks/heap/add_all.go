package heap

import (
	"fmt"
	"reflect"
)

func addAll(n int, nums []int) int {
	mh := make([]int, n)
	copy(mh, nums)
	buildMinHeap(n, mh)
	var res int
	for len(mh) != 0 {
		mh1, first := minHeapPop(mh)
		mh2, second := minHeapPop(mh1)
		num := first + second
		res += first + second
		if len(mh2) != 0 {
			mh = minHeapPush(mh2, num)
		} else {
			break
		}
	}
	return res
}

func ConsoleAddAll() {
	for {
		var n int
		fmt.Scanln(&n)
		if n == 0 {
			break
		}
		nums := make([]int, 0, n)
		var temp int
		for i := 0; i < n; i++ {
			fmt.Scanf("%d", &temp)
			nums = append(nums, temp)
		}
		fmt.Println(addAll(n, nums))
	}
}

func TestAddAll() {
	type args struct {
		n    int
		nums []int
	}
	cases := []struct {
		args args
		want int
	}{
		{
			args: args{
				n:    3,
				nums: []int{1, 2, 3},
			},
			want: 9,
		},
		{
			args: args{
				n:    4,
				nums: []int{1, 2, 3, 4},
			},
			want: 19,
		},
	}
	for _, v := range cases {
		if got := addAll(v.args.n, v.args.nums); !reflect.DeepEqual(got, v.want) {
			fmt.Printf("args: %+v - got: %v / want: %v\n", v.args, got, v.want)
		}
	}
}
