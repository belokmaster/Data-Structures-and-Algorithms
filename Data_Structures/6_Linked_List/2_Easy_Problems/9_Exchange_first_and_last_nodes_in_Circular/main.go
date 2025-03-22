package main

import (
	linkedList "algs/Data_Structures/6_Linked_List/1_Basics_of_Linked_List/1_Singly_Linked_List"
)

func exchangeNodes(head *linkedList.Node) *linkedList.Node {
	if head == nil {
		return nil
	}

	if head == head.Next {
		return head
	}

	current := head
	var prev *linkedList.Node = nil

	// Найти последний узел
	for current.Next != head {
		prev = current
		current = current.Next
	}

	if prev != nil {
		prev.Next = head
	}

	current.Next = head.Next
	head.Next = current
	head = current

	return head
}

func main() {
	head := &linkedList.Node{Val: 1}
	for i := 2; i < 5; i++ {
		head = linkedList.InsertAtEnd(head, i)
	}
	head = linkedList.MakeCircularLinkedList(head)
	linkedList.TraverseCircularLinkedList(head)

	head = exchangeNodes(head)
	linkedList.TraverseCircularLinkedList(head)
}
