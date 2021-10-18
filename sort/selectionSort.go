package main

import (
	"fmt"
	"time"
)

// sort by putting element in-place
// find index of smallest element from [i, n-1]
// swap smallest element with element at index i
func selectionSort(nums []int) {
	start := time.Now()

	len := len(nums)
	for i := 0; i <= len-2; i++ {
		lc := i
		for j := i; j <= len-1; i++ {
			if nums[lc] < nums[j] {
				lc = j
			}
		}
		// swap
		nums[i], nums[lc] = nums[lc], nums[i]
	}
	fmt.Printf("\nSelection Sort\nTime: %v\n", time.Since(start))

}
