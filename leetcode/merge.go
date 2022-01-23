package main

import "fmt"

func merge(nums1 []int, m int, nums2 []int, n int) []int {
	// i is index for nums1
	// j is index for nums2
	// z is index for temp
	i, j, z := 0, 0, 0
	temp := make([]int, m)

	for z < m+n {
		temp[z] = nums1[z]
		if (i < m && j < n && temp[i] < nums2[j]) || j >= n {
			nums1[z] = temp[i]
			i++
		} else {
			nums1[z] = nums2[j]
			j++
		}
		z++
	}

	return nums1
}

func testMerge() {
	inputs := []struct {
		nums1 []int
		m     int
		nums2 []int
		n     int
	}{}

	for _, v := range inputs {
		fmt.Println(v.nums1, v.m)
		fmt.Println(v.nums2, v.n)
		res := merge(v.nums1, v.m, v.nums2, v.n)
		fmt.Println(res)
	}
}
