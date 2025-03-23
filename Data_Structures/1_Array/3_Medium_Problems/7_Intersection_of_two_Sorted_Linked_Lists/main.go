package main

import (
	linkedList "algs/Data_Structures/6_Linked_List/1_Basics_of_Linked_List/1_Singly_Linked_List"
	"fmt"
)

// найти пересечение двух связанных списков
func findIntersection(head1, head2 *linkedList.Node) []int {
	ans := []int{}
	current := head1
	digits1 := make(map[int]bool)

	for current != nil {
		digits1[current.Val] = true
		current = current.Next
	}

	current = head2
	digits2 := make(map[int]bool)
	for current != nil {
		if digits1[current.Val] && !digits2[current.Val] {
			ans = append(ans, current.Val)
			digits2[current.Val] = true
		}
		current = current.Next
	}

	return ans
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

	// Создаем список
	head2 := &linkedList.Node{Val: 1}
	head2 = linkedList.InsertAtEnd(head2, 2)
	head2 = linkedList.InsertAtEnd(head2, 2)
	head2 = linkedList.InsertAtEnd(head2, 2)
	head2 = linkedList.InsertAtEnd(head2, 3)
	head2 = linkedList.InsertAtEnd(head2, 3)
	head2 = linkedList.InsertAtEnd(head2, 8)
	head2 = linkedList.InsertAtEnd(head2, 4)
	head2 = linkedList.InsertAtEnd(head2, 3)
	head2 = linkedList.InsertAtEnd(head2, 4)
	head2 = linkedList.InsertAtEnd(head2, 2)

	ans := findIntersection(head, head2)
	fmt.Println(ans)
}
