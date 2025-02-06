package main

import "math"

//Time Complexity: O(log n)
//Auxiliary Space: O(log n)
func gcd(a, b float64) float64 {
	if a < b {
		return gcd(b, a)
	}

	// Базовый случай
	if math.Abs(b) < 0.001 {
		return a
	}

	return gcd(b, a-math.Floor(a/b)*b)
}
