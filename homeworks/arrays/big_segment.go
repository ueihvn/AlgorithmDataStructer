package arrays

import "fmt"

type segment struct {
	l int
	r int
}

func consoleBigSegment() {
	var n int
	fmt.Scanf("%d", &n)
	segments := make([]segment, 0, n)
	var l, r int
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &l)
		fmt.Scanf("%d", &r)
		segments = append(segments, segment{l: l, r: r})
	}
	fmt.Println(bigSegment(n, segments))
}

func bigSegment(n int, segments []segment) int {
	minL := segments[0].l
	maxR := segments[0].r
	for i := 1; i < n; i++ {
		if segments[i].l < minL {
			minL = segments[i].l
		}
		if segments[i].r > maxR {
			maxR = segments[i].r
		}
	}
	for i := 0; i < n; i++ {
		if segments[i].r == maxR && segments[i].l == minL {
			return i + 1
		}
	}
	return -1
}

func testBigSegment() {
	type args struct {
		n        int
		segments []segment
	}
	type test struct {
		args args
		want int
	}
	tests := []test{
		{
			args: args{
				n: 3,
				segments: []segment{
					{l: 3, r: 3},
					{l: 4, r: 4},
					{l: 5, r: 5},
				},
			},
			want: -1,
		},
		{
			args: args{
				n: 6,
				segments: []segment{
					{l: 1, r: 5},
					{l: 2, r: 3},
					{l: 1, r: 10},
					{l: 7, r: 10},
					{l: 7, r: 7},
					{l: 10, r: 10},
				},
			},
			want: 3,
		},
	}
	for _, v := range tests {
		if out := bigSegment(v.args.n, v.args.segments); out != v.want {
			fmt.Printf("args: %+v - out: %d / expected: %d\n", v.args, out, v.want)
		}
	}
}
