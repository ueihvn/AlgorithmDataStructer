package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generateSliceInt(n int, max int) []int {
	rand.Seed(time.Now().UnixNano())
	nums := make([]int, n)
	for i := range nums {
		nums[i] = rand.Intn(max)
	}

	return nums
}

func main() {
	lenList := 30
	maxNum := 100
	nums := generateSliceInt(lenList, maxNum)
	fmt.Println("Array nums:\n", nums)
	// InterchangeSort(nums, lenList)
	BubbleSort(nums, lenList)
	fmt.Println("Sorted Array nums:\n", nums)
}
