package main

import (
	"fmt"
	"slices"
	"sort"
)

// Дан массив arr[] размером n. Задача состоит в том, чтобы найти все лидирующие элементы в массиве.
// Элемент является лидирующим, если он больше или равен всем элементам справа от него.

// Идея состоит в том, чтобы просмотреть все элементы массива справа налево и отслеживать
// максимальное значение. Когда максимальное значение меняется, добавьте его к результату.
// Наконец, переверните результат
func findLeaders(arr []int) []int {
	if len(arr) == 0 {
		return nil
	}

	lead := arr[len(arr)-1]
	leaders := []int{lead}

	for i := len(arr) - 2; i >= 0; i-- {
		if arr[i] >= lead {
			lead = arr[i]
			leaders = append(leaders, lead)
		}
	}

	slices.Reverse(leaders)
	return leaders
}

func main() {
	arr := []int{16, 17, 4, 3, 5, 2}
	ans := findLeaders(arr)
	fmt.Println(arr)
	fmt.Println(ans)
	sort.Ints(arr)
	fmt.Println(arr)
}
