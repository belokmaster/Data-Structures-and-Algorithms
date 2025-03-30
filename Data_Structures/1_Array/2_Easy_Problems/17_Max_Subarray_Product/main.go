package main

import (
	"fmt"
	"math/rand"
)

// Учитывая массив arr[] размером n, задача состоит в том, чтобы вернуть индекс равновесия
// (если он есть) или -1, если индекс равновесия не существует. Индекс равновесия массива — это индекс,
// при котором сумма всех элементов с меньшими индексами равна сумме всех элементов с большими индексами.
func maxProductNaive(arr []int) int {
	n := len(arr)
	res := arr[0]
	for i := 1; i < n; i++ {
		mul := 1
		for j := i; j < n; j++ {
			mul *= arr[j]
			res = max(res, mul)
		}
	}

	return res
}

func maxProduct(arr []int) int {
	n := len(arr)
	if n == 0 {
		return 0
	}

	// Максимальное и минимальное произведение на текущем индексе
	currMax, currMin := arr[0], arr[0]
	maxProd := arr[0]

	for i := 1; i < n; i++ {
		// Временная переменная для хранения старого currMax
		tempMax := max(arr[i], arr[i]*currMax, arr[i]*currMin)

		// Обновляем currMin перед изменением currMax
		currMin = min(arr[i], arr[i]*currMax, arr[i]*currMin)
		currMax = tempMax

		maxProd = max(maxProd, currMax)
	}

	return maxProd
}

func main() {
	arr := []int{}
	for i := 0; i < 10; i++ {
		arr = append(arr, rand.Intn(200)-100)
	}
	fmt.Println(arr)
	ans := maxProduct(arr)
	ans1 := maxProductNaive(arr)
	fmt.Println(ans)
	fmt.Println(ans1)
}
