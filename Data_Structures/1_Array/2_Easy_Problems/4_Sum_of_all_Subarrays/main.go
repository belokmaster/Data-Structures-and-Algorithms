package main

import (
	"fmt"
	"math/rand"
)

// найдите сумму всех подмассивов данного массива.
func subarraySum(arr []int) int {
	n := len(arr)
	res := 0

	for i := 0; i < n; i++ {
		temp := 0
		for j := i; j < n; j++ {
			temp += arr[j]
			res += temp
		}
	}

	return res
}

func subarraySumFormula(arr []int) int {
	n := len(arr)
	res := 0

	for i := 0; i < n; i++ {
		res += (arr[i] * (i + 1) * (n - i))
	}

	return res
}

func main() {
	arr := []int{}
	for i := 0; i < 10; i++ {
		arr = append(arr, rand.Intn(100))
	}
	fmt.Println(arr)
	ans := subarraySum(arr)
	ans2 := subarraySumFormula(arr)
	fmt.Println(ans, ans2)
}
