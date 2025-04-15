package main

import (
	"fmt"
	"math"
	"math/rand"
	"runtime"
	"time"
)

// для небольшого массива норм
func bubbleSort(arr []int) {
	n := len(arr)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
}

func main() {
	var n int
	n = rand.Intn(5001) + 1

	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = rand.Intn(int(math.Pow(10, 9)))
	}

	runtime.GC()

	var memStart runtime.MemStats
	runtime.ReadMemStats(&memStart)

	start := time.Now()
	bubbleSort(arr)
	duration := time.Since(start)

	var memEnd runtime.MemStats
	runtime.ReadMemStats(&memEnd)

	memoryUsed := memEnd.Alloc - memStart.Alloc
	memoryTotal := memEnd.TotalAlloc - memStart.TotalAlloc

	fmt.Printf("Количество элементов в массиве: %d\n", len(arr))
	fmt.Printf("Время выполнения: %v\n", duration)
	fmt.Printf("Использовано памяти (Alloc): %d байт (%.2f KB)\n", memoryUsed, float64(memoryUsed)/1024)
	fmt.Printf("Всего выделено памяти (TotalAlloc): %d байт (%.2f KB)\n", memoryTotal, float64(memoryTotal)/1024)
}
