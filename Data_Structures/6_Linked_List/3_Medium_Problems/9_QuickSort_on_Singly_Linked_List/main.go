package main

import (
	linkedList "algs/Data_Structures/6_Linked_List/1_Basics_of_Linked_List/1_Singly_Linked_List"
)

// Алгоритм быстрой сортировки работает путем выбора опорного элемента из массива.
// Затем массив разбивается таким образом, что все элементы, меньшие опорного, размещаются слева,
// а те, что больше опорного, размещаются справа. После завершения разбиения алгоритм рекурсивно применяет
// тот же процесс к левому и правому подмассивам.

// находит последний узел списка
func getTail(head *linkedList.Node) *linkedList.Node {
	if head == nil {
		return nil
	}

	for head.Next != nil {
		head = head.Next
	}

	return head
}

// Эта функция выполняет разбиение списка относительно опорного элемента
func partition(head, tail *linkedList.Node) *linkedList.Node {
	pivot := head
	pre := head
	curr := head

	for curr != nil && curr != tail.Next {
		if curr.Val < pivot.Val {
			pre = pre.Next
			pre.Val, curr.Val = curr.Val, pre.Val
		}
		curr = curr.Next
	}

	pivot.Val, pre.Val = pre.Val, pivot.Val
	return pre
}

// рекурсивно сортирует список
func quickSortHelper(head, tail *linkedList.Node) {
	if head == nil || head == tail {
		return
	}

	pivot := partition(head, tail)
	quickSortHelper(head, pivot)
	quickSortHelper(pivot.Next, tail)
}

func quickSort(head *linkedList.Node) *linkedList.Node {
	tail := getTail(head)
	quickSortHelper(head, tail)
	return head
}

func main() {
	head := &linkedList.Node{Val: 30, Next: &linkedList.Node{Val: 3, Next: &linkedList.Node{Val: 4, Next: &linkedList.Node{Val: 20, Next: &linkedList.Node{Val: 5}}}}}
	head = quickSort(head)
	linkedList.TraverseLinkedList(head)
}
