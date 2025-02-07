package main

import "math"

// Учитывая три числа a, b и k, найдите k-ю цифру в a^b с правой стороны
// 3^3 = 27 for k = 1. First digit is 7 in 27

// O(log p)
func kthdigit(a, b, k int) int {
	p := int(math.Pow(float64(a), float64(b)))
	count := 0

	for p > 0 && count < k {
		rem := p % 10
		count++
		if count == k {
			return rem
		}
		p /= 10
	}

	return 0
}
