package main

import (
	linkedList "algs/Data_Structures/6_Linked_List/1_Basics_of_Linked_List/1_Singly_Linked_List"
	"fmt"
)

// Дан отсортированный односвязный список. Задача состоит в том, чтобы удалить дубликаты

func removeDuplicates(head *linkedList.Node) *linkedList.Node {
	if head == nil {
		return nil
	}

	current := head
	for current.Next != nil {
		if current.Val == current.Next.Val {
			current.Next = current.Next.Next
		} else {
			current = current.Next
		}
	}

	return head
}

func removeDuplicatesTwoPointers(head *linkedList.Node) *linkedList.Node {
	if head == nil {
		return nil
	}

	prev := head
	current := head

	for current != nil {
		if current.Val == prev.Val {
			prev.Next = current.Next
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
	head = linkedList.InsertAtEnd(head, 2)
	head = linkedList.InsertAtEnd(head, 2)
	head = linkedList.InsertAtEnd(head, 3)
	head = linkedList.InsertAtEnd(head, 3)
	head = linkedList.InsertAtEnd(head, 4)

	// Выводим элементы списка
	fmt.Print("Список элементов: ")
	linkedList.TraverseLinkedList(head)

	head = removeDuplicatesTwoPointers(head)
	linkedList.TraverseLinkedList(head)
}
