package main

import (
	linkedList "algs/Data_Structures/6_Linked_List/1_Basics_of_Linked_List/1_Singly_Linked_List"
	"fmt"
)

// удалить середину связанного списка

// два поинтета. алгоритм флойда
func deleteMiddle(head *linkedList.Node) *linkedList.Node {
	if head == nil {
		return nil
	}

	var prev *linkedList.Node
	slow := head
	fast := head

	for fast != nil && fast.Next != nil {
		prev = slow
		slow = slow.Next
		fast = fast.Next.Next
	}

	prev.Next = slow.Next

	return head
}

func deleteMiddleTwoPass(head *linkedList.Node) *linkedList.Node {
	if head == nil {
		return nil
	}

	sizeLinkedList := 0
	current := head
	for current != nil {
		sizeLinkedList++
		current = current.Next
	}

	mid := sizeLinkedList / 2
	count := 0
	current = head

	var prev *linkedList.Node
	for current != nil {
		if count == mid {
			// Удаление элемента: prev.Next указывает на current.Next
			if prev != nil {
				prev.Next = current.Next
			} else {
				// Если мы удаляем первый элемент
				head = current.Next
			}
			break
		}
		prev = current
		current = current.Next
		count++
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

	head = deleteMiddleTwoPass(head)
	linkedList.TraverseLinkedList(head)
}
