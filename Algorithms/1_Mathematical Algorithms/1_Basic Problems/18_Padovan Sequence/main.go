package main

import "strconv"

// Задано число n, задача состоит в том, чтобы вернуть количество цифр в этом числе.

func cointDigit(n int) int {
	count := 0
	for n > 0 {
		n /= 10
		count++
	}
	return count
}

func fromIntToString(n int) int {
	result := strconv.Itoa(n)
	return len(result)
}
