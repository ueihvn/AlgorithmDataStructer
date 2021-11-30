package main

import "fmt"

// 1 -> 1
// 2 -> 2  // 1 + 1 | 2
// 3 -> 3 // 1 + 1 + 1 | 1 + 2 | 2 + 1
// 4 -> 5 // 1 + 1 + 1 + 1 | 1 + 1 + 2 | 1 + 2 + 1 | 2 + 1 + 1 | 2 + 2
// just like fibonanci

func climbStairs(n int) int {
	if n == 1 || n == 2 {
		return n
	}

	a, b := 1, 2
	var temp int
	for i := 3; i <= n; i++ {
		temp = a + b
		a = b
		b = temp
	}

	return temp
}

// easy to understand but take longer
func climbStairsRecursive(n int) int {
	if n <= 2 {
		return n
	}

	return climbStairs(n-1) + climbStairs(n-2)
}

func testClimbStairs() {
	inputs := []int{2, 3, 4, 45}
	for _, v := range inputs {
		fmt.Println("n: ", v, "climbStair(n): ", climbStairs(v))
		fmt.Println("n: ", v, "climbStairRecursive(n): ", climbStairsRecursive(v))
	}
}
