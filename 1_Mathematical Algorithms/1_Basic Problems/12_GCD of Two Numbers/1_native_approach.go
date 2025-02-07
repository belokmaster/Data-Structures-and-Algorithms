package main

import "math"

/*
The basic idea is to find the minimum of the two numbers
and find its highest factor which is also a factor of the other number.

Time Complexity: O(min(a,b))
Auxiliary Space: O(1)
*/

func gcdNativeApproach(a, b int) int {
	result := int(math.Min(float64(a), float64(b)))
	for result > 0 {
		if a%result == 0 && b%result == 0 {
			break
		}
		result--
	}
	return result
}
