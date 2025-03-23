package main

import (
	"fmt"
	"math/rand"
)

// перевернуть массив
func reverseArray(arr []int) []int {
	left := 0
	right := len(arr) - 1

	for left < right {
		arr[left], arr[right] = arr[right], arr[left]
		left++
		right--
	}

	return arr
}

func main() {
	arr := []int{}
	for i := 0; i < 10; i++ {
		arr = append(arr, rand.Intn(100))
	}
	fmt.Println(arr)
	arr = reverseArray(arr)
	fmt.Println(arr)
}
