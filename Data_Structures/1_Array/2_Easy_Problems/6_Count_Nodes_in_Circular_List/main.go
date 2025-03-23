package main

import (
	linkedList "algs/Data_Structures/6_Linked_List/1_Basics_of_Linked_List/1_Singly_Linked_List"
	"fmt"
)

func countNodesCircular(head *linkedList.Node) int {
	if head == nil {
		return 0
	}

	count := 1 // Учтем первый узел
	current := head.Next

	for current != head { // Пока не вернулись к начальному узлу
		count++
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
	head = linkedList.InsertAtEnd(head, 6)

	head = linkedList.MakeCircularLinkedList(head)
	count := countNodesCircular(head)
	fmt.Println(count)
}
