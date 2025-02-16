package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Дан массив arr[] неотрицательных чисел, задача состоит в том, чтобы найти НОД всех элементов массива.

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func findGCD(arr []int) int {
	res := arr[0]

	for _, num := range arr {
		res = gcd(res, num)
		if res == 1 {
			return 1
		}
	}

	return res
}

func main() {
	fmt.Print("Введите числа через пробел: ")
	var input string
	fmt.Scanln(&input)

	parts := strings.Split(input, " ") // Разбиваем строку по пробелам
	arr := make([]int, len(parts))

	for i, p := range parts {
		num, _ := strconv.Atoi(p)
		arr[i] = num
	}

	fmt.Println(findGCD(arr))
}
