package main

import "fmt"

// isHeap проверяет, является ли массив максимальной кучей
func isHeap(arr []int, n int) bool {
	// Начинаем с корня и идем до последнего внутреннего узла
	for i := 0; i <= (n-2)/2; i++ {
		// Если левый ребенок больше, возвращаем false
		if arr[2*i+1] > arr[i] {
			return false
		}

		// Если правый ребенок больше, возвращаем false
		if 2*i+2 < n && arr[2*i+2] > arr[i] {
			return false
		}
	}
	return true
}

func main() {
	arr := []int{90, 15, 10, 7, 12, 2, 7, 3}
	n := len(arr)

	if isHeap(arr, n) {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
