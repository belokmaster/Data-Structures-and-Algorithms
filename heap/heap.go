package main

import "fmt"

// MinHeap представляет структуру мин-кучи.
type MinHeap struct {
	elements []int // Срез, хранящий элементы кучи
}

// Insert добавляет элемент в мин-кучу.
func (h *MinHeap) Insert(val int) {
	h.elements = append(h.elements, val) // Добавляем новый элемент в конец среза
	h.SiftUp(len(h.elements) - 1)        // Восстанавливаем свойства кучи с помощью SiftUp
}

// SiftUp восстанавливает свойства кучи после добавления элемента.
func (h *MinHeap) SiftUp(index int) {
	for index > 0 { // Пока не достигнем корня кучи
		parentIndex := (index - 1) / 2 // Индекс родителя текущего узла
		if h.elements[parentIndex] <= h.elements[index] {
			break // Если родитель меньше или равен текущему, выходим из цикла
		}
		// Меняем местами родителя и текущий элемент
		h.elements[parentIndex], h.elements[index] = h.elements[index], h.elements[parentIndex]
		index = parentIndex // Переходим на уровень выше
	}
}

// ExtractMin извлекает минимальный элемент из мин-кучи.
func (h *MinHeap) ExtractMin() (int, error) {
	if len(h.elements) == 0 {
		return 0, fmt.Errorf("Пустая мин-куча.") // Обработка случая, когда куча пуста
	}

	min := h.elements[0]                  // Запоминаем минимальный элемент (корень)
	lastIndex := len(h.elements) - 1      // Индекс последнего элемента
	h.elements[0] = h.elements[lastIndex] // Перемещаем последний элемент на место корня
	h.elements = h.elements[:lastIndex]   // Удаляем последний элемент (он теперь перемещён)
	h.SiftDown(0)                         // Восстанавливаем свойства кучи с помощью SiftDown

	return min, nil // Возвращаем минимальный элемент
}

// SiftDown восстанавливает свойства кучи после извлечения элемента.
func (h *MinHeap) SiftDown(index int) {
	lastIndex := len(h.elements) - 1 // Индекс последнего элемента
	leftChildIndex := 2*index + 1    // Индекс левого дочернего элемента
	rightChildIndex := 2*index + 2   // Индекс правого дочернего элемента

	for leftChildIndex <= lastIndex { // Пока существует хотя бы один дочерний элемент
		smallerChildIndex := leftChildIndex // Предполагаем, что левый дочерний элемент — меньший
		// Проверяем, есть ли правый дочерний элемент и меньше ли он левого
		if rightChildIndex <= lastIndex && h.elements[rightChildIndex] < h.elements[leftChildIndex] {
			smallerChildIndex = rightChildIndex // Обновляем индекс меньшего дочернего элемента
		}

		if h.elements[index] <= h.elements[smallerChildIndex] {
			break // Если текущий узел меньше или равен меньшему дочернему узлу, выходим из цикла
		}

		// Меняем местами родителя и меньшего дочернего элемента
		h.elements[index], h.elements[smallerChildIndex] = h.elements[smallerChildIndex], h.elements[index]
		index = smallerChildIndex    // Переходим на уровень ниже
		leftChildIndex = 2*index + 1 // Обновляем индексы дочерних элементов
		rightChildIndex = 2*index + 2
	}
}
