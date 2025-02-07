package main

/*
Iterative way to find the GCD of two numbers using Euclidean algorithm.

Time Complexity: O(log(min(a,b)))
Auxiliary Space: O(1)
*/

func gcdEuclideanAlg(a, b int) int {
	for a > 0 && b > 0 {
		if a > b {
			a %= b
		} else {
			b %= a
		}
	}

	if a == 0 {
		return b
	}

	return a
}
