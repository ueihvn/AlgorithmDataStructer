package main

import "fmt"

func plusOne(digits []int) []int {
	incremenet, i := 1, len(digits)-1

	for incremenet != 0 && i >= 0 {
		temp := digits[i] + incremenet
		if temp >= 10 {
			digits[i] = temp % 10
		} else {
			digits[i] = temp
			incremenet = 0
		}
		i--
	}

	if incremenet != 0 {
		res := make([]int, len(digits)+1)
		res[0] = 1
		for i := 0; i < len(digits); i++ {
			res[i+1] = digits[0]
		}
		return res

	}

	return digits
}

func testPlusOne() {
	testCases := [][]int{
		{1, 2, 3},
		{4, 3, 2, 1},
		{0},
		{9},
	}

	for i := range testCases {
		fmt.Println(testCases[i])
		fmt.Println("plus one: ", plusOne(testCases[i]))
	}
}
