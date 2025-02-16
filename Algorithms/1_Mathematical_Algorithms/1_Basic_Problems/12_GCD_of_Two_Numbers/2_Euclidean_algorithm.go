package main

/*
The idea of this algorithm is, the GCD of two numbers doesn’t change if
the smaller number is subtracted from the bigger number.
This is the Euclidean algorithm by subtraction.
It is a process of repeat subtraction, carrying the result forward each time
until the result is equal to any one number being subtracted.

Time Complexity: O(min(a,b))
Auxiliary Space: O(min(a,b)) because it uses internal stack data structure in recursion.
*/

func gcdEuclideanAlgorithm(a, b int) int {
	// Everything divides 0
	if a == 0 {
		return b
	}
	if b == 0 {
		return a
	}

	// Base case
	if a == b {
		return a
	}

	if a > b {
		return gcdEuclideanAlgorithm(a-b, b)
	}
	return gcdEuclideanAlgorithm(b-a, a)
}
