package main

func deleteDuplicates(l *LinkedList) *LinkedList {
	if !(l.Head != nil && l.Head.Next != nil) {
		return l
	}

	lastNode := l.Head
	currNode := l.Head.Next

	for currNode != nil {
		if lastNode.Val == currNode.Val {
			// duplicate so set Next = nil
			lastNode.Next = nil
			currNode = currNode.Next
			l.Length--
		} else {
			lastNode.Next = currNode
			lastNode = currNode
			currNode = currNode.Next
			// last node of linked list -> set Next = nil
			lastNode.Next = nil
		}
	}
	return l
}

func testDeleteDuplicates() {
	input := [][]int{
		{1, 1, 2},
		{1, 1, 2, 3, 3},
	}

	for _, v := range input {
		lInput := fromArrayToLinkedList(v)
		lInput.print()
		deleteDuplicates(&lInput).print()
	}
}
