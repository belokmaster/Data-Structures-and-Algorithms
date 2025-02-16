package main

import "fmt"

// Дано число n (n >= 0), найдите его факториал .

func factorial(n int) int {
	res := 1

	for i := 2; i <= n; i++ {
		res *= i
	}

	return res
}

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Println(factorial(n))
}
