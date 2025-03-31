package main

import (
	"fmt"
	"math/rand"
	"sort"
)

// Дан массив arr[] из различных элементов. Найдите минимальное количество перестановок,
// необходимых для сортировки массива.
func minSwaps(arr []int) int {
	temp := make([]int, len(arr))
	copy(temp, arr)
	sort.Ints(temp)

	pos := make(map[int]int)
	for i, val := range arr {
		pos[val] = i
	}

	swaps := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] != temp[i] {
			ind := pos[temp[i]]

			arr[i], arr[ind] = arr[ind], arr[i]

			pos[arr[ind]] = ind
			pos[arr[i]] = i

			swaps++
		}
	}

	return swaps
}

func main() {
	arr := []int{}
	for i := 0; i < 10; i++ {
		arr = append(arr, rand.Intn(100))
	}
	arr[3] = 3
	fmt.Println(arr)
	ans := minSwaps(arr)
	fmt.Println(ans)
}
