package heap

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
)

func restaurantRating(n int, qs [][2]int) []string {
	var res []string
	var maxHeap []int
	var minHeap []int
	var count int
	for _, q := range qs {
		ratings := count / 3
		switch q[0] {
		case 1:
			count++
			maxHeap = maxHeapPush(maxHeap, q[1])
		default:
			if ratings <= 0 {
				res = append(res, "No reviews yet")
				continue
			}
			i := len(minHeap)
			for i < ratings {
				mxh, removed := maxHeapPop(maxHeap)
				maxHeap = mxh
				minHeap = minHeapPush(minHeap, removed)
				i++
			}
			res = append(res, strconv.Itoa(minHeap[0]))
		}
	}
	return res
}

func ReadInputRestaurantRating() (int, [][2]int) {
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

func ConsoleRestaurantRating() {
	n, qs := ReadInputRestaurantRating()
	res := restaurantRating(n, qs)
	for _, v := range res {
		fmt.Println(v)
	}
}

func TestRestaurantRating() {
	type args struct {
		n  int
		qs [][2]int
	}

	cases := []struct {
		args args
		want []string
	}{
		{
			args: args{
				n: 10,
				qs: [][2]int{
					{1, 1},
					{1, 7},
					{2, 0},
					{1, 9},
					{1, 21},
					{1, 8},
					{1, 5},
					{2, 0},
					{1, 9},
					{2, 0},
				},
			},
			want: []string{"No reviews yet", "9", "9"},
		},
		{
			args: args{
				n: 24,
				qs: [][2]int{
					{1, 970610158},
					{2, 0},
					{1, 355274745},
					{2, 0},
					{1, 533358198},
					{2, 0},
					{1, 453583769},
					{2, 0},
					{1, 858301599},
					{2, 0},
					{1, 857023846},
					{2, 0},
					{1, 197585277},
					{1, 680212757},
					{2, 0},
					{2, 0},
					{2, 0},
					{2, 0},
					{2, 0},
					{1, 279831991},
					{2, 0},
					{1, 277978870},
					{1, 325368879},
					{2, 0},
					{2, 0},
					{1, 274102029},
					{1, 569598161},
					{2, 0},
				},
			},
			want: []string{
				"No reviews yet",
				"No reviews yet",
				"970610158",
				"970610158",
				"970610158",
				"858301599",
				"858301599",
				"858301599",
				"858301599",
				"858301599",
				"858301599",
				"857023846",
				"857023846",
				"857023846",
				"680212757",
			},
		},
	}
	for _, v := range cases {
		if got := restaurantRating(v.args.n, v.args.qs); !reflect.DeepEqual(got, v.want) {
			fmt.Printf("args: %+v - got: %v / want: %v\n", v.args, got, v.want)
		}
	}
}
