package main

import (
	"fmt"
	"math/rand"
)

// Дан массив arr[] из n целых чисел. Создайте массив-произведение res[] (того же размера),
// в котором res[i] равно произведению всех элементов массива arr[], кроме arr[i].

func productExceptSelf(arr []int) []int {
	n := len(arr)
	zeroesCount := 0
	idx := -1
	prod := 1

	for i := 0; i < n; i++ {
		if arr[i] == 0 {
			zeroesCount++
			idx = i
		} else {
			prod *= arr[i]
		}
	}

	ans := make([]int, n)
	if zeroesCount == 0 {
		for i := 0; i < n; i++ {
			ans[i] = prod / arr[i]
		}
	} else if zeroesCount == 1 {
		ans[idx] = prod
	}

	return ans
}

func main() {
	arr := []int{}
	for i := 0; i < 10; i++ {
		arr = append(arr, rand.Intn(10))
	}
	fmt.Println(arr)
	fmt.Println(productExceptSelf(arr))
}
