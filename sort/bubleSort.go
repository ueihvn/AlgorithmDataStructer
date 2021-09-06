package main

import (
	"fmt"
	"time"
)

// Time complextity -> O(n^2)
// suitable for small array
func bubleSortAsc(nums []int) {
	start := time.Now()

	len := len(nums)
	for i := 0; i < len-1; i++ {
		// swap the lagrest num to the ends
		// - i -> every i loop, the largest num to the ends loop
		// Don't need compare it in next i loop
		// Ex: i = 0, first i loop the largest num at len-1.
		// i = 1, largest at len -1 -> don't need compare => j < len -1 - 1
		for j := 0; j < len-i-1; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
	fmt.Printf("\nBuble Sort\nTime: %v\n", time.Since(start))

}

func bubleSortDesc(nums []int) {
	start := time.Now()
	len := len(nums)
	for i := 0; i < len-1; i++ {
		for j := 0; j < len-i-1; j++ {
			if nums[j] < nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
	fmt.Printf("\nBuble Sort\nTime: %v\n", time.Since(start))

}
