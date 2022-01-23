package main

import (
	"fmt"
)

// FILO
type Stack struct {
	Head   *Node
	Length int
}

// Add new node to the begin of linked list
func (s *Stack) Push(_val int) {
	if s == nil {
		fmt.Println("Stack is nil, declare new stack")
		return
	}

	newNode := NewNode(_val)

	if s.Head == nil {
		s.Head = newNode
		s.Length++
		return
	}

	temp := s.Head
	s.Head = newNode
	s.Head.Next = temp
	s.Length++
}

// Delete the node at the end of linked list
func (s *Stack) Pop() int {
	if s == nil {
		fmt.Println("Stack is nil, declare new stack")
		return 0
	}

	if s.Head == nil {
		fmt.Println("Stack is empty")
		return 0
	}

	temp := s.Head.Next
	res := s.Head.Val
	s.Head = temp
	s.Length--

	return res
}

func (s *Stack) IsEmpty() bool {
	if s.Head == nil && s.Length == 0 {
		return true
	} else {
		return false
	}
}
