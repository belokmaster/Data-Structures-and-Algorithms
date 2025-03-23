package main

import (
	linkedList "algs/Data_Structures/6_Linked_List/1_Basics_of_Linked_List/1_Singly_Linked_List"
	"fmt"
)

func makeCircular(head *linkedList.Node) *linkedList.Node {
	if head == nil {
		return nil
	}

	current := head
	for current.Next != nil {
		current = current.Next
	}

	current.Next = head
	return head
}

func main() {
	head := &linkedList.Node{Val: 1}
	for i := 0; i < 5; i++ {
		head = linkedList.InsertAtEnd(head, i+1)
	}

	head = makeCircular(head)
	linkedList.TraverseCircularLinkedList(head)

	current := head
	fmt.Println(current)
	for current != nil {
		current = current.Next
		if current == head {
			break
		}
	}

	fmt.Println(current)
}
