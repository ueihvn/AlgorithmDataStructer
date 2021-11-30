package main

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	var res *ListNode
	cur := head

	for cur != nil {
		temp := res
		res = cur
		cur = cur.Next
		// cur now is different with res
		// res is old cur
		// cur is cur.Next
		// res, cur is pointer variables
		res.Next = temp
	}
	return res
}

func testReverseLists() {
	input := [][]int{
		{1, 2, 4},
		// {1, 3, 4},
		// {},
		// {0},
	}

	for _, v := range input {
		l := fromArrayToListNode(v)
		l.print()
		reverseList(l).print()
	}
}
