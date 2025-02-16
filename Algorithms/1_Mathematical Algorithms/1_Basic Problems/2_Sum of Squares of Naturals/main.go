package main

import (
	"fmt"
	"time"
)

// Дано n, найти сумму квадратов первых n натуральных чисел.

// O(n)
func nativeApproachSummation(n int) int {
	finalSum := 0
	for i := 1; i <= n; i++ {
		finalSum += i * i
	}
	return finalSum
}

// O(1)
func effectiveSolutionSum(n int) int {
	return (n * (n + 1) * (2*n + 1)) / 6
}

// O(1)
func effectiveSolutionSumWithoutOverflow(n int) int {
	return (n * (n + 1) / 2) * (2*n + 1) / 3
}

func main() {
	var input int
	fmt.Scan(&input)

	start := time.Now()
	result := effectiveSolutionSumWithoutOverflow(input)
	duration := time.Since(start)

	fmt.Println("Результат:", result)
	fmt.Println("Время выполнения:", duration)
}
