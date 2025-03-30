package main

import (
	"fmt"
	"math/rand"
)

// Даны два массива a[] и b[], задача состоит в том, чтобы найти пересечение этих двух массивов.
// Пересечением двух массивов называются элементы, которые встречаются в обоих массивах.
// Пересечение не должно содержать повторяющиеся элементы, а результат должен содержать элементы
// в любом порядке.
func intersection(a []int, b []int) []int {
	setA := make(map[int]struct{})
	result := []int{}

	for _, num := range a {
		setA[num] = struct{}{}
	}

	for _, num := range b {
		if _, exists := setA[num]; exists {
			result = append(result, num)
			delete(setA, num)
		}
	}

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
	ans := intersection(arr, arr2)
	fmt.Println(ans)
}
