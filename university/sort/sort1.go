package main

// travel through all pair of number(a[i], a[j]) in array and swap if wrong order
// T(n) = O(n^2)
// best case don't have to swap
// worst case have to swap (n-1)/2
func InterchangeSort(num []int, n int) {
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			if num[i] > num[j] {
				num[i], num[j] = num[j], num[i]
			}
		}
	}
}

// find smallest from i to n
// swap that num to the ith
// reapeat with i+1 to n
// T(n) = O(n^2)
// best case don't have to swap
// worst case have to swap (n-1)/2
func SelectionSort(num []int, n int) {
	min := 0
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			if num[j] <= num[min] {
				min = j
			}
		}
		num[min], num[i] = num[i], num[min]
	}
}

// compare each pair of num[i] num[i+1]
// swap the largest num to the end
// T(n) = O(n^2)
// best case don't have to swap
// worst case have to swap (n-1)/2
func BubbleSort(num []int, n int) {
	for i := 0; i < n-i; i++ {
		for j := 0; j < n-i-1; j++ {
			if num[j] > num[j+1] {
				num[j], num[j+1] = num[j+1], num[j]
			}
		}
	}
}
