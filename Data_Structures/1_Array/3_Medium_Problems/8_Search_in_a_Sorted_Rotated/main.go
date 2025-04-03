package main

import (
	"fmt"
	"math/rand"
	"sort"
)

// Дан отсортированный и повёрнутый массив arr[] из n различных элементов.
// Задача состоит в том, чтобы найти индекс заданного ключа в массиве.
// Если ключа нет в массиве, верните -1.
func search(arr []int, key int) int {
	left, right := 0, len(arr)-1

	for left <= right {
		mid := left + (right-left)/2

		if arr[mid] == key {
			return mid
		}

		if arr[mid] >= arr[left] {
			if key >= arr[left] && key < arr[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			if key > arr[mid] && key <= arr[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}

	// Ключ не найден
	return -1
}
func main() {
	arr := []int{}
	for i := 0; i < 10; i++ {
		arr = append(arr, rand.Intn(100))
	}
	sort.Ints(arr)
	fmt.Println(search(arr, arr[3]))
}
