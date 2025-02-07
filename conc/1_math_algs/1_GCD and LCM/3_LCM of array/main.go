package main

/*
lcm(a, b) = (a * b) / gcd(a, b)

Time Complexity: O(n * log(min(a, b))), where n represents the size of the given array.
Auxiliary Space: O(n*log(min(a, b))) due to recursive stack space.
*/

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func findLCM(arr []int) int {
	ans := arr[0]

	// ans contains LCM of arr[0], ..arr[i]
	// after i'th iteration,
	for i := 1; i < len(arr); i++ {
		ans = (arr[i] * ans) / gcd(arr[i], int(ans))
	}

	return ans
}
