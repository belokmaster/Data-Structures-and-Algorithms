package main

import (
	"fmt"
	"time"
)

// Задано число n, найдите сумму первых натуральных чисел.

// O(n)
func nativeApproachSum(n int) int {
	finalSum := 0
	for i := 1; i <= n; i++ {
		finalSum += i
	}
	return finalSum
}

// O(1)
func effectiveSolutionSum(n int) int {
	return n * (n - 1) / 2
}

func main() {
	var input int
	fmt.Scan(&input)

	start := time.Now()
	result := nativeApproachSum(input)
	duration := time.Since(start)

	fmt.Println("Результат:", result)
	fmt.Println("Время выполнения:", duration)
}
