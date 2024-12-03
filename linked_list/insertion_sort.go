package main

import "fmt"

/*
Этот код сортирует связный список с помощью алгоритма сортировки вставками:
он проходит по каждому элементу списка, вынимает его,
находит подходящее место в уже отсортированной части списка и вставляет туда,
пока весь список не будет отсортирован. В конце обновляя голову списка,
чтобы она указывала на отсортированную часть.
*/

func (ll *LinkedList) InsertionSort() {
	// Проверяем, пуст ли список или состоит ли он только из одного элемента.
	if ll.head == nil || ll.head.next == nil {
		fmt.Println("Список пуст.")
		return
	}

	var sorted *Node // Переменная для отсортированного списка.  Начало отсортированного списка
	// Проходим по всему списку
	current := ll.head
	for current != nil {
		// Берем следующий элемент для вставки
		next := current.next
		// Вставляем текущий элемент в отсортированную часть
		if sorted == nil || sorted.value >= current.value {
			// Если отсортированный список пуст или текущий элемент меньше или равен первому элементу отсортированного списка,
			// добавляем его в начало отсортированного списка
			current.next = sorted
			sorted = current
		} else {
			// Ищем место для вставки текущего элемента
			sortedCurrent := sorted
			for sortedCurrent.next != nil && sortedCurrent.next.value < current.value {
				sortedCurrent = sortedCurrent.next
			}
			// Вставляем элемент между sortedCurrent и sortedCurrent.next
			current.next = sortedCurrent.next
			sortedCurrent.next = current
		}
		// Переходим к следующему элементу в списке
		current = next
	}

	// Переносим отсортированный список обратно в исходный список
	ll.head = sorted
}
