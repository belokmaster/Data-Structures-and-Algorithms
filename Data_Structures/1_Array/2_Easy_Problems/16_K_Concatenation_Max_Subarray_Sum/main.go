package main

import (
	"fmt"
	"math"
	"math/rand"
)

// Дан массив и число k. Найдите наибольшую сумму элементов в модифицированном массиве,
// который получается путём повторения заданного массива k раз.
func maxSubArraySumRepeated(arr []int, k int) int {
	n := len(arr)
	maxSoFar := math.MinInt
	maxEndingHere := 0

	for i := 0; i < n; i++ {
		maxEndingHere += arr[i%n]
		maxSoFar = max(maxSoFar, maxEndingHere)
		if maxEndingHere < 0 {
			maxEndingHere = 0
		}
	}

	return maxSoFar
}

func main() {
	arr := []int{}
	for i := 0; i < 10; i++ {
		arr = append(arr, rand.Intn(200)-100)
	}
	fmt.Println(arr)
	fmt.Println(maxSubArraySumRepeated(arr, 3))
}
