package main

import "fmt"

func mySqrt(x int) int {
	left, right, mid := 0, x, 0

	for left <= right {
		mid = (left + right) / 2
		temp := mid * mid
		if temp > x {
			right = mid - 1
		} else if temp < x {
			left = mid + 1
		} else {
			return mid
		}
	}

	return right
}

func testMySqrt() {
	input := []int{4, 8}
	for i := range input {
		fmt.Println(input[i], " sqrt() = ", mySqrt(input[i]))
	}
}
