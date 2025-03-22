package main

import (
	linkedList "algs/Data_Structures/6_Linked_List/1_Basics_of_Linked_List/1_Singly_Linked_List"
	"fmt"
)

func deleteFirstNode(head *linkedList.Node) *linkedList.Node {
	if head == nil {
		return nil
	}

	if head.Next == head {
		return nil
	}

	newHead := head.Next
	current := head

	for current.Next != head {
		current = current.Next
	}

	current.Next = newHead
	return newHead
}

func deleteLastNode(head *linkedList.Node) *linkedList.Node {
	if head == nil || head.Next == head {
		return nil
	}

	current := head
	prev := head
	for current.Next != head {
		current = current.Next
		if current.Next != head {
			prev = current
		} else {
			prev.Next = head
			break
		}
	}

	return head
}

func deleteIndexNode(head *linkedList.Node, index int) *linkedList.Node {
	if head == nil {
		return nil
	}

	if index == 0 {
		if head.Next == head {
			return nil
		}

		newHead := head.Next
		current := head

		for current.Next != head {
			current = current.Next
		}

		current.Next = newHead
		return newHead
	}

	count := 0
	current := head
	prev := head

	for current.Next != nil {
		count++
		if count > index {
			fmt.Println("Индекс выходит за границы списка.")
			return head
		}

		prev = current
		current = current.Next

		if count == index {
			prev.Next = current.Next
			return head
		}

		if current == head {
			fmt.Println("Индекс выходит за границы списка.")
			return head
		}
	}

	return head
}

func main() {
	// Создаем список
	head := &linkedList.Node{Val: 1}
	for i := 2; i < 7; i++ {
		head = linkedList.InsertAtEnd(head, i)
	}

	linkedList.TraverseLinkedList(head)

	head = linkedList.MakeCircularLinkedList(head)
	linkedList.TraverseCircularLinkedList(head)

	fmt.Println()
	fmt.Println("After delete: ")
	head = deleteLastNode(head)
	head = deleteFirstNode(head)
	linkedList.TraverseCircularLinkedList(head)

	fmt.Println()
	index := 10
	fmt.Printf("After delete at position %v:\n", index)
	head = deleteIndexNode(head, index)
	linkedList.TraverseCircularLinkedList(head)
}
