package main

// call itself 2 time -> binary recursive
// use recursive to sort 2 sub array have index [right, mid] and [mid+1,right]
// are seperated by array which have index from left to right
func mergeSort(nums []int, left, right int) {
	// range from left to right have more than 1 element
	if left < right {
		// middle of array
		mid := (left + right) / 2
		// sort left
		mergeSort(nums, left, mid)
		// sort right
		mergeSort(nums, mid+1, right)

		// 2 function above done -> left and right is sorted
		// merge two sorted array into 1 sorted array
		merge(nums, left, mid, right)
	}
}

func merge(nums []int, left, mid, right int) {
	i := left
	j := mid + 1
	temp := make([]int, right-left+1)
	// track index of slice temp
	m := 0
	for !(i > mid && j > right) {
		if (nums[i] < nums[j] && i <= mid && j <= right) || (j > right) {
			temp[m] = nums[i]
			i++
			m++
		} else {
			temp[m] = nums[j]
			j++
			m++
		}
	}

	for i := 0; i < m; i++ {
		nums[left+i] = temp[i]
	}
}

func merge2(nums []int, left, mid, right int) {
	// create slice have mid-left + 1 element
	// store all element from left...mid
	// cap phat vung nho co kich thuoc gap mid -left + 1 int
	b := make([]int, mid-left+1)
	k := 0
	for i := left; i <= mid; i++ {
		b[k] = nums[i]
		k++
	}

	// create slice have right - (mid +1) + 1 element
	// store all element from mid+1...right
	c := make([]int, right-mid)
	l := 0
	for i := mid + 1; i <= right; i++ {
		c[l] = nums[i]
		l++
	}

	// create slice have right - (mid +1) + 1 element
	// store all element from left...right
	temp := make([]int, right-left+1)
	// track index of slice temp
	m := 0
	mix(b, c, temp, k, l, &m)

	for i := 0; i < m; i++ {
		nums[left+i] = temp[i]
	}
}

func mix(a, b, c []int, lena, lenb int, lenc *int) {
	var i, j int
	*lenc = 0
	for !(i >= lena && j >= lenb) {
		if i < lena && j < lenb && a[i] < b[j] || j >= lenb {
			c[*lenc] = a[i]
			*lenc++
			i++
		} else {
			c[*lenc] = b[j]
			*lenc++
			j++
		}
	}
}
