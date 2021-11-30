package main

func removeElement(nums []int, val int) int {
	left, right, numsVal := 0, len(nums)-1, 0
	for left <= right && right < len(nums) {
		for left <= right && nums[left] != val {
			left++
		}
		for right >= left && nums[right] == val {
			numsVal++
			right--
		}
		if left < right {
			numsVal++
			temp := nums[left]
			nums[left] = nums[right]
			nums[right] = temp
			left++
			right--
		}
	}
	return len(nums) - numsVal
}

// move all element not equal val to begin of nums
func removeElement2(nums []int, val int) int {
	i, j := 0, 0
	for j < len(nums) {
		if nums[j] != val {
			nums[i] = nums[j]
			i++
			j++
		}
		j++
	}
	return i
}
