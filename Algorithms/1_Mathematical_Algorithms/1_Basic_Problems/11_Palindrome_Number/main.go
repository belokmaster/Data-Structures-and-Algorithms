package main

import (
	"slices"
	"strconv"
)

// Дано положительное целое число n , определите, является ли число палиндромом или нет.
// Число является палиндромом, если оно остается прежним при перестановке цифр.

func checkPalindrome(n int) bool {
	reversedDigit := 0
	temp := n

	for temp > 0 {
		reversedDigit = reversedDigit*10 + temp%10
		temp /= 10
	}

	return n == reversedDigit
}

func checkPalindromeUsingString(n int) bool {
	reversedDigit := []rune(strconv.Itoa(n))
	slices.Reverse(reversedDigit)
	reversedDigitInt, _ := strconv.Atoi(string(reversedDigit))
	return reversedDigitInt == n
}
