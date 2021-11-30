package search

func BinarySearch(nums []int, num int) int {
	left, mid, right := 0, 0, len(nums)-1
	for left <= right {
		mid = (left + right) / 2
		if nums[mid] == num {
			return mid
		} else if nums[mid] > num {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}
