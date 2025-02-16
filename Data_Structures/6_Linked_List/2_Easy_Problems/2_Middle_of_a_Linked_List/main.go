package main

import (
	linkedList "algs/Data_Structures/6_Linked_List/1_Basics_of_Linked_List/1_Singly_Linked_List"
	"fmt"
)

// алгоритм флойда, он же черепаха и заяц
func FindMiddle(head *linkedList.Node) *linkedList.Node {
	if head == nil {
		return nil
	}

	slow := head
	fast := head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow
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
	middle := FindMiddle(head)
	if middle != nil {
		fmt.Println("Средний элемент:", middle.Val)
	} else {
		fmt.Println("Список пуст.")
	}
}
