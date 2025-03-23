package main

import (
	"fmt"
	"math/rand"
)

func findMax(arr []int) int {
	maxElement := arr[0]
	for i := 1; i < len(arr); i++ {
		if maxElement < arr[i] {
			maxElement = arr[i]
		}
	}
	return maxElement
}

func main() {
	arr := []int{}
	for i := 0; i < 10; i++ {
		arr = append(arr, rand.Intn(100))
	}

	ans := findMax(arr)
	fmt.Println(arr)
	fmt.Println(ans)
}
