package heap

import (
	"fmt"
	"reflect"
)

func monkAndMultiplication(n int, a []int) []int {
	res := make([]int, 0, n)
	maxH := []int{}
	for i := 0; i < n; i++ {
		if i < 2 {
			res = append(res, -1)
			maxH = maxHeapPush(maxH, a[i])
			continue
		} else if i == 2 {
			res = append(res, a[0]*a[1]*a[2])
			maxH = maxHeapPush(maxH, a[i])
			continue
		}
		maxH = maxHeapPush(maxH, a[i])
		max1, first := maxHeapPop(maxH)
		max2, second := maxHeapPop(max1)
		max3, third := maxHeapPop(max2)
		res = append(res, first*second*third)
		maxH = maxHeapPush(max3, third)
		maxH = maxHeapPush(maxH, second)
		maxH = maxHeapPush(maxH, first)
	}
	return res
}

func ConsoleMonkAndMultiplication() {
	var n int
	fmt.Scanln(&n)
	a := make([]int, 0, n)
	var temp int
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &temp)
		a = append(a, temp)
	}
	res := monkAndMultiplication(n, a)
	for _, v := range res {
		fmt.Println(v)
	}
}

func TestMonkAndMultiplication() {
	type args struct {
		n int
		a []int
	}

	cases := []struct {
		args args
		want []int
	}{
		{
			args: args{
				n: 5,
				a: []int{1, 2, 3, 4, 5},
			},
			want: []int{-1, -1, 6, 24, 60},
		},
	}
	for _, v := range cases {
		if got := monkAndMultiplication(v.args.n, v.args.a); !reflect.DeepEqual(got, v.want) {
			fmt.Printf("args: %+v - got: %v / want: %v\n", v.args, got, v.want)
		}
	}

}
