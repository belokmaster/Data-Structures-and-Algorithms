package main

import (
	"fmt"
)

// Дан массив целых чисел. Все числа встречаются по два раза, кроме одного, которое встречается один раз.
// Идея состоит в том, чтобы использовать XOR. XOR всех элементов массива даст число с одним вхождением.
func search(arr []int) int {
	n := len(arr)
	res := 0

	for i := 0; i < n; i++ {
		res = res ^ arr[i]
	}

	return res
}

func main() {
	arr := []int{0, 0, 1, 2, 2, 3, 4, 4, 3, 6, 7, 7, 6}
	fmt.Println(arr)
	ans := search(arr)
	fmt.Println(ans)
}
