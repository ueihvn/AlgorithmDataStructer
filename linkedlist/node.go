package main

type Node struct {
	val  int
	next *Node
}

func NewNode(n int) *Node {
	return &Node{
		val: n,
	}
}
