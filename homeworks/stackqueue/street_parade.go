package stackqueue

import (
	"fmt"
)

func StreetParade(n int, mobiles []int) string {
	stack := NewStack[int]()
	ordered := make([]int, 0, n)
	for i := 0; i < n-1; i++ {
		if mobiles[i] <= mobiles[i+1] {
			if stack.IsEmpty() {
				ordered = append(ordered, mobiles[i])
				continue
			}
			if mobiles[i] <= stack.Top() {
				ordered = append(ordered, mobiles[i])
			} else {
				ordered = append(ordered, stack.Top())
				stack.Pop()
			}
			continue
		}
		if !stack.IsEmpty() {
			if mobiles[i] >= stack.Top() {
				return "no"
			}
		}
		stack.Push(mobiles[i])
	}
	//var isNMinusOne bool
	//for !stack.IsEmpty() {
	//	if !isNMinusOne && mobiles[n-1] <= stack.Top() {
	//		ordered = append(ordered, mobiles[n-1])
	//		isNMinusOne = true
	//		continue
	//	}
	//	ordered = append(ordered, stack.Top())
	//	stack.Pop()
	//}
	//if len(ordered) == len(mobiles) {
	//	return "yes"
	//}
	//return "no"
	return "yes"
}

func ConsoleStreetParade() {
	for {
		var t int
		fmt.Scanln(&t)
		if t == 0 {
			break
		}
		mobiles := make([]int, 0, t)
		for i := 0; i < t; i++ {
			var temp int
			fmt.Scanf("%d", &temp)
			mobiles = append(mobiles, temp)
		}
		fmt.Println(StreetParade(t, mobiles))
	}
}

func TestStreetParade() {
	type args struct {
		n       int
		mobiles []int
	}
	cases := []struct {
		args args
		want string
	}{
		{
			args: args{
				n:       5,
				mobiles: []int{5, 1, 2, 4, 3},
			},
			want: "yes",
		},
		{
			args: args{
				n:       20,
				mobiles: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20},
			},
			want: "yes",
		},
	}
	for _, v := range cases {
		if got := StreetParade(v.args.n, v.args.mobiles); got != v.want {
			fmt.Printf("args: %+v - got: %v / want: %v\n", v.args, got, v.want)
		}
	}
}
