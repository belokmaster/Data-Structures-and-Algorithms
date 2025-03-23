package main

import (
	"fmt"
	"math/rand"
	"sort"
)

// Дан массив положительных целых чисел arr[] размером n.
// Задача состоит в том, чтобы найти второй по величине неповторяющийся элемент в массиве.
func findSecondMax(arr []int) int {
	n := len(arr)
	if n == 1 {
		return arr[0]
	}

	firstMax := arr[0]
	secondMax := arr[1]
	if secondMax > firstMax {
		firstMax, secondMax = secondMax, firstMax
	}

	if n == 2 {
		return secondMax
	}

	for i := 2; i < n; i++ {
		if arr[i] > firstMax {
			temp := firstMax
			firstMax = arr[i]
			secondMax = temp
		} else if arr[i] > secondMax {
			secondMax = arr[i]
		}
	}

	return secondMax
}

func main() {
	arr := []int{}
	for i := 0; i < 10; i++ {
		arr = append(arr, rand.Intn(100))
	}

	ans := findSecondMax(arr)
	fmt.Println(arr)
	fmt.Println(ans)
	sort.Ints(arr)
	fmt.Println(arr)
}
