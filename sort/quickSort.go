package main

// pation is seperate array into 2 group, gt pivot & lt pivot
// index of pivot is between two group
// all element gt pivot is placed to the right of pivot
// all element lt pivot is placed to the left of pivot
func partion(nums []int, l, r int) int {
	// chosse pivot at r (the last element of array)
	pivot := nums[r]
	i := l - 1
	for j := l; j < r; j++ {
		if nums[j] < pivot {
			i++
			// swap num at i with nums at j
			nums[i], nums[j] = nums[j], nums[i]
		}

	}

	nums[i+1], nums[r] = nums[r], nums[i+1]
	return i + 1
}

// Choose pivot
// partion array -> return index of pivot which nums in left < pivot and nums in right > pivot
// recursie with array in left and in right pivot. Stop if l >= r
func quickSort(nums []int, l, r int) {
	if l >= r {
		return
	}

	// partion array
	p := partion(nums, l, r)

	quickSort(nums, l, p-1)
	quickSort(nums, p+1, r)

}
