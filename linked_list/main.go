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

	fmt.Println("Добавляем элемент 40 в список. Попробуем удалить его.")
	ll.Add(6, 40)
	ll.Display()
	ll.RemoveElement(40)
	ll.Display()
}
