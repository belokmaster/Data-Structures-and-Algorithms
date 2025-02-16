package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Дан массив arr[] неотрицательных чисел, задача состоит в том, чтобы найти НОК всех элементов массива.

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

// Функция для вычисления НОК (lcm) двух чисел
func lcm(a, b int) int {
	return (a * b) / gcd(a, b)
}

// Функция для вычисления НОК всех элементов массива
func findLCM(arr []int) int {
	ans := arr[0]

	for _, num := range arr {
		ans = lcm(ans, num)
	}

	return ans
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

	fmt.Println(findLCM(arr))
}
