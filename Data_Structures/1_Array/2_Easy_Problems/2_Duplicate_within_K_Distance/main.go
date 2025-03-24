package main

import (
	"fmt"
	"math/rand"
)

// Дан целочисленный массив arr[] и целое число k. Определите, существуют ли два индекса i и j,
// таких что arr[i] == arr[j] и |i – j| ≤ k. Если такая пара существует, верните «Да»,
// иначе верните «Нет».
func checkDuplicatesWithinK(arr []int, k int) bool {
	n := len(arr)
	for i := 0; i < n; i++ {
		for c := 1; c <= k && (i+c) < n; c++ {
			j := i + c
			if arr[i] == arr[j] {
				return true
			}
		}
	}
	return false
}

func main() {
	arr := []int{}
	for i := 0; i < 10; i++ {
		arr = append(arr, rand.Intn(100))
	}
	fmt.Println(arr)
	ans := checkDuplicatesWithinK(arr, 4)
	fmt.Println(ans)
}
