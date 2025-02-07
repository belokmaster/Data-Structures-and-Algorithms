package main

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func largestGCD1Subset(arr []int, n int) int {
	curGCD := arr[0]
	for i := 1; i < n; i++ {
		curGCD = gcd(curGCD, arr[i])

		if curGCD == 1 {
			return n
		}
	}

	return 0
}
