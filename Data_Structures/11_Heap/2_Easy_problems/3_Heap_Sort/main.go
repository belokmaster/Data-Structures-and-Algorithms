package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Функция для перестройки кучи
func heapify(arr []int, n, i int) {
	largest := i
	left := 2*i + 1
	right := 2*i + 2

	// Если левый дочерний элемент больше корня
	if left < n && arr[left] > arr[largest] {
		largest = left
	}

	// Если правый дочерний элемент больше текущего наибольшего
	if right < n && arr[right] > arr[largest] {
		largest = right
	}

	// Если наибольший элемент не корень
	if largest != i {
		arr[i], arr[largest] = arr[largest], arr[i]

		// Рекурсивно перестроить поддерево
		heapify(arr, n, largest)
	}
}

// Основная функция для сортировки с использованием кучи
func heapSort(arr []int) {
	n := len(arr)

	// Построение кучи (перестройка массива)
	for i := n/2 - 1; i >= 0; i-- {
		heapify(arr, n, i)
	}

	// Поочередно извлекаем элементы из кучи
	for i := n - 1; i > 0; i-- {
		// Перемещаем корень в конец
		arr[0], arr[i] = arr[i], arr[0]

		// Перестраиваем кучу для уменьшенного массива
		heapify(arr, i, 0)
	}
}

// Функция для вывода массива
func printArray(arr []int) {
	for _, value := range arr {
		fmt.Printf("%d ", value)
	}
	fmt.Println()
}

// Основная функция
func main() {
	n := 20
	arr := make([]int, n)
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < n; i++ {
		x := rand.Intn(100)
		arr[i] = x
	}

	// Вызов функции сортировки
	heapSort(arr)

	fmt.Println("Sorted array is:")
	printArray(arr)
}
