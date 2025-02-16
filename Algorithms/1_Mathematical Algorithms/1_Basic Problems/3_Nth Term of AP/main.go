package main

// Учитывая первый член (a), общую разность (d) и целое число N ряда
// арифметической прогрессии, найти N-й член ряда.

// O(n)
func loopMethod(a, d, n int) int {
	nthTerm := a
	for i := 1; i < n; i++ {
		nthTerm += d
	}
	return nthTerm
}

// O(1)
func effectiveMethod(a, d, n int) int {
	return (a + (n-1)*d)
}
