package main

import (
	"fmt"
	"math/rand"
)

// Дан массив элементов длиной n, от 0 до n – 1. Все элементы могут отсутствовать в массиве.
// Если элемент отсутствует, то в массиве будет -1. Переставьте элементы массива так,
// чтобы arr[i] = i, а если i отсутствует, то выведите -1 на это место.
func modifyArray(arr []int) {
	n := len(arr)
	temp := make([]int, n) // Создаем временный массив

	for i := range temp {
		temp[i] = -1 // Инициализируем -1
	}

	// Заполняем temp корректными значениями
	for _, val := range arr {
		if val != -1 && val < n {
			temp[val] = val
		}
	}

	// Обновляем исходный массив
	copy(arr, temp)
}

// modifyArray переставляет элементы так, чтобы arr[i] = i, заполняя -1, где нужно.
func modifyArrayExpected(arr []int) {
	i := 0
	for i < len(arr) {
		// Если arr[i] не -1 и не находится на "правильном" месте, меняем местами
		if arr[i] != -1 && arr[i] < len(arr) && arr[i] != arr[arr[i]] {
			arr[i], arr[arr[i]] = arr[arr[i]], arr[i] // Swap элементов
		} else {
			// Если элемент уже на своем месте или -1, просто двигаемся дальше
			i++
		}
	}
}

func main() {
	arr := []int{}
	for i := 0; i < 10; i++ {
		arr = append(arr, rand.Intn(100))
	}
	arr[3] = 3
	fmt.Println(arr)
	modifyArray(arr)
	fmt.Println(arr)
}
