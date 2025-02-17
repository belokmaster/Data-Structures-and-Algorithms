package main

import (
	linkedList "algs/Data_Structures/6_Linked_List/1_Basics_of_Linked_List/1_Singly_Linked_List"
	"fmt"
)

// Учитывая заголовок односвязного списка, задача состоит в том, чтобы найти, является ли данный
// связанный список циклическим или нет. Связанный список называется циклическим,
// если его последний узел указывает обратно на его первый узел.
func printList(head *linkedList.Node) bool {
	slow := head
	fast := head

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			return true
		}
	}

	return false
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
	printList(head)
}
