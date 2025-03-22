package main

import (
	linkedList "algs/Data_Structures/6_Linked_List/1_Basics_of_Linked_List/1_Singly_Linked_List"
	"fmt"
)

// Дан связанный список и число N.
// Найдите значение N-го узла от конца связанного списка.
// Если N-го узла от конца нет, выведите -1.

func findNthFromLast(head *linkedList.Node, index int) int {
	if head == nil {
		return -1
	}

	sizeLinkedList := 0
	current := head

	for current != nil {
		sizeLinkedList++
		current = current.Next
	}

	if index > sizeLinkedList || index < 0 {
		fmt.Println("Индекс выходит за границы списка.")
		return -1
	}

	dif := sizeLinkedList - index + 1
	size := 0
	temp := head
	for temp != nil {
		size++
		if size == dif {
			return temp.Val
		}
		temp = temp.Next
	}

	return -1
}

// два указателя. один двигаем на index. когда он доходит до конца, второй указатель уже на нужном числе.
func findNthFromLastTwoPointer(head *linkedList.Node, index int) int {
	if head == nil {
		return -1
	}

	mainPtr := head
	refPtr := head

	// Переместить refPtr на N-й узел от начала
	for i := 0; i < index; i++ {
		if refPtr == nil {
			return -1
		}
		refPtr = refPtr.Next
	}

	for refPtr != nil {
		refPtr = refPtr.Next
		mainPtr = mainPtr.Next
	}

	return mainPtr.Val
}

func main() {
	// Создаем список
	head := &linkedList.Node{Val: 1}
	head = linkedList.InsertAtEnd(head, 2)
	head = linkedList.InsertAtEnd(head, 3)
	head = linkedList.InsertAtEnd(head, 4)
	head = linkedList.InsertAtEnd(head, 5)
	head = linkedList.InsertAtEnd(head, 6)

	// Выводим элементы списка
	fmt.Print("Список элементов: ")
	linkedList.TraverseLinkedList(head)

	fmt.Println()
	ans := findNthFromLast(head, 3)
	ans2 := findNthFromLastTwoPointer(head, 3)
	fmt.Println(ans, ans2)
}
