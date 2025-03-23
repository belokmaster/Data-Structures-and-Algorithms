package main

import (
	"fmt"
	"math/rand"
)

// Дано множество arr[], задача состоит в том, чтобы сгенерировать все возможные подмножества данного множества.
func subArray(arr []int) {
	n := len(arr)
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			for k := i; k <= j; k++ {
				fmt.Print(arr[k], " ")
			}
			fmt.Println()
		}
	}
}

func main() {
	arr := []int{}
	for i := 0; i < 10; i++ {
		arr = append(arr, rand.Intn(100))
	}
	fmt.Println(arr)
	subArray(arr)
}
