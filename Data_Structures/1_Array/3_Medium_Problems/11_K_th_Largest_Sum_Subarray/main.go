package main

import (
	"fmt"
	"math/rand"
	"sort"
)

// Дан массив целых чисел. Напишите программу для поиска K-й по величине суммы смежных
// подмассивов в массиве чисел, в котором есть как отрицательные, так и положительные числа.

/*
	func kthLargestSum(arr []int, k int) int {
		var result []int

		// Генерация всех подмассивов и их сумм
		for i := 0; i < len(arr); i++ {
			sum := 0
			for j := i; j < len(arr); j++ {
				sum += arr[j]
				result = append(result, sum)
			}
		}

		// Сортировка по убыванию
		sort.Slice(result, func(i, j int) bool {
			return result[i] > result[j]
		})

		return result[k-1]
	}
*/
func kthLargestSum(arr []int, k int) int {
	n := len(arr)

	// Считаем префиксные суммы
	prefix := make([]int, n+1)
	for i := 0; i < n; i++ {
		prefix[i+1] = prefix[i] + arr[i]
	}

	// Собираем все суммы подмассивов
	var subarraySums []int
	for i := 0; i < n; i++ {
		for j := i + 1; j <= n; j++ {
			subarraySums = append(subarraySums, prefix[j]-prefix[i])
		}
	}

	// Сортируем по убыванию
	sort.Slice(subarraySums, func(i, j int) bool {
		return subarraySums[i] > subarraySums[j]
	})

	// Возвращаем k-й максимум (индексы с 0, но k передаётся с 1)
	return subarraySums[k-1]
}

func main() {
	arr := []int{}
	for i := 0; i < 10; i++ {
		arr = append(arr, rand.Intn(10)-3)
	}
	k := 3
	fmt.Println(arr)
	fmt.Println(kthLargestSum(arr, k))
}
