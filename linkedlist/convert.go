package main

func fromArrayToLinkedList(nums []int) LinkedList {
	l := LinkedList{}
	for _, v := range nums {
		l.append(v)
	}
	return l
}
