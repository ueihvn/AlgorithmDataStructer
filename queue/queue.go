package main

import (
	"fmt"
)

// FIFO
type Queue struct {
	Head   *Node
	Length int
}

// Add node to begin of linked list
func (s *Queue) Enqueue(_val int) {
	if s == nil {
		fmt.Println("Queue is nil, declare new queue")
		return
	}

	newNode := NewNode(_val)
	if s.Head == nil {
		s.Head = newNode
		s.Length++
		return
	}

	newNode.Next = s.Head
	s.Head = newNode
	s.Length++

}

// Remove the node at end of linked list
func (s *Queue) Dequeue() int {
	if s == nil {
		fmt.Println("Queue is nil, declare new queue")
		return 0
	}

	if s.Head == nil {
		fmt.Println("Queue is empty")
		return 0
	}

	cur := s.Head
	for cur.Next != nil {
		cur = cur.Next
	}

	result := cur.Val
	s.Length--
	cur = nil

	return result
}

func (s *Queue) IsEmpty() bool {
	if s.Head == nil && s.Length == 0 {
		return true
	} else {
		return false
	}
}
