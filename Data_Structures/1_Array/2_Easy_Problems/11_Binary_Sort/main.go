package main

import (
	"fmt"
	"math/rand"
)

// Вам дан массив из нулей и единиц в случайном порядке.
func segregate0and1(arr []int) []int {
	zeroesCount := 0
	for _, x := range arr {
		if x == 0 {
			zeroesCount++
		}
	}

	for i := 0; i < zeroesCount; i++ {
		arr[i] = 0
	}

	for i := zeroesCount; i < len(arr); i++ {
		arr[i] = 1
	}

	return arr
}

func main() {
	arr := []int{}
	for i := 0; i < 10; i++ {
		arr = append(arr, rand.Intn(2))
	}
	fmt.Println(arr)
	ans := segregate0and1(arr)
	fmt.Println(ans)
}
