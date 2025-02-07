package main

/*
If we notice the previous approach, we can see at some point, one number becomes
a factor of the other so instead of repeatedly subtracting till both become equal,
we can check if it is a factor of the other.

Time Complexity: O(min(a, b))
Auxiliary Space: O(1)
*/

func gcdEuclideanAlgorithmOptimizationCheckingDivisibility(a, b int) int {
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
		if a%b == 0 {
			return b
		}
		return gcdEuclideanAlgorithmOptimizationCheckingDivisibility(a-b, b)
	}

	if b%a == 0 {
		return a
	}
	return gcdEuclideanAlgorithmOptimizationCheckingDivisibility(b-a, a)
}
