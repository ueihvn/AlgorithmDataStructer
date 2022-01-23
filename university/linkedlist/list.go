package linkedlist

import "fmt"

type node struct {
	// storing data with specific data type(basic data type, user define data type)
	info int
	// storing address of the next node
	pNext *node
}

type list struct {
	//storing address of the first node
	pHead *node
	//storing address of the last node
	pTail *node
}

func GetNode(x int) *node {
	n := new(node)
	if n == nil {
		return nil
	}
	n.info = x
	n.pNext = nil
	return n
}

func NewList() *list {
	return new(list)
}

func (l *list) IsEmpty() bool {
	if l.pHead == nil {
		return true
	}
	return false
}

func (l *list) AddHead(p *node) {
	if l.pHead == nil {
		l.pHead = p
		l.pTail = p
	} else {
		p.pNext = l.pHead
		l.pHead = p
	}
}

func (l *list) AddTail(p *node) {
	if l.pHead == nil {
		l.pHead = p
		l.pTail = p
	} else {
		l.pTail.pNext = p
		l.pTail = p
	}
}

func (l *list) InsertAfterQ(p, q *node) {
	if q != nil {
		p.pNext = q.pNext
		q.pNext = p
		if q == l.pTail {
			l.pTail = p
		}
	}
	l.AddHead(p)
}

func (l list) Print() {
	p := l.pHead
	for p != nil {
		fmt.Println(p.info)
		p = p.pNext
	}
}

func (l *list) RemoveHead() bool {
	if l.pHead != nil {
		l.pHead = l.pHead.pNext
		// if list have 1 node -> delete node -> empty
		// update pTail
		if l.pHead == nil {
			l.pTail = nil
		}
		return true
	}
	return false
}

func (l *list) RemoveAfterQ(q *node) bool {
	if q != nil {
		// p is delete node
		p := q.pNext
		// q is not the last node (pTail)
		if p != nil {
			q.pNext = p.pNext
			// delete node is last node
			// update pTail
			if p == l.pTail {
				l.pTail = q
			}
			return true
		}
	}
	return false
}

func (l *list) RemoveX(x int) bool {
	// q is the previous node before p
	q := new(node)
	// p is the node have info = x
	p := l.pHead
	for p != nil && p.info != x {
		q = p
		p = p.pNext
	}

	// can't find node have info = x
	if p == nil {
		return false
	}
	// find a node have info = x

	if q != nil {
		l.RemoveAfterQ(q)
	} else {
		// q = nil -> node have info = x is pHead
		l.RemoveHead()
	}

	return true
}

func (l *list) Search(x int) *node {
	p := l.pHead
	for p != nil && p.info != x {
		p = p.pNext
	}
	return p
}
