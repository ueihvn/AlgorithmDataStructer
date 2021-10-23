package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/ueihvn/AlgorithmDataStructer/search"
)

// function return slice of int have length = n, and < max
func generateSliceInt(n int, max int) []int {
	rand.Seed(time.Now().UnixNano())
	nums := make([]int, n)
	for i := range nums {
		nums[i] = rand.Intn(max)
	}

	return nums
}

func main() {
	nums := generateSliceInt(5, 100)
	fmt.Println("Array nums:\n", nums)

	// bubleSortAsc(nums)
	// fmt.Println("Array nums Asc:\n", nums)

	// bubleSortDesc(nums)
	// fmt.Println("Array nums Desc:\n", nums)

	// insertionSortAsc(nums)
	// fmt.Println("Array nums After Sort:\n", nums)

	// insertionSortDesc(nums)
	// fmt.Println("Array nums After Sort:\n", nums)

	// insertionSort(nums)
	// fmt.Println("Array nums After Sort:\n", nums)

	// quickSort(nums, 0, 19)
	// fmt.Println("Array nums After Sort:\n", nums)

	mergeSort(nums, 0, len(nums)-1)
	fmt.Println("Array nums After Sort:\n", nums)
	fmt.Println(search.LinearSearch(nums, 22))
}
