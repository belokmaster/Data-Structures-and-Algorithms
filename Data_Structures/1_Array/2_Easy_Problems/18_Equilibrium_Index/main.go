package main

import (
	"fmt"
)

// Учитывая массив arr[] размером n, задача состоит в том, чтобы вернуть индекс равновесия (если он есть)
// или -1, если индекс равновесия не существует.
// Индекс равновесия массива — это индекс, при котором сумма всех элементов с меньшими индексами
// равна сумме всех элементов с большими индексами.
func findEquilibrium(arr []int) int {
	n := len(arr)
	pref := make([]int, n)
	suff := make([]int, n)

	pref[0] = arr[0]
	suff[n-1] = arr[n-1]

	// Вычислить сумму префикса для всех индексов
	// от начала до i
	for i := 1; i < n; i++ {
		pref[i] = pref[i-1] + arr[i]
	}

	// Вычисление суммы суффиксов для всех индексов
	// от i до конца массива
	for i := n - 2; i >= 0; i-- {
		suff[i] = suff[i+1] + arr[i]
	}

	for i := 0; i < n; i++ {
		if pref[i] == suff[i] {
			return i
		}
	}

	return -1
}

func main() {
	arr := []int{-7, 1, 5, 2, -4, 3, 0}
	fmt.Println(arr)
	ans := findEquilibrium(arr)
	fmt.Println(ans)
}
