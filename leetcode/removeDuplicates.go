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

// move first appear element  to begin of nums
// 0 1 2 3 3 3
// i
//	 j
//         k
//		     j
// numU = 3
func removeDuplicates2(nums []int) int {
	i, j := 0, 1

	for j < len(nums) {
		if nums[i] != nums[j] {
			i++
			nums[i] = nums[j]
		}
		j++
	}
	return i + 1
}
