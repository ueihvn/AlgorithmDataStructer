package main

func mergeTwoLists(l1 LinkedList, l2 LinkedList) LinkedList {
	templ1 := l1.Head
	templ2 := l2.Head
	var res LinkedList
	var cur *Node

	for !(templ1 == nil && templ2 == nil) {
		if (templ1 != nil && templ2 != nil && templ1.Val < templ2.Val) || (templ2 == nil) {
			if res.Head == nil {
				res.Head = &Node{Val: templ1.Val}
				cur = res.Head
				res.Length++
			} else {
				cur.Next = &Node{Val: templ1.Val}
				cur = cur.Next
				res.Length++
			}
			templ1 = templ1.Next
		} else {
			if res.Head == nil {
				res.Head = &Node{Val: templ2.Val}
				cur = res.Head
				res.Length++
			} else {
				cur.Next = &Node{Val: templ2.Val}
				cur = cur.Next
				res.Length++
			}
			templ2 = templ2.Next
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
