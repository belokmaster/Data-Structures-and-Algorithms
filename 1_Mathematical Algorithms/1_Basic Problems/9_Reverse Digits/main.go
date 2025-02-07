package main

import (
	"slices"
	"strconv"
)

// Дано целое число n, найдите обратные его цифры.

func reverseDigits(n int) int {
	revNum := 0
	for n > 0 {
		revNum = revNum*10 + n%10
		n /= 10
	}
	return revNum
}

func reverseDigitsUsingString(n int) int {
	result := []rune(strconv.Itoa(n))
	slices.Reverse(result)
	finalResult, _ := strconv.Atoi(string(result))
	return finalResult
}
