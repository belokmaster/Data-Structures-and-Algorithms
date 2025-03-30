package main

import (
	"fmt"
	"math"
	"sort"
)

// Дан массив arr[] из n целых чисел, где arr[i] обозначает количество шоколадок в i-м пакете.
// В каждом пакете может быть разное количество шоколадок. Есть m учеников, задача состоит в том,
// чтобы распределить пакеты с шоколадками так, чтобы:
// Каждый ученик получает ровно один пакетик. Разница между максимальным и минимальным
// количеством шоколадок в пакетиках, которые получают ученики, сводится к минимуму.
func findMinDiff(arr []int, m int) int {
	n := len(arr)
	if m > n {
		return -1 // Нельзя раздать больше пакетов, чем есть
	}

	sort.Ints(arr)
	midDiff := math.MaxInt

	for i := 0; i < n; i++ {
		// вычисляем разницу текущего окна
		diff := arr[i+m-1] - arr[i]
		if diff < midDiff {
			midDiff = diff
		}
	}
	return midDiff
}

func main() {
	arr := []int{7, 3, 2, 4, 9, 12, 56}
	fmt.Println(arr)
	ans := findMinDiff(arr, 2)
	fmt.Println(ans)
}
