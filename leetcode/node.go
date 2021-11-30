package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) print() {
	cur := l
	for cur != nil {
		fmt.Printf("%d ", cur.Val)
		cur = cur.Next
	}
	fmt.Println()
}

func fromArrayToListNode(nums []int) *ListNode {
	var l *ListNode
	for _, v := range nums {
		if l == nil {
			l = &ListNode{Val: v}
		} else {
			cur := l
			for cur.Next != nil {
				cur = cur.Next
			}
			cur.Next = &ListNode{Val: v}
		}
	}
	return l
}
