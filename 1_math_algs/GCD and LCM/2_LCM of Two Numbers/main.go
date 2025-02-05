package main

import "math"

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

// use gcd to calc lcm
// Time Complexity: O(log(min(a,b))
// Auxiliary Space: O(log(min(a,b))
func lcmByGCD(a, b int) int {
	return (a * b) / gcd(a, b)
}

// Time Complexity: O(min(a,b))
// Auxiliary Space: O(1)
func lcm(a, b int) int {
	greater := int(math.Max(float64(a), float64(b)))
	smallest := int(math.Min(float64(a), float64(b)))
	for i := greater; ; i += greater {
		if i%smallest == 0 {
			return i
		}
	}
}
