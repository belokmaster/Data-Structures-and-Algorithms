package main

import (
	linkedList "algs/Data_Structures/6_Linked_List/1_Basics_of_Linked_List/1_Singly_Linked_List"
)

// Даны два отсортированных связанных списка, состоящих из n и m узлов соответственно.
// Задача состоит в том, чтобы объединить оба списка и вернуть начало объединенного списка.

func mergeLinkedList(head1, head2 *linkedList.Node) *linkedList.Node {
	mergedLinkedList := &linkedList.Node{}
	mergedHead := mergedLinkedList

	current1 := head1
	current2 := head2

	for current1 != nil && current2 != nil {
		if current1.Val < current2.Val {
			mergedHead.Next = current1
			current1 = current1.Next
		} else {
			mergedHead.Next = current2
			current2 = current2.Next
		}
		mergedHead = mergedHead.Next
	}

	// Присоединяем оставшийся узел (если есть)
	if current1 != nil {
		mergedHead.Next = current1
	} else {
		mergedHead.Next = current2
	}

	return mergedLinkedList.Next
}

func main() {
	head1 := &linkedList.Node{Val: 1, Next: &linkedList.Node{Val: 3, Next: &linkedList.Node{Val: 5}}}
	head2 := &linkedList.Node{Val: 2, Next: &linkedList.Node{Val: 4, Next: &linkedList.Node{Val: 6}}}

	mergedHead := mergeLinkedList(head1, head2)
	linkedList.TraverseLinkedList(mergedHead)
}
