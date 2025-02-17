package main

import (
	linkedList "algs/Data_Structures/6_Linked_List/1_Basics_of_Linked_List/1_Singly_Linked_List"
	"fmt"
)

func ReverseLinkedList(head *linkedList.Node) *linkedList.Node {
	var prev *linkedList.Node
	current := head

	for current != nil {
		next := current.Next
		current.Next = prev
		prev = current
		current = next
	}

	return prev
}

func main() {
	// Создаем список
	head := &linkedList.Node{Val: 1}
	head = linkedList.InsertAtEnd(head, 2)
	head = linkedList.InsertAtEnd(head, 3)
	head = linkedList.InsertAtEnd(head, 4)
	head = linkedList.InsertAtEnd(head, 5)
	head = linkedList.InsertAtEnd(head, 5)

	// Выводим элементы списка
	fmt.Print("Список элементов: ")
	linkedList.TraverseLinkedList(head)

	head = ReverseLinkedList(head)
	linkedList.TraverseLinkedList(head)
}
