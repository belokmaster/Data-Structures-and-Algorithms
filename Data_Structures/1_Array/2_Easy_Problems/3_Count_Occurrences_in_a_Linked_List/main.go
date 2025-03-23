package main

import (
	linkedList "algs/Data_Structures/6_Linked_List/1_Basics_of_Linked_List/1_Singly_Linked_List"
	"fmt"
)

// алгоритм флойда, он же черепаха и заяц
func Count(head *linkedList.Node, target int) int {
	count := 0
	current := head

	for current != nil {
		if current.Val == target {
			count++
		}
		current = current.Next
	}

	return count
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

	// Находим и выводим средний элемент
	fmt.Println(Count(head, 3))
}
