package main

import (
	"fmt"
	"time"
)

// seperated array into sorted list (left of position) and unsorted list(rigt of position)
// every i loop, insert position into sorted list
func insertionSortAsc(nums []int) {
	start := time.Now()

	len := len(nums)
	for i := 1; i < len; i++ {
		position := i
		curr := nums[position]
		for position >= 0 {
			if position == 0 {
				nums[0] = curr
				break
			}
			// curr > nums at index position -1 => curr > nums at index < position -1(insert curr at position)
			if curr > nums[position-1] {
				nums[position] = curr
				break
			} else {
				// curr < num[position-1] => find index of curr in sorted list
				// move to position, update position -1
				nums[position] = nums[position-1]
				position--
			}
		}

	}
	fmt.Printf("\nInsertion Sort Asc\nTime: %v\n", time.Since(start))
}

func insertionSortDesc(nums []int) {
	start := time.Now()

	len := len(nums)
	for i := 1; i < len; i++ {
		position := i
		curr := nums[position]
		for position >= 0 {
			if position == 0 {
				nums[0] = curr
				break
			}
			// curr < nums at index position -1 => curr < nums at index < position -1(insert curr at position)
			if curr < nums[position-1] {
				nums[position] = curr
				break
			} else {
				// curr > num[position-1] => find index of curr in sorted list
				// move to position, update position -1
				nums[position] = nums[position-1]
				position--
			}
		}

	}
	fmt.Printf("\nInsertion Sort Desc\nTime: %v\n", time.Since(start))
}

// O(n^2)
// do n time.Each time can do n compare
// Best case O(n)
// compare is more complex than swap(two pointer)
// using binary search with nums[0,key-1] -> nlog(n) => binary insertion Sort
func insertionSort(nums []int) {
	start := time.Now()

	len := len(nums)
	for i := 1; i < len; i++ {
		position := i
		curr := nums[position]
		for position > 0 {
			// curr > nums at index position -1 => curr > nums at index < position -1(insert curr at position)
			if curr > nums[position-1] {
				nums[position] = curr
				break
			} else {
				// curr < num[position-1] => find index of curr in sorted list
				// swap 2 pair of nums to find index of nums[i]
				nums[position], nums[position-1] = nums[position-1], nums[position]
				position--
			}
		}

	}
	fmt.Printf("\nInsertion Sort\nTime: %v\n", time.Since(start))
}
