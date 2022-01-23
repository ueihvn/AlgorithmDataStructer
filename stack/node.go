package main

type Node struct {
	Val  int
	Next *Node
}

func NewNode(n int) *Node {
	return &Node{
		Val: n,
	}
}
