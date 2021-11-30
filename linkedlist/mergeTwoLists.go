package main

func mergeTwoLists(l1 LinkedList, l2 LinkedList) LinkedList {
	templ1 := l1.head
	templ2 := l2.head
	var res LinkedList
	var cur *Node

	for !(templ1 == nil && templ2 == nil) {
		if (templ1 != nil && templ2 != nil && templ1.val < templ2.val) || (templ2 == nil) {
			if res.head == nil {
				res.head = &Node{val: templ1.val}
				cur = res.head
				res.length++
			} else {
				cur.next = &Node{val: templ1.val}
				cur = cur.next
				res.length++
			}
			templ1 = templ1.next
		} else {
			if res.head == nil {
				res.head = &Node{val: templ2.val}
				cur = res.head
				res.length++
			} else {
				cur.next = &Node{val: templ2.val}
				cur = cur.next
				res.length++
			}
			templ2 = templ2.next
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
		l1 := fromArrayToLinkedList(v[0])
		l2 := fromArrayToLinkedList(v[1])
		l1.print()
		l2.print()
		l3 := mergeTwoLists(l1, l2)
		l3.print()
	}
}
