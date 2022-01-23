package main

func deleteDuplicates(head *ListNode) *ListNode {
	if !(head != nil && head.Next != nil) {
		return head
	}

	lastNode := head
	currNode := head.Next
	for currNode != nil {
		if lastNode.Val == currNode.Val {
			// duplicate so set next = nil
			lastNode.Next = nil
			currNode = currNode.Next
		} else {
			lastNode.Next = currNode
			lastNode = currNode
			currNode = currNode.Next
			// last node of linked list -> set next = nil
			lastNode.Next = nil
		}
	}
	return head
}

func testDeleteDuplicates() {
	input := [][]int{
		{1, 1},
		{1, 1, 2},
		{1, 1, 2, 3, 3},
	}

	for _, v := range input {
		lInput := fromArrayToListNode(v)
		lInput.print()
		deleteDuplicates(lInput).print()
	}
}
