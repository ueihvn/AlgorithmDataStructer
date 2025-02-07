package heap

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
)

func qHeap1(q int, qs [][2]int) []int {
	var res []int
	minHeap := []int{}
	for i := 0; i < q; i++ {
		switch qs[i][0] {
		case 1:
			minHeap = minHeapPush(minHeap, qs[i][1])
		case 2:
			for j, v := range minHeap {
				if v != qs[i][1] {
					continue
				}
				minHeap[j] = minHeap[len(minHeap)-1]
				minHeap = minHeap[:len(minHeap)-1]
				minHeapify(minHeap, j)
				break
			}
		default:
			res = append(res, minHeap[0])
		}
	}
	return res
}

func ReadInputQHeap1() (int, [][2]int) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	q, err := strconv.Atoi(scanner.Text())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading number of lines: %s\n", err)
		return 0, nil
	}

	var res [][2]int

	// Prepare to read the next n lines of data
	for i := 0; i < q && scanner.Scan(); i++ {
		line := scanner.Text()
		parts := strings.Split(line, " ")
		row := make([]int, len(parts))

		for j, part := range parts {
			row[j], err = strconv.Atoi(part)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error converting string to integer: %s\n", err)
				return 0, nil
			}
		}
		if len(parts) == 2 {
			val := [2]int{row[0], row[1]}
			res = append(res, val)
		} else {
			val := [2]int{row[0], 0}
			res = append(res, val)
		}
	}

	// Check for scanning errors
	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading input: %s\n", err)
		return 0, nil
	}
	return q, res
}

func ConsoleQHeap1() {
	q, qs := ReadInputQHeap1()
	res := qHeap1(q, qs)
	for _, v := range res {
		fmt.Println(v)
	}
}

func TestQHeap1() {
	type args struct {
		q  int
		qs [][2]int
	}
	cases := []struct {
		args args
		want []int
	}{
		/*
			{
				args: args{
					q: 5,
					qs: [][2]int{
						{1, 4},
						{1, 9},
						{3, 0},
						{2, 4},
						{3, 0},
					},
				},
				want: []int{4, 9},
			},
		*/
		{
			args: args{
				q: 13,
				qs: [][2]int{
					{1, -890112362},
					{3, 0},
					{2, -890112362},
					{1, 310432408},
					{3, 0},
					{3, 0},
					{3, 0},
					{1, 151263095},
					{3, 0},
					{1, 450851372},
					{1, 189370081},
					{2, 151263095},
					{3, 0},
				},
			},
			want: []int{-890112362, 310432408, 310432408, 310432408, 151263095, 189370081},
		},
	}
	for _, v := range cases {
		if got := qHeap1(v.args.q, v.args.qs); !reflect.DeepEqual(got, v.want) {
			fmt.Printf("args: %+v - got: %v / want: %v\n", v.args, got, v.want)
		}
	}
}
