package main

import (
	"fmt"
	"math/rand"
	"slices"
)

// перевернуть массив
func rotateArr(arr []int, d int) {
	n := len(arr)
	for i := 0; i < d; i++ {
		// Повернуть массив вправо на одну позицию
		last := arr[n-1]
		for j := n - 1; j > 0; j-- {
			arr[j] = arr[j-1]
		}
		arr[0] = last
	}
}

func rotateArrReversal(arr []int, d int) {
	n := len(arr)
	if n == 0 {
		return
	}

	// Обработать случай, если d > n
	d %= n
	if d == 0 {
		return
	}

	// Развернуть весь массив
	slices.Reverse(arr)
	// Развернуть первые d элементов
	slices.Reverse(arr[:d])
	// Развернуть оставшиеся n-d элементов
	slices.Reverse(arr[d:])
}

func main() {
	arr := []int{}
	for i := 0; i < 10; i++ {
		arr = append(arr, rand.Intn(100))
	}
	arr2 := make([]int, len(arr))
	copy(arr2, arr)

	fmt.Println(arr)
	fmt.Println(arr2)

	rotateArr(arr, 2)
	fmt.Println(arr)

	rotateArrReversal(arr2, 2)
	fmt.Println(arr2)
}
