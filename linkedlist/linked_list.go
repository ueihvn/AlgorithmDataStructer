package main

import "fmt"

type LinkedList struct {
	head   *Node
	length int
}

func (l *LinkedList) prepend(n int) {
	newNode := NewNode(n)
	newNode.next = l.head
	l.head = newNode
	l.length++
}

func (l *LinkedList) append(n int) {
	newNode := NewNode(n)
	if l.length == 0 {
		l.head = newNode
		l.length++
		return
	}
	temp := l.head
	for temp.next != nil {
		temp = temp.next
	}
	temp.next = newNode
	l.length++
}

func (l *LinkedList) deleteFirst() {
	if l.length < 1 {
		fmt.Println("empty linked list cann't delete frist")
		return
	}
	l.head = l.head.next
	l.length--
}

func (l *LinkedList) deleteEnd() {
	if l.length < 1 {
		fmt.Println("empty linked list cann't delete end")
		return
	}

	// only 1 node is head -> head = nil
	if l.length == 1 {
		l.head = nil
		l.length--
		return
	}

	// length > 1
	temp := l.head
	for temp.next.next != nil {
		temp = temp.next
	}
	temp.next = nil
	l.length--
}

func (l LinkedList) print() {
	temp := l.head
	for l.length > 0 {
		fmt.Printf("%d ", temp.val)
		temp = temp.next
		l.length--
	}
	fmt.Println()
}
