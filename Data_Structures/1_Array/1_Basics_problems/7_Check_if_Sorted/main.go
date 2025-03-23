package main

import (
	"fmt"
	"math/rand"
	"sort"
)

// Проверить отсортирован ли массив
func checkIsSorted(arr []int) bool {
	for i := 1; i < len(arr); i++ {
		if arr[i-1] > arr[i] {
			return false
		}
	}
	return true
}

func main() {
	arr := []int{}
	for i := 0; i < 10; i++ {
		arr = append(arr, rand.Intn(100))
	}

	fmt.Println(arr)
	fmt.Println(checkIsSorted(arr))

	sort.Ints(arr)
	fmt.Println(checkIsSorted(arr))
}
