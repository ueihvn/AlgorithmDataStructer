package arrays

import (
	"fmt"
)

func consoleNightAtMuseum() {
	var exhibit string
	fmt.Scanln(&exhibit)
	fmt.Println(nightAtMuseum(exhibit))
}

func nightAtMuseum(exhibit string) int {
	// a -> z () clockwise
	// z -> a () counterClockwise
	start := byte('a')
	var result int
	for i := 0; i < len(exhibit); i++ {
		cw := clockwiseDistance(start, exhibit[i])
		ccw := counterClockwiseDistance(start, exhibit[i])
		if cw < ccw {
			result += cw
		} else {
			result += ccw
		}
		start = exhibit[i]
	}
	return result
}

func clockwiseDistance(from, to byte) int {
	max := byte('z')
	min := byte('a')
	if to < from {
		// b -> a
		// c -> a
		return int(max-from+to-min) + 1
	} else if to > from {
		// a -> b
		return int(to - from)
	}
	return 0
}

func counterClockwiseDistance(from, to byte) int {
	max := byte('z')
	min := byte('a')
	if to < from {
		// b -> a
		// c -> a
		return int(from - to)
	} else if to > from {
		// c -> a
		return int(from-min+max-to) + 1
	}
	return 0
}

func testNightAtMuseum() {
	type args struct {
		in  string
		out int
	}
	tests := []args{
		{
			in:  "ares",
			out: 34,
		},
		{
			in:  "zeus",
			out: 18,
		},
		{
			in:  "map",
			out: 35,
		},
	}
	for _, v := range tests {
		if out := nightAtMuseum(v.in); out != v.out {
			fmt.Printf("in: %s - out: %d / expected: %d", v.in, out, v.out)
		}
	}
}
