package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generateRandomArray(size int) []int {
	rng := rand.New(rand.NewSource(time.Now().UnixNano())) // Создание нового генератора с начальным значением
	arr := make([]int, size)                               // Создание массива заданного размера

	for i := 0; i < size; i++ {
		arr[i] = rng.Intn(100)
	}

	return arr
}

func main() {
	randomArr := generateRandomArray(10)
	fmt.Println("Массив до сортировки", randomArr)
	BubbleSort(randomArr)
	fmt.Println("После сортировки:", randomArr)

	randomArr = generateRandomArray(10)
	fmt.Println("Массив до сортировки", randomArr)
	InsertionSort(randomArr)
	fmt.Println("После сортировки:", randomArr)

	randomArr = generateRandomArray(10)
	fmt.Println("Массив до сортировки", randomArr)
	sortArray := CountingSort(randomArr)
	fmt.Println("После сортировки:", sortArray)
}
