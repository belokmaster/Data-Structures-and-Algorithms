package main

// Дано целое положительное число n. Задача состоит в том, чтобы найти сумму всех чисел до n.

// O(n)
func seriousSum(n int) int {
	sum := 0
	for i := 1; i < n; i++ {
		sum += i * (i + 1) / 2
	}
	return sum
}

func effectiveSeriousSum(n int) int {
	return (n * (n + 1) * (n + 2)) / 6
}
