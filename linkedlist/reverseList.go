package main

func reverseList(l LinkedList) LinkedList {
	if l.head == nil {
		return l
	}
	cur := l.head
	res := LinkedList{}

	for cur != nil {
		temp := res.head
		res.head = cur
		// set cur = cur.net before res.head.next
		// res.head is cur so if res.head change -> cur change
		//  set cur = cur.net before so cur now != res.head, we can change res.head.next = temp
		// don't affect cur
		cur = cur.next
		res.head.next = temp
		res.length++
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
