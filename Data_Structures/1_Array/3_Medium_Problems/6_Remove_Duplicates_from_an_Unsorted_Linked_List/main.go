package main

import (
	linkedList "algs/Data_Structures/6_Linked_List/1_Basics_of_Linked_List/1_Singly_Linked_List"
	"fmt"
)

// Задача состоит в том, чтобы удалить повторяющиеся узлы из несортированного связанного списка,
// содержащего n узлов, сохранив исходный порядок.
func removeDuplicates(head *linkedList.Node) *linkedList.Node {
	current := head
	var prev *linkedList.Node
	elements := make(map[int]bool)

	for current != nil {
		_, exists := elements[current.Val]
		if exists {
			prev.Next = current.Next
		} else {
			elements[current.Val] = true
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
	head = linkedList.InsertAtEnd(head, 3)
	head = linkedList.InsertAtEnd(head, 4)
	head = linkedList.InsertAtEnd(head, 2)

	// Выводим элементы списка
	fmt.Print("Список элементов: ")
	linkedList.TraverseLinkedList(head)

	head = removeDuplicates(head)
	linkedList.TraverseLinkedList(head)
}
