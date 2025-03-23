package main

import (
	"fmt"
	"math/rand"
)

// Дан массив целых чисел arr[], задача состоит в том, чтобы переместить все нули в
// конец массива, сохранив относительный порядок всех ненулевых элементов.
func pushZerosToEnd(arr []int) {
	count := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] != 0 {
			arr[count] = arr[i]
			count++
		}
	}

	for count < len(arr) {
		arr[count] = 0
		count++
	}
}

func main() {
	arr := []int{}
	for i := 0; i < 10; i++ {
		arr = append(arr, rand.Intn(100))
	}
	arr[2] = 0
	arr[5] = 0

	fmt.Println(arr)
	pushZerosToEnd(arr)
	fmt.Println(arr)
}
