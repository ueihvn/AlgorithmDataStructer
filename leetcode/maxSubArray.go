package main

import "fmt"

// using brute force
func maxSubArray(nums []int) int {
	max := nums[0]
	for i := 0; i < len(nums); i++ {
		maxi, sumi := nums[i], nums[i]
		for j := i + 1; j < len(nums); j++ {
			sumi += nums[j]
			if maxi < sumi {
				maxi = sumi
			}
		}
		if max < maxi {
			max = maxi
		}
	}

	return max
}

// using kadane algorithm
//[1,-3,2]
// maximum subarray of subarray at index 0 = [1]
// maximum subarray of subarray at index 1 = [-3; [1,-3]]
// maximum subarray of subarray at index 2 = [2 | [2,-3] | [2,-3,1]]
// -> maximum subarray at index = [nums at index | num at index + maximum subarray at previous index]
// -> maximum subarray of entire array = maximum subarray at index has  maximum subarray = max[maximum subarray at index [0 -> n-1]]
// ex: max_at_index_2 = [2 |2 +  max_at_1] = [2 , 1,-3]]
func maxSubArray2(nums []int) int {
	// max_ending_here contain max with current i
	max_ending_here, max_so_far := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		max_ending_here = max(nums[i], max_ending_here+nums[i])
		max_so_far = max(max_ending_here, max_so_far)
	}
	return max_so_far
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func testMaxSubArray() {
	arrs := [][]int{
		{-2, 1, -3, 4, -1, 2, 1, -5, 4},
		{1},
		{5, 4, -1, 7, 8},
	}

	for i := range arrs {
		fmt.Println(arrs[i])
		fmt.Println(maxSubArray2(arrs[i]))
	}
}
