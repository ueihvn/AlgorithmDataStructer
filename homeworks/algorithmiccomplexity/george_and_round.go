package algorithmiccomplexity

import "fmt"

func ConsoleGeorgeAndRound() {
	var n int
	fmt.Scanf("%d", &n)
	var m int
	fmt.Scanf("%d", &m)
	a := make([]int, 0, n)
	var temp int
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &temp)
		a = append(a, temp)
	}
	b := make([]int, 0, m)
	for i := 0; i < m; i++ {
		fmt.Scanf("%d", &temp)
		b = append(b, temp)
	}
	fmt.Println(georgeAndRound(n, m, a, b))
}

func georgeAndRound(n, m int, a, b []int) int {
	var i, j, count int
	for j < m && i < n {
		if a[i] > b[j] {
			j++
			continue
		}
		// a[i] <= b[j]
		count++
		i++
		j++
	}
	return n - count
}

func TestGeorgeAndRound() {
	type args struct {
		n int
		m int
		a []int
		b []int
	}
	type test struct {
		args args
		want int
	}

	cases := []test{
		{
			args: args{
				n: 3,
				m: 5,
				a: []int{1, 2, 3},
				b: []int{1, 2, 2, 3, 3},
			},
			want: 0,
		},
		{
			args: args{
				n: 3,
				m: 5,
				a: []int{1, 2, 3},
				b: []int{1, 1, 1, 1, 1},
			},
			want: 2,
		},
		{
			args: args{
				n: 3,
				m: 1,
				a: []int{2, 3, 4},
				b: []int{1},
			},
			want: 3,
		},
	}

	for _, v := range cases {
		if got := georgeAndRound(v.args.n, v.args.m, v.args.a, v.args.b); got != v.want {
			fmt.Printf("args: %+v - got: %d / want: %d\n", v.args, got, v.want)
		}
	}
}
