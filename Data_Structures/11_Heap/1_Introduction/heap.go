package heap

import (
	"fmt"
	"math"
)

// MinHeap представляет структуру мин-кучи
type MinHeap struct {
	arr      []int // Массив для хранения элементов кучи
	capacity int   // Максимальная ёмкость кучи
	heapSize int   // Текущий размер кучи
}

// NewMinHeap создает новый объект MinHeap с заданной максимальной ёмкостью
func NewMinHeap(capacity int) *MinHeap {
	return &MinHeap{
		arr:      make([]int, capacity),
		capacity: capacity,
		heapSize: 0,
	}
}

// Parent возвращает индекс родителя элемента в куче
func (h *MinHeap) Parent(i int) int {
	return (i - 1) / 2
}

// Left возвращает индекс левого потомка элемента в куче
func (h *MinHeap) Left(i int) int {
	return 2*i + 1
}

// Right возвращает индекс правого потомка элемента в куче
func (h *MinHeap) Right(i int) int {
	return 2*i + 2
}

// swap меняет местами два элемента в массиве кучи
func (h *MinHeap) swap(i, j int) {
	h.arr[i], h.arr[j] = h.arr[j], h.arr[i]
}

// MinHeapify восстанавливает свойства кучи для поддерева с корнем в i
func (h *MinHeap) MinHeapify(i int) {
	l := h.Left(i)  // Левый потомок
	r := h.Right(i) // Правый потомок
	smallest := i   // Предполагаем, что текущий элемент — минимальный

	// Проверяем, если левый потомок меньше текущего элемента
	if l < h.heapSize && h.arr[l] < h.arr[i] {
		smallest = l
	}
	// Проверяем, если правый потомок меньше текущего минимального элемента
	if r < h.heapSize && h.arr[r] < h.arr[smallest] {
		smallest = r
	}
	// Если находим меньший элемент, меняем местами и продолжаем восстановление кучи
	if smallest != i {
		h.swap(i, smallest)
		h.MinHeapify(smallest)
	}
}

// InsertKey вставляет новый ключ в мин-кучу
func (h *MinHeap) InsertKey(k int) {
	// Проверяем, не переполнилась ли куча
	if h.heapSize == h.capacity {
		fmt.Println("Overflow: Could not insertKey")
		return
	}

	// Вставляем новый элемент в конец кучи
	h.heapSize++
	i := h.heapSize - 1
	h.arr[i] = k

	// Восстанавливаем свойства мин-кучи, если они нарушены
	for i != 0 && h.arr[h.Parent(i)] > h.arr[i] {
		h.swap(i, h.Parent(i))
		i = h.Parent(i)
	}
}

// DecreaseKey уменьшает значение элемента в куче до newVal
func (h *MinHeap) DecreaseKey(i, newVal int) {
	// Устанавливаем новое значение для элемента
	h.arr[i] = newVal
	// Восстанавливаем свойства кучи, если они нарушены
	for i != 0 && h.arr[h.Parent(i)] > h.arr[i] {
		h.swap(i, h.Parent(i))
		i = h.Parent(i)
	}
}

// ExtractMin извлекает и возвращает минимальный элемент из кучи (корень)
func (h *MinHeap) ExtractMin() int {
	// Если куча пуста, возвращаем максимально возможное значение
	if h.heapSize <= 0 {
		return math.MaxInt
	}
	// Если в куче только один элемент, удаляем его
	if h.heapSize == 1 {
		h.heapSize--
		return h.arr[0]
	}

	// Сохраняем корень, меняем его с последним элементом и восстанавливаем кучу
	root := h.arr[0]
	h.arr[0] = h.arr[h.heapSize-1]
	h.heapSize--
	h.MinHeapify(0)

	return root
}

// DeleteKey удаляет элемент на заданном индексе
func (h *MinHeap) DeleteKey(i int) {
	// Сначала уменьшаем значение элемента до минус бесконечности, чтобы переместить его в корень
	h.DecreaseKey(i, math.MinInt)
	// Затем извлекаем минимальный элемент, который теперь и является удаленным
	h.ExtractMin()
}

// GetMin возвращает минимальный элемент (корень) кучи
func (h *MinHeap) GetMin() int {
	return h.arr[0]
}
