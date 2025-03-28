package main

import (
	"fmt"
	"math"
	"math/rand"
)

// Дан массив из n целых чисел. Найдите три элемента, для которых a[i] < a[j] < a[k]
// и i < j < k, за время 0(n). Если таких троек несколько, выведите любую из них.
func find3Numbers(arr []int) []int {
	// Если элементов меньше 3, тройка невозможна
	if len(arr) < 3 {
		return nil
	}

	// Минимальное число в массиве
	minNum := arr[0]

	// Наименьшее максимальное число в лучшей последовательности
	maxSeq := math.MaxInt32

	// Сохранение arr[i]
	storeMin := minNum

	for i := 1; i < len(arr); i++ {
		if arr[i] == minNum {
			continue
		} else if arr[i] < minNum {
			minNum = arr[i]
			continue
		} else if arr[i] < maxSeq {
			// Данное условие срабатывает, когда текущая последовательность = 2
			maxSeq = arr[i]
			storeMin = minNum
		} else if arr[i] > maxSeq {
			return []int{storeMin, maxSeq, arr[i]}
		}
	}

	return nil
}

func main() {
	arr := []int{}
	for i := 0; i < 10; i++ {
		arr = append(arr, rand.Intn(100))
	}
	fmt.Println(arr)
	ans := find3Numbers(arr)
	fmt.Println(ans)
}
