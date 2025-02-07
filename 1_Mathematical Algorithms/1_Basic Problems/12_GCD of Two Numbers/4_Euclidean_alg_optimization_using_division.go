package main

/*
Instead of the Euclidean algorithm by subtraction, a better approach can be used.
We don’t perform subtraction here. we continuously divide the bigger number by the smaller number.

Time Complexity: O(log(min(a,b)))
Auxiliary Space: O(log(min(a,b))
*/

func gcdEuclideanAlgOptimizationUsingDivision(a, b int) int {
	if b == 0 {
		return a
	}
	return gcdEuclideanAlgOptimizationUsingDivision(b, a%b)
}
