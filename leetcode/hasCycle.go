package main

import "fmt"

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}

	isVisited := map[*ListNode]int{}
	cur := head
	index := 0
	for cur != nil {
		index++
		if _, ok := isVisited[cur]; ok {
			return true
		}
		isVisited[cur] = index
		cur = cur.Next

	}
	return false
}

func hasCycle1(head *ListNode) bool {
	if head == nil {
		return false
	}

	cur := head
	for cur != nil {
		localCurr := cur.Next
		for localCurr != nil {
			if localCurr == cur {
				return true
			}
			localCurr = localCurr.Next
		}
		cur = cur.Next
	}
	return false
}

func testHasCycle() {
	input := [][]int{
		{1, 1},
		{1, 1, 2},
		{1, 1, 2, 3, 3},
	}

	for _, v := range input {
		lInput := fromArrayToListNode(v)
		lInput.print()
		fmt.Println(hasCycle1(lInput))
	}
}
