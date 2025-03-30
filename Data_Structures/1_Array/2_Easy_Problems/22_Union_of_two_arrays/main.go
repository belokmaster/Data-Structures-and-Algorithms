package main

import (
	"fmt"
	"math/rand"
	"sort"
)

// Даны два массива a[] и b[], задача состоит в том, чтобы вернуть объединение обоих массивов
// в любом порядке. Примечание: объединение двух массивов — это массив, содержащий все
// уникальные элементы, присутствующие в обоих массивах.
func findUnion(a []int, b []int) []int {
	set := make(map[int]struct{})

	// Добавляем все элементы из a
	for _, num := range a {
		set[num] = struct{}{}
	}

	// Добавляем все элементы из b
	for _, num := range b {
		set[num] = struct{}{}
	}

	result := make([]int, 0, len(set))
	for num := range set {
		result = append(result, num)
	}

	sort.Ints(result)

	return result
}

func main() {
	arr := []int{}
	for i := 0; i < 10; i++ {
		arr = append(arr, rand.Intn(100))
	}
	arr2 := []int{}
	for i := 0; i < 10; i++ {
		arr2 = append(arr2, rand.Intn(100))
	}
	fmt.Println(arr)
	fmt.Println(arr2)
	ans := findUnion(arr, arr2)
	fmt.Println(ans)
}
