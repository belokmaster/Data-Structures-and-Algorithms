package main

import (
	linkedList "algs/Data_Structures/6_Linked_List/1_Basics_of_Linked_List/1_Singly_Linked_List"
	"fmt"
)

// Дан односвязанный список, задача состоит в том,
// чтобы поменять местами элементы последовательности попарно.
func pairwiseSwap(head *linkedList.Node) *linkedList.Node {
	current := head
	for current != nil && current.Next != nil {
		current.Val, current.Next.Val = current.Next.Val, current.Val
		current = current.Next.Next // на два
	}
	return head
}

func main() {
	// Создаем список
	head := &linkedList.Node{Val: 1}
	head = linkedList.InsertAtEnd(head, 2)
	head = linkedList.InsertAtEnd(head, 3)
	head = linkedList.InsertAtEnd(head, 4)
	head = linkedList.InsertAtEnd(head, 5)
	head = linkedList.InsertAtEnd(head, 6)

	fmt.Println(linkedList.IsCurcular(head))
	// Выводим элементы списка
	linkedList.TraverseLinkedList(head)

	head = pairwiseSwap(head)
	linkedList.TraverseLinkedList(head)
}
