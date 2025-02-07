package binarysearch

import "sort"

func whereIsTheMarble(n, q int, marbles, qs []int) []string {
	var res []string
	sort.Ints(marbles)
	for _, query := range qs {
		i := searchFirst(query, marbles)
		if i == -1 {
		}
	}
	return res
}

func searchFirst(n int, nums []int) int {
	l := 0
	r := len(nums) - 1
	for l <= r {
		mid := l + (r-l)/2
		if nums[mid] == n && (mid == l || n != nums[mid-1]) {
			return mid
		} else if nums[mid] < n {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return -1
}
