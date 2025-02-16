package main

import (
	"fmt"
	"math"
)

// Учитывая n, подсчитайте все ‘a’ и ‘b’, которые удовлетворяют условию a³ + b³ = n
// Где (a, b) и (b, a) считаются двумя разными парами

// Временная сложность: O(n^1/3)
func countPairs(n int) int {
	count := 0

	// Проверяем все числа от 1 до кубического корня из n
	for i := 1; i*i*i <= n; i++ {
		cb := i * i * i
		diff := n - cb

		// Проверяем, является ли разность идеальным кубом
		cbrtDiff := int(math.Round(math.Pow(float64(diff), 1.0/3.0)))

		if cbrtDiff*cbrtDiff*cbrtDiff == diff {
			count++
		}
	}

	return count
}

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Println(countPairs(n))
}
