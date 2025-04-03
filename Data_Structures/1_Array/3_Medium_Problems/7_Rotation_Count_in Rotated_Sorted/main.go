package main

import (
	"fmt"
)

// Дан массив arr[], содержащий различные числа, отсортированные в порядке возрастания,
// и массив, повёрнутый вправо (то есть последний элемент циклически сдвигается в начало массива) k раз.
// Задача состоит в том, чтобы найти значение k.
func findMin(arr []int) int {
	low, high := 0, len(arr)-1

	for low < high {
		// Если подмассив уже отсортирован, минимум находится в low
		if arr[low] <= arr[high] {
			return low
		}

		mid := (low + high) / 2

		// Если правая половина не отсортирована, минимум в правой половине
		if arr[mid] > arr[high] {
			low = mid + 1
		} else {
			// Левая половина содержит минимум
			high = mid
		}
	}

	return low
}

func main() {
	arr := []int{5, 6, 1, 2, 3, 4}
	ans := findMin(arr)
	fmt.Println(ans)
}
