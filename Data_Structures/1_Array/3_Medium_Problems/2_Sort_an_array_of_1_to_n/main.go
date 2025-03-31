package main

import (
	"fmt"
)

// Нам дан массив, содержащий от 1 до n элементов. Наша задача — эффективно отсортировать этот массив.
// O(N)
func cycleSort(arr []int) {
	i := 0
	for i < len(arr) {
		// Если элемент не на своём месте, меняем его
		if arr[i] != arr[arr[i]-1] {
			arr[i], arr[arr[i]-1] = arr[arr[i]-1], arr[i]
		} else {
			i++
		}
	}
}

func main() {
	arr := []int{}
	for i := 0; i < 10; i++ {
		arr = append(arr, i+1)
	}
	cycleSort(arr)
	fmt.Println(arr)
}
