package main

import (
	linkedList "algs/Data_Structures/6_Linked_List/1_Basics_of_Linked_List/1_Singly_Linked_List"
	"fmt"
)

// Дан заголовок связанного списка. Задача состоит в том, чтобы найти длину цикла в связанном списке.
// Если цикл отсутствует, вернуть 0.
func countNodesLoop(head *linkedList.Node) int {
	if head == nil {
		return 0
	}

	slow := head
	fast := head

	for slow != nil && fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		if slow == fast {
			// Если нашли цикл, подсчитываем количество узлов в цикле
			count := 1
			current := slow.Next
			for current != slow {
				count++
				current = current.Next
			}
			return count
		}
	}

	return 0
}

func main() {
	// Создаем список
	head := &linkedList.Node{Val: 1}
	head = linkedList.InsertAtEnd(head, 2)
	head = linkedList.InsertAtEnd(head, 3)
	head = linkedList.InsertAtEnd(head, 4)
	head = linkedList.InsertAtEnd(head, 5)

	// Создаем цикл в списке: 5 -> 2
	head.Next.Next.Next.Next.Next = head.Next

	fmt.Println(countNodesLoop(head))
}
