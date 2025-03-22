package main

import (
	linkedList "algs/Data_Structures/6_Linked_List/1_Basics_of_Linked_List/1_Singly_Linked_List"
	"fmt"
)

// удалить последнее вхождение в списке
func deleteLastOccurrence(head *linkedList.Node, target int) *linkedList.Node {
	if head == nil {
		return nil
	}

	var prev *linkedList.Node
	var last *linkedList.Node
	var lastPrev *linkedList.Node
	current := head

	for current != nil {
		if current.Val == target {
			lastPrev = prev
			last = current
		}
		prev = current
		current = current.Next
	}

	if last != nil {
		if lastPrev != nil {
			lastPrev.Next = last.Next
		} else {
			head = head.Next
		}
	}

	return head
}

func main() {
	// Создаем список
	head := &linkedList.Node{Val: 1}
	head = linkedList.InsertAtEnd(head, 2)
	head = linkedList.InsertAtEnd(head, 3)
	head = linkedList.InsertAtEnd(head, 4)
	head = linkedList.InsertAtEnd(head, 5)
	head = linkedList.InsertAtEnd(head, 6)
	head = linkedList.InsertAtEnd(head, 5)

	// Выводим элементы списка
	fmt.Print("Список элементов: ")
	linkedList.TraverseLinkedList(head)

	head = deleteLastOccurrence(head, 6)
	linkedList.TraverseLinkedList(head)
}
