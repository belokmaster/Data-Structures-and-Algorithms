package main

import (
	"fmt"
	"math/rand"
)

// Дан массив arr из n целых чисел и целое число x.
// Определите, присутствует ли элемент x в массиве.
// Верните индекс первого вхождения x в массив или -1, если его нет.

func search(array []int, target int) int {
	for i, num := range array {
		if num == target {
			return i
		}
	}
	return -1
}

func main() {
	arr := []int{}
	for i := 0; i < 10; i++ {
		arr = append(arr, rand.Intn(100))
	}

	ans := search(arr, 10)
	fmt.Println(arr)
	fmt.Println(ans)
}
