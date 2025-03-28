package main

import (
	"fmt"
	"math/rand"
	"sort"
)

// Дан несортированный массив целых чисел. Отсортируйте его в виде волнового массива.
// Массив arr отсортирован в виде волны, если: arr[0] >= arr[1] <= arr[2] >= arr[3] <= arr[4] >= …..
func sortInWaveNaive(arr []int) {
	sort.Ints(arr)

	for i := 0; i < len(arr); i += 2 {
		arr[i], arr[i+1] = arr[i+1], arr[i]
	}
}

func sortInWave(arr []int) {
	n := len(arr)
	for i := 0; i < n; i += 2 {
		// если не первый
		if i > 0 && arr[i-1] > arr[i] {
			arr[i-1], arr[i] = arr[i], arr[i-1]
		}

		// если не ласт
		if i < n-1 && arr[i] < arr[i+1] {
			arr[i], arr[i+1] = arr[i+1], arr[i]
		}
	}
}

func main() {
	arr := []int{}
	for i := 0; i < 10; i++ {
		arr = append(arr, rand.Intn(100))
	}
	fmt.Println(arr)
	sortInWave(arr)
	fmt.Println(arr)
}
