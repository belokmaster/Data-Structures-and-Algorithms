package main

import "fmt"

func (ll *LinkedList) RemoveDuplicate() {
	if ll.head == nil {
		fmt.Println("Список пуст.")
		return
	}

	seen := make(map[int]bool) // создаем мапу для отслеживания значений
	current := ll.head
	seen[current.value] = true

	for current != nil && current.next != nil {
		if seen[current.next.value] {
			// Если значение уже встречалось, удаляем узел
			current.next = current.next.next
		} else {
			// Если значение не встречалось, добавляем его в мапу
			seen[current.next.value] = true
			current = current.next
		}
	}
}
