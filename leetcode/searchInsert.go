package main

import "fmt"

func searchInsert(nums []int, target int) int {
	left, right := 0, len(nums)-1
	mid := 0
	for left <= right {
		mid = (left + right) / 2
		if nums[mid] == target {
			fmt.Printf("nums: %d, target: %d\n", nums, target)
			fmt.Printf("left: %d, mid: %d, right: %d\n", left, mid, right)
			return mid
		} else if target < nums[mid] {
			right = mid - 1
		} else if target > nums[mid] {
			left = mid + 1
		}
	}
	fmt.Printf("nums: %d, target: %d\n", nums, target)
	fmt.Printf("left: %d, mid: %d, right: %d\n", left, mid, right)
	// return left
	// because else if target < nums[mid] before -> right = res -1, now left = right = res
	// next loop right < left then break for loop
	// now right = res - 1, left = res
	return left
}

func testSearchInsert() {
	nums := []int{1, 3, 5, 6}
	targets := []int{5, 2, 7, 0}

	for i := range targets {
		fmt.Println(searchInsert(nums, targets[i]))
		fmt.Println()
	}
}
