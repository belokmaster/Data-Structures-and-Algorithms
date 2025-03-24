package main

import (
	"fmt"
)

// Дан массив arr[] размером n-1 с различными целыми числами в диапазоне [1, n].
// Этот массив представляет собой перестановку целых чисел от 1 до n с одним отсутствующим элементом.
// Найдите отсутствующий элемент в массиве.
func missingNum(arr []int) int {
	n := len(arr) + 1

	actualSum := 0
	for _, num := range arr { // Суммируем элементы массива
		actualSum += num
	}

	expectedSum := (n * (n - 1)) / 2 // Используем n*(n-1)/2, так как числа идут от 0 до n-1
	return expectedSum - actualSum
}

func missingNumXOR(arr []int) int {
	n := len(arr) + 1

	xor1 := 0
	for _, num := range arr { // XOR всех элементов массива
		xor1 ^= num
	}

	xor2 := 0
	for i := 0; i < n; i++ { // XOR всех чисел от 0 до n-1
		xor2 ^= i
	}

	return xor1 ^ xor2
}

func main() {
	arr := []int{0, 1, 2, 4, 5}
	fmt.Println(arr)
	ans := missingNumXOR(arr)
	ans2 := missingNum(arr)
	fmt.Println(ans2)
	fmt.Println(ans)
}
