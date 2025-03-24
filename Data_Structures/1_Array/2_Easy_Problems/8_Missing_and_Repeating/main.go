package main

import (
	"fmt"
	"math"
	"math/rand"
)

// Дан несортированный массив размером n. Элементы массива находятся в диапазоне от 1 до n.
// Одно число из множества {1, 2, …, n} отсутствует, а одно число встречается в массиве дважды. Задача состоит в том, чтобы найти эти два числа.
func findTwoElement(arr []int) (int, int) {
	n := len(arr)

	repeating := -1

	// Поиск повторяющегося элемента
	for i := 0; i < n; i++ {
		val := int(math.Abs(float64(arr[i])))

		if arr[val-1] > 0 {
			arr[val-1] = -arr[val-1]
		} else {
			repeating = val // Повторяющийся элемент
		}
	}

	missing := -1

	// Поиск пропущенного элемента
	for i := 0; i < n; i++ {
		if arr[i] > 0 {
			missing = i + 1
			break
		}
	}

	return repeating, missing
}

func main() {
	arr := []int{}
	for i := 0; i < 10; i++ {
		arr = append(arr, rand.Intn(100))
	}
	fmt.Println(arr)
	ans, ans2 := findTwoElement(arr)
	fmt.Println(ans, ans2)
}
