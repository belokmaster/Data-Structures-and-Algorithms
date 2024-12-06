package main

// split делит список на две половины
func (ll *LinkedList) split(head *Node) (*Node, *Node) {
	// Проверяем, пуст ли список или состоит ли он только из одного элемента.
	if head == nil || head.next == nil {
		return head, nil
	}

	// Инициализация указателей
	slow := head
	fast := head
	var prev *Node

	// Ищем середину списка
	for fast != nil && fast.next != nil {
		prev = slow
		slow = slow.next
		fast = fast.next.next
	}

	// Разделяем список на две части
	prev.next = nil
	return head, slow
}

// merge объединяет два отсортированных списка
func merge(left, right *Node) *Node {
	// Базовые случаи: один из списков пуст
	if left == nil {
		return right
	}
	if right == nil {
		return left
	}

	// Инициализация результата
	var result *Node

	// Сравнение первых элементов и рекурсивный вызов для оставшихся элементов
	if left.value < right.value {
		result = left
		result.next = merge(left.next, right)
	} else {
		result = right
		result.next = merge(left, right.next)
	}
	return result
}

// mergeSort выполняет сортировку слиянием для списка
func (ll *LinkedList) mergeSort(head *Node) *Node {
	if head == nil || head.next == nil {
		return head
	}

	// Разделяем список на две части
	left, right := ll.split(head)

	// Рекурсивно сортируем каждую половину
	sortedLeft := ll.mergeSort(left)
	sortedRight := ll.mergeSort(right)

	// Объединяем отсортированные половины
	return merge(sortedLeft, sortedRight)
}

// MergeSort выполняет сортировку списка с использованием сортировки слиянием
func (ll *LinkedList) MergeSort() {
	ll.head = ll.mergeSort(ll.head)
}
