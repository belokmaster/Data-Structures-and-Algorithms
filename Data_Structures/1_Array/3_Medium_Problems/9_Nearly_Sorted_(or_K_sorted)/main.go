package main

import (
	"fmt"
	"math/rand"
	"sort"
)

// Дан массив arr[] и число k. Массив отсортирован таким образом, что каждый элемент находится
// на расстоянии не более k от своей отсортированной позиции. Это означает, что если мы полностью
// отсортируем массив, то индекс элемента может находиться в диапазоне от i – k до i + k,
// где i — индекс в данном массиве. Наша задача — полностью отсортировать массив.
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
