package main

import (
	"fmt"
)

type Node struct {
	Val  int
	Next *Node
	Prev *Node
}

func TraverseLinkedList(head *Node) {
	if head == nil {
		fmt.Println("Список пуст.")
		return
	}

	current := head
	for current != nil {
		fmt.Print(current.Val, " ")
		current = current.Next
	}

	fmt.Println()
}

func length(head *Node) int {
	res := 0
	if head == nil {
		return res
	}

	current := head
	for current != nil {
		res++
		current = current.Next
	}

	return res
}

func main() {
	head := &Node{Val: 1}
	second := &Node{Val: 2}
	third := &Node{Val: 3}

	head.Next = second
	second.Next = third
	third.Prev = second

	TraverseLinkedList(head)
	fmt.Println(length(head))
}
