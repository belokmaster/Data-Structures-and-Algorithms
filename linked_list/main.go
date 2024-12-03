package main

import "fmt"

func main() {
	// Создаем новый список
	ll := GenerateRandomList(10, 20)
	fmt.Println("Список выглядит так: ")
	ll.Display()

	fmt.Println("Перевернутый список выглядит так: ")
	ll.Reverse()
	ll.Display()

	fmt.Println("Добавляем элемент 20 в список.")
	ll.Add(4, 20)
	ll.Display()
	ll.FindElement(20)

	fmt.Println("Список с НОД между двумя узлами:")
	ll.InsertGCD()
	ll.Display()

	middleNode := ll.Middle()
	fmt.Printf("Средний элемент: %d\n", middleNode.value)
	ll.Display()

	// Добавляем элементы в начало списка
	ll.PushFront(10)
	ll.PushFront(20)
	ll.PushFront(30)
	ll.Display() // Ожидается: 30 -> 20 -> 10 -> nil

	// Добавляем элементы в конец списка
	ll.PushBack(40)
	ll.PushBack(50)
	ll.Display() // Ожидается: 30 -> 20 -> 10 -> 40 -> 50 -> nil

	// Удаляем элемент из начала списка
	ll.PopFront()
	ll.Display() // Ожидается: 20 -> 10 -> 40 -> 50 -> nil

	// Удаляем элемент из конца списка
	ll.PopBack()
	ll.Display() // Ожидается: 20 -> 10 -> 40 -> nil

	// Добавляем элемент по индексу
	ll.Add(1, 25)
	ll.Display() // Ожидается: 20 -> 25 -> 10 -> 40 -> nil

	// Удаляем элемент по индексу
	ll.Remove(2)
	ll.Display() // Ожидается: 20 -> 25 -> 40 -> nil

	// Удаляем все элементы
	ll.DeleteAll()
	ll.Display() // Ожидается: nil
}
