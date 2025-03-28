package main

import (
	"fmt"
	"math/rand"
)

// Дан массив arr[] целых чисел. Задача состоит в том, чтобы разделить массив на чётные и
// нечётные элементы. Разделение должно быть стабильным, то есть относительный порядок
// всех чётных элементов должен оставаться неизменным до и после разделения, и то же
// самое должно быть верно для всех нечётных элементов.
func partition(arr []int) []int {
	n := len(arr)
	tempEven := []int{}
	tempOdd := []int{}

	for i := 0; i < n; i++ {
		if arr[i]%2 == 0 {
			tempEven = append(tempEven, arr[i])
		} else {
			tempOdd = append(tempOdd, arr[i])
		}
	}

	tempEven = append(tempEven, tempOdd...)
	return tempEven
}

func main() {
	arr := []int{}
	for i := 0; i < 10; i++ {
		arr = append(arr, rand.Intn(100))
	}
	fmt.Println(arr)
	ans := partition(arr)
	fmt.Println(ans)
}
