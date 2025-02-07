package main

func check(x int) bool {
	sum := 0
	for x > 0 {
		sum += x % 10
		x /= 10
	}
	return sum%3 == 0
}

func checkDivisibilityForThree(x int) bool {
	return x%3 == 0
}
