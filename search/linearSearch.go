package search

func LinearSearch(nums []int, num int) int {
	for i := range nums {
		if nums[i] == num {
			return i
		} else if nums[i] > num {
			return -1
		}
	}
	return -1
}
