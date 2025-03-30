package main

import (
	"fmt"
	"math/rand"
	"sort"
)

// Даны два отсортированных массива a[] и b[], задача состоит в том, чтобы вернуть их пересечение.
// Пересечением двух массивов называются элементы, которые есть в обоих массивах.
// Пересечение не должно содержать повторяющиеся элементы, а результат должен содержать элементы
// в отсортированном порядке.
func intersection(a []int, b []int) []int {
	res := []int{}
	m, n := len(a), len(b)
	i, j := 0, 0

	// Два указателя, аналог merge sort
	for i < m && j < n {
		// Пропускаем дубликаты в a
		if i > 0 && a[i] == a[i-1] {
			i++
			continue
		}

		// Пропускаем меньший элемент
		if a[i] < b[j] {
			i++
		} else if a[i] > b[j] {
			j++
		} else { // a[i] == b[j]
			res = append(res, a[i])
			i++
			j++
		}
	}

	return res
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
	sort.Ints(arr)
	sort.Ints(arr2)

	fmt.Println(arr)
	fmt.Println(arr2)
	ans := intersection(arr, arr2)
	fmt.Println(ans)
}
