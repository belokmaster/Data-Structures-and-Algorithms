package main

import (
	linkedList "algs/Data_Structures/6_Linked_List/1_Basics_of_Linked_List/1_Singly_Linked_List"
	"fmt"
)

// При наличии односвязного списка задача состоит в том, чтобы удалить каждый k-й узел связанного списка.
func deleteK(head *linkedList.Node, k int) *linkedList.Node {
	if head == nil || k < 0 {
		return nil
	}

	current := head
	var prev *linkedList.Node

	count := 0

	for current != nil {
		count++

		if count%k == 0 {
			if prev != nil {
				prev.Next = current.Next
			} else {
				head = current.Next
			}
		} else {
			prev = current
		}
		current = current.Next
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
	head = linkedList.InsertAtEnd(head, 5)

	// Выводим элементы списка
	fmt.Print("Список элементов: ")
	linkedList.TraverseLinkedList(head)

	k := 2
	head = deleteK(head, k)
	linkedList.TraverseLinkedList(head)
}
