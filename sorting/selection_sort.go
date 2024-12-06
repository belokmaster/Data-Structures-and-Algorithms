package main

// Функция сортировки выбором с использованием minIndex
func SelectionSort(arr []int) {
	n := len(arr)
	for i := 0; i < n; i++ {
		// Ищем минимальный элемент в неотсортированной части массива
		minIndex := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		// Если минимальный элемент не на текущей позиции, меняем их местами
		if minIndex != i {
			arr[i], arr[minIndex] = arr[minIndex], arr[i]
		}
	}
}
