package main

import (
	"fmt"
	"time"
)

func insertionSort(nums []int) {
	start := time.Now()

	len := len(nums)
	for i := 1; i < len; i++ {
		for position := i; position > 0; position-- {
			curr := nums[position]

			if curr > nums[position-1] {
				break
			} else {
				nums[position] = nums[position-1]
				nums[position-1] = curr
			}
		}
	}
	fmt.Printf("\nInsertion Sort\nTime: %v\n", time.Since(start))

}
