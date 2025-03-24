package main

import (
	"fmt"
	"math/rand"
)

// Дан массив размером n, заполненный числами от 1 до n-1 в случайном порядке.
// В массиве есть только один повторяющийся элемент. Задача состоит в том, чтобы
// найти повторяющийся элемент.
func findDuplicate(arr []int) int {
	n := len(arr)
	sum := 0
	for _, num := range arr {
		sum += num
	}
	expectedSum := (n - 1) * n / 2

	duplicate := sum - expectedSum
	return duplicate
}

func main() {
	arr := []int{}
	for i := 0; i < 10; i++ {
		arr = append(arr, rand.Intn(100))
	}
	fmt.Println(arr)
	ans := findDuplicate(arr)
	fmt.Println(ans)
}
