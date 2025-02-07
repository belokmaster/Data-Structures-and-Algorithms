package main

import "math"

// Учитывая первый член a, общее отношение r и целое число N ряда
// геометрической прогрессии, найти N-й член ряда.

// O(n)
func loopMethod(a, r, n int) int {
	nthTerm := a
	for i := 1; i < n; i++ {
		nthTerm *= r
	}
	return nthTerm
}

// O(1)
func effectiveMethod(a, r, n int) int {
	return a * int(math.Pow(float64(r), float64(n-1)))
}
