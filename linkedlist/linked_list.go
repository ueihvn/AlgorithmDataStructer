package main

import "fmt"

type LinkedList struct {
	Head   *Node
	Length int
}

func (l *LinkedList) prepend(n int) {
	newNode := NewNode(n)
	newNode.Next = l.Head
	l.Head = newNode
	l.Length++
}

func (l *LinkedList) append(n int) {
	newNode := NewNode(n)
	if l.Length == 0 {
		l.Head = newNode
		l.Length++
		return
	}
	temp := l.Head
	for temp.Next != nil {
		temp = temp.Next
	}
	temp.Next = newNode
	l.Length++
}

func (l *LinkedList) deleteFirst() {
	if l.Length < 1 {
		fmt.Println("empty linked list can't delete frist")
		return
	}
	l.Head = l.Head.Next
	l.Length--
}

func (l *LinkedList) deleteEnd() {
	if l.Length < 1 {
		fmt.Println("empty linked list can't delete end")
		return
	}

	// only 1 node is head -> head = nil
	if l.Length == 1 {
		l.Head = nil
		l.Length--
		return
	}

	// length > 1
	temp := l.Head
	for temp.Next.Next != nil {
		temp = temp.Next
	}
	temp.Next = nil
	l.Length--
}

func (l LinkedList) print() {
	temp := l.Head
	for l.Length > 0 {
		fmt.Printf("%d ", temp.Val)
		temp = temp.Next
		l.Length--
	}
	fmt.Println()
}
