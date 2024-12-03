package main

import "fmt"

func (ll *LinkedList) RemoveElement(value int) {
	// Если список пуст, выводится сообщение об ошибке и выполнение функции завершается.
	if ll.head == nil {
		fmt.Println("Список пуст, элемент не может быть удален.")
		return
	}

	// Если головной узел содержит значение, которое нужно удалить
	if ll.head.value == value {
		ll.head = ll.head.next
		return
	}
	// Поиск узла, предшествующего узлу с нужным значением
	current := ll.head
	for current.next != nil && current.next.value != value {
		current = current.next
	}
	// Если узел найден, удаляем его
	if current.next != nil {
		current.next = current.next.next
	} else {
		fmt.Println("Элемент не найден.")
	}
}
