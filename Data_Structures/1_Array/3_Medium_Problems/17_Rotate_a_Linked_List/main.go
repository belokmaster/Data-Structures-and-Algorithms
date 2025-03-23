package main

import (
	linkedList "algs/Data_Structures/6_Linked_List/1_Basics_of_Linked_List/1_Singly_Linked_List"
)

// Дан односвязный список и целое число k . Задача состоит в том, чтобы повернуть связанный список влево на k позиций .

// [Наивный подход] Смещение головного узла в конец k раз – O(n * k) времени и O(1) пространства
// Чтобы повернуть связанный список влево на k позиций,  многократно перемещаем головной узел в конец списка k раз.

func rotate(head *linkedList.Node, k int) *linkedList.Node {
	if k == 0 || head == nil {
		return head
	}

	// Поворачиваем список на k узлов
	for i := 0; i < k; i++ {
		current := head
		for current.Next != nil {
			current = current.Next
		}

		// Перемещаем первый узел в конец
		current.Next = head
		current = current.Next
		head = head.Next
		current.Next = nil
	}
	return head
}

func main() {
	head := &linkedList.Node{Val: 1}
	head = linkedList.InsertAtEnd(head, 2)
	head = linkedList.InsertAtEnd(head, 3)
	head = linkedList.InsertAtEnd(head, 4)
	head = linkedList.InsertAtEnd(head, 5)
	head = linkedList.InsertAtEnd(head, 5)

	head = rotate(head, 6)
	linkedList.TraverseLinkedList(head)
}
