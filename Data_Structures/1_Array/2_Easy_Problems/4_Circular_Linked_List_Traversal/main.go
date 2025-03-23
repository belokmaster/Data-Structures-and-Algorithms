package main

import (
	linkedList "algs/Data_Structures/6_Linked_List/1_Basics_of_Linked_List/1_Singly_Linked_List"
	"fmt"
)

// Функция для вывода элементов циклического списка
func printList(head *linkedList.Node) {
	// Если список пустой
	if head == nil {
		return
	}

	current := head
	for {
		fmt.Print(current.Val, " ")
		current = current.Next
		if current == head { // если снова вернулись к голове
			break
		}
	}

	fmt.Println() // для новой строки после вывода
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
