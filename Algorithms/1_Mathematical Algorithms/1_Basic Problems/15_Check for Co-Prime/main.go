package main

import "fmt"

// Два числа A и B называются простейшими или взаимно простыми, если их наибольший общий делитель равен 1.
// Вам даны два числа A и B, найдите, являются ли они простейшими или нет.

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func coPrime(a, b int) bool {
	if gcd(a, b) == 1 {
		return true
	}

	return false
}

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	fmt.Println(coPrime(a, b))
}
