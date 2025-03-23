package main

import (
	linkedList "algs/Data_Structures/6_Linked_List/1_Basics_of_Linked_List/1_Singly_Linked_List"
)

// Дан связный список и значение x. Разбейте его на части так, чтобы сначала шли все узлы со
// значением меньше x, затем все узлы со значением, равным x, и, наконец, узлы со значением больше x.
// Исходный относительный порядок узлов в каждой из трёх частей должен быть сохранён.
func partition(head *linkedList.Node, x int) *linkedList.Node {
	if head == nil {
		return nil
	}

	lessHead := &linkedList.Node{}
	equalHead := &linkedList.Node{}
	greaterHead := &linkedList.Node{}

	less := lessHead
	equal := equalHead
	greater := greaterHead

	current := head
	for current != nil {
		if current.Val < x {
			less.Next = current
			less = less.Next
		} else if current.Val == x {
			equal.Next = current
			equal = equal.Next
		} else {
			greater.Next = current
			greater = greater.Next
		}
		current = current.Next
	}

	greater.Next = nil            // Завершаем список "больше x"
	equal.Next = greaterHead.Next // Соединяем равные и большие
	less.Next = equalHead.Next    // Соединяем меньше и равные

	// Новый заголовок перераспределенного списка
	return lessHead.Next
}

func main() {
	head := &linkedList.Node{Val: 3}
	head = linkedList.InsertAtEnd(head, 5)
	head = linkedList.InsertAtEnd(head, 8)
	head = linkedList.InsertAtEnd(head, 5)
	head = linkedList.InsertAtEnd(head, 10)
	head = linkedList.InsertAtEnd(head, 2)
	head = linkedList.InsertAtEnd(head, 1)

	head = partition(head, 5)
	linkedList.TraverseLinkedList(head)
}
