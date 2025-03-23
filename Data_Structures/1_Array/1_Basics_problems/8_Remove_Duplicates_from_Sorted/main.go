package main

import (
	"fmt"
	"math/rand"
	"sort"
)

// удалить дубликаты из отсортированного массива.
func removeDuplicates(arr []int) []int {
	newArr := []int{arr[0]}
	for i := 1; i < len(arr); i++ {
		if arr[i-1] != arr[i] {
			newArr = append(newArr, arr[i])
		}
	}
	return newArr
}

func main() {
	arr := []int{}
	for i := 0; i < 10; i++ {
		arr = append(arr, rand.Intn(100))
	}
	arr = append(arr, arr[0])
	arr = append(arr, arr[6])
	sort.Ints(arr)
	fmt.Println(arr)
	fmt.Println(removeDuplicates(arr))
}
