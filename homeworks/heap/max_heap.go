package heap

import "fmt"

func maxHeapify(nums []int, i int) {
	greatest := i
	left := 2*i + 1
	right := 2*i + 2
	if left < len(nums) && nums[left] > nums[greatest] {
		greatest = left
	}
	if right < len(nums) && nums[right] > nums[greatest] {
		greatest = right
	}
	if greatest == i {
		return
	}
	swap[int](&nums[greatest], &nums[i])
	maxHeapify(nums, greatest)
}

func buildMaxHeap(n int, nums []int) {
	for i := n / 2; i >= 0; i-- {
		maxHeapify(nums, i)
	}
}

func maxHeapPop(nums []int) ([]int, int) {
	if len(nums) == 0 {
		return nums, 0
	}
	i := len(nums) - 1
	swap[int](&nums[i], &nums[0])
	removed := nums[i]

	nums = nums[:i]
	maxHeapify(nums, 0)
	return nums, removed
}

func maxHeapPush(nums []int, v int) []int {
	if len(nums) == 0 {
		return append(nums, v)
	}
	nums = append(nums, v)
	i := len(nums) - 1
	for i > 0 && nums[(i-1)/2] < nums[i] {
		swap[int](&nums[(i-1)/2], &nums[i])
		i = (i - 1) / 2
	}
	return nums
}

func TestBuildMaxHeap() {
	h := []int{7, 12, 6, 10, 17, 15, 2, 4}
	buildMaxHeap(len(h), h)
	fmt.Println(h)
	h = maxHeapPush(h, 11)
	h, removed := maxHeapPop(h)
	fmt.Println(h)
	fmt.Println(removed)
}
