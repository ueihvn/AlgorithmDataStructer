package linkedlist

func SelectionSort(l *list) {
	var p, q, min *node
	p = l.pHead
	for p != nil {
		min = p
		q = p.pNext
		for q != nil {
			if q.info <= min.info {
				min = q
			}
			q = q.pNext
		}
		p.info, min.info = min.info, p.info
		p = p.pNext
	}
}
