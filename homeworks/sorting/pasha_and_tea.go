package sorting

import (
	"fmt"
	"math"
	"sort"
)

func pashaAndTea(n, w int, a []int) float64 {
	sort.Ints(a)
	amount := math.Min(float64(a[0]), float64(a[n])/2)
	return math.Min(float64(w), float64(3*n)*amount)
}

func ConsolePashaAndTea() {
	var n int
	fmt.Scanf("%d", &n)
	var w int
	fmt.Scanf("%d", &w)
	a := make([]int, 0, 2*n)
	var temp int
	for i := 0; i < 2*n; i++ {
		fmt.Scanf("%d", &temp)
		a = append(a, temp)
	}
	fmt.Println(pashaAndTea(n, w, a))
}

func TestPashaAndTea() {
	type args struct {
		n int
		w int
		a []int
	}

	cases := []struct {
		args args
		want float64
	}{
		{
			args: args{
				n: 2,
				w: 4,
				a: []int{1, 1, 1, 1},
			},
			want: 3,
		},
		{
			args: args{
				n: 3,
				w: 18,
				a: []int{4, 4, 4, 2, 2, 2},
			},
			want: 18,
		},
		{
			args: args{
				n: 1,
				w: 5,
				a: []int{2, 3},
			},
			want: 4.5,
		},
	}
	for _, v := range cases {
		if got := pashaAndTea(v.args.n, v.args.w, v.args.a); got != v.want {
			fmt.Printf("args: %+v - got: %v / want: %v\n", v.args, got, v.want)
		}
	}
}
