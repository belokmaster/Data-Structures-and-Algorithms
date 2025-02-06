package main

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func lcm(a, b int) int {
	return (a * b) / gcd(a, b)
}

// findValue вычисляет GCD(LCM(x, y), LCM(x, z))
func findValue(x, y, z int) int {
	g := gcd(y, z) // GCD(y, z)
	return lcm(x, g)
}
