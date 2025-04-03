package main

import (
	"fmt"
)

// Дан отсортированный массив arr[] из n различных элементов, который повёрнут
// в некоторой неизвестной точке. Задача состоит в том, чтобы найти в нём максимальный элемент
func findMinBinarySearch(arr []int) int {
	left := 0
	right := len(arr) - 1

	for left < right {
		if arr[left] < arr[right] {
			return arr[right]
		}

		mid := (left + right) / 2
		if arr[mid] > arr[right] {
			left = mid
		} else {
			right = mid - 1
		}
	}

	return arr[left]
}

func main() {
	arr := []int{5, 6, 1, 2, 3, 4}
	ans := findMinBinarySearch(arr)
	fmt.Println(ans)
}
