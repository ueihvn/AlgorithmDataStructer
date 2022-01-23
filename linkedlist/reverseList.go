package main

func reverseList(l LinkedList) LinkedList {
	if l.Head == nil {
		return l
	}
	cur := l.Head
	res := LinkedList{}

	for cur != nil {
		temp := res.Head
		res.Head = cur
		// set cur = cur.net before res.Head.Next
		// res.Head is cur so if res.Head change -> cur change
		//  set cur = cur.net before so cur now != res.Head, we can change res.Head.Next = temp
		// don't affect cur
		cur = cur.Next
		res.Head.Next = temp
		res.Length++
	}

	return res
}

func testReverseList() {
	input := [][]int{
		{1, 2, 4},
		{1, 3, 4},
		{},
		{0},
	}

	for _, v := range input {
		l := fromArrayToLinkedList(v)
		l.print()
		res := reverseList(l)
		res.print()
	}
}
