package main

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	currL1 := l1
	currL2 := l2
	var res *ListNode
	var curr *ListNode

	for !(currL1 == nil && currL2 == nil) {
		if (currL1 != nil && currL2 != nil && currL1.Val < currL2.Val) || currL2 == nil {
			if res == nil {
				res = &ListNode{Val: currL1.Val}
				curr = res
			} else {
				curr.Next = &ListNode{Val: currL1.Val}
				curr = curr.Next
			}
			currL1 = currL1.Next
		} else {
			if res == nil {
				res = &ListNode{Val: currL2.Val}
				curr = res
			} else {
				curr.Next = &ListNode{Val: currL2.Val}
				curr = curr.Next
			}
			currL2 = currL2.Next
		}
	}
	return res
}

func testMergeTwoLists() {
	input := [][][]int{
		{
			{1, 2, 4},
			{1, 3, 4},
		},
		{
			{},
			{},
		},
		{
			{},
			{0},
		},
	}

	for _, v := range input {
		l1 := fromArrayToListNode(v[0])
		l2 := fromArrayToListNode(v[1])
		l1.print()
		l2.print()
		l3 := mergeTwoLists(l1, l2)
		l3.print()
	}
}
