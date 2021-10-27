package main

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	k := 1
	numUnique := nums[0]

	for i := 1; i < len(nums); i++ {
		if nums[i] != numUnique {
			nums[k] = nums[i]
			k++
			numUnique = nums[i]
		}
	}

	for j := k; j < len(nums); j++ {
		nums[j] = 0
	}

	return k
}
