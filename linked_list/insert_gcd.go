package main

// gcd вычисляет наибольший общий делитель (НОД) двух чисел
func gcd(a, b int) int {
	// Если второе число равно 0, то НОД это первое число
	if b == 0 {
		return a
	}
	// Рекурсивный вызов функции для нахождения НОД
	return gcd(b, a%b)
}

// InsertGCD вставляет новый узел с НОД значений соседних узлов в связанный список
func (ll *LinkedList) InsertGCD() {
	current := ll.head // Начинаем с головы списка

	// Проходим по списку, пока текущий и следующий узлы не равны nil
	for current != nil && current.next != nil {
		// Вычисляем НОД значений текущего узла и следующего
		gcdValue := gcd(current.value, current.next.value)

		// Создаем новый узел с вычисленным НОД
		newNode := &Node{value: gcdValue}

		// Новый узел должен следовать за текущим
		newNode.next = current.next
		current.next = newNode

		// Переходим к следующему узлу после вставленного
		current = newNode.next
	}
}
