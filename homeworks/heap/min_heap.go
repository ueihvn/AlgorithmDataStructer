package heap

import (
	"fmt"
)

func minHeapify(nums []int, i int) {
	smallest := i
	left := 2*i + 1
	right := 2*i + 2
	if left < len(nums) && nums[left] < nums[smallest] {
		smallest = left
	}
	if right < len(nums) && nums[right] < nums[smallest] {
		smallest = right
	}
	if smallest == i {
		return
	}
	swap[int](&nums[smallest], &nums[i])
	minHeapify(nums, smallest)
}

func buildMinHeap(n int, nums []int) {
	for i := n / 2; i >= 0; i-- {
		minHeapify(nums, i)
	}
}

func minHeapPop(nums []int) ([]int, int) {
	if len(nums) == 0 {
		return nums, 0
	}
	i := len(nums) - 1
	swap[int](&nums[i], &nums[0])
	removed := nums[i]
	nums = nums[:i]
	minHeapify(nums, 0)
	return nums, removed
}

func minHeapPush(nums []int, v int) []int {
	nums = append(nums, v)
	i := len(nums) - 1
	for i > 0 && nums[(i-1)/2] > nums[i] {
		swap[int](&nums[(i-1)/2], &nums[i])
		i = (i - 1) / 2
	}
	return nums
}

func TestBuildMinHeap() {
	h := []int{7, 12, 6, 10, 17, 15, 2, 4}
	buildMinHeap(len(h), h)
	fmt.Println(h)
	h = minHeapPush(h, 11)
	fmt.Println(h)
	h, removed := minHeapPop(h)
	fmt.Println(h)
	fmt.Println(removed)
}
