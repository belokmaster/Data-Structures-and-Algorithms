package main

import (
	"fmt"
	"math/rand"
	"sort"
)

// Даны два отсортированных массива a[] и b[], задача состоит в том, чтобы вернуть объединение
// обоих массивов в отсортированном порядке.
// Объединение двух массивов — это массив, содержащий все уникальные элементы, присутствующие в
// обоих массивах. Входные массивы могут содержать дубликаты.
func findUnion(a []int, b []int) []int {
	res := []int{}
	n, m := len(a), len(b)
	i, j := 0, 0

	for i < n && j < m {
		// Пропускаем дубликаты в a
		if i > 0 && a[i] == a[i-1] {
			i++
			continue
		}
		// Пропускаем дубликаты в b
		if j > 0 && b[j] == b[j-1] {
			j++
			continue
		}
		// Добавляем меньший элемент в результат
		if a[i] < b[j] {
			res = append(res, a[i])
			i++
		} else if a[i] > b[j] {
			res = append(res, b[j])
			j++
		} else { // a[i] == b[j], добавляем один раз
			res = append(res, a[i])
			i++
			j++
		}
	}

	// Добавляем оставшиеся элементы из a, пропуская дубликаты
	for i < n {
		if i > 0 && a[i] == a[i-1] {
			i++
			continue
		}
		res = append(res, a[i])
		i++
	}

	// Добавляем оставшиеся элементы из b, пропуская дубликаты
	for j < m {
		if j > 0 && b[j] == b[j-1] {
			j++
			continue
		}
		res = append(res, b[j])
		j++
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
	ans := findUnion(arr, arr2)
	fmt.Println(ans)
}
