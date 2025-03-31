package main

import (
	"fmt"
)

// Даны два целочисленных массива одинакового размера, «arr[]» и «index[]».
// Переставьте элементы в «arr[]» в соответствии с заданным массивом индексов.
func reoder(arr, index []int, n int) {
	temp := make([]int, n)

	for i := 0; i < n; i++ {
		temp[index[i]] = arr[i]
	}

	copy(arr, temp)
}

func reorderExpect(arr, index []int, n int) {
	for i := 0; i < n; i++ {
		for index[i] != i {
			// Сохраняем целевые значения перед перемещением
			oldTargetI := index[index[i]]
			oldTargetE := arr[index[i]]

			// Размещаем arr[i] на его правильную позицию
			arr[index[i]] = arr[i]
			index[index[i]] = index[i]

			// Перемещаем старые значения в текущую позицию
			index[i] = oldTargetI
			arr[i] = oldTargetE
		}
	}
}

func main() {
	arr := []int{50, 40, 70, 60, 90}
	index := []int{3, 0, 4, 1, 2}
	reoder(arr, index, len(arr))
	fmt.Println(arr)
}
