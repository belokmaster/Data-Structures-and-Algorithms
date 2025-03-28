package main

import (
	"fmt"
	"math/rand"
)

// Дан массив arr[] целых чисел. Задача состоит в том, чтобы упорядочить их таким образом,
// чтобы все отрицательные числа оказались перед всеми положительными числами в массиве,
// не используя никаких дополнительных структур данных, таких как хеш-таблица, массивы и т. д.
// Порядок следования должен быть сохранён.
func rearrange(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		if arr[i] > 0 {
			continue
		}

		temp := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > 0 {
			arr[j+1] = arr[j]
			j--
		}

		arr[j+1] = temp
	}

	return arr
}

func main() {
	arr := []int{}
	for i := 0; i < 10; i++ {
		arr = append(arr, rand.Intn(100))
	}
	fmt.Println(arr)
	ans := rearrange(arr)
	fmt.Println(ans)
}
