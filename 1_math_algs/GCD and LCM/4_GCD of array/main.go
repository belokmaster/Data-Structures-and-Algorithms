package main

// iterative implementation
func getGCD(a, b int) int {
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

func gcdArray(arr []int) int {
	gcd := 1
	for i := 0; i < len(arr); i++ {
		gcd = getGCD(gcd, arr[i])
	}
	return gcd
}
