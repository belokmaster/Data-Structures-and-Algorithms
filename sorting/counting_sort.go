package main

import "math"

// min возвращает меньшее из двух чисел
func min(arr []int) int {
	min := math.MaxInt64
	for _, digit := range arr {
		if digit < min {
			min = digit
		}
	}
	return min
}

// max возвращает большее из двух чисел
func max(arr []int) int {
	max := math.MinInt64
	for _, digit := range arr {
		if digit > max {
			max = digit
		}
	}
	return max
}

// Функция Count для подсчета вхождений элемента в срез
func count(arr []int, item int) int {
	counts := 0
	for _, digit := range arr {
		if digit == item {
			counts++
		}
	}
	return counts
}

func CountingSort(arr []int) []int {
	minDigit := min(arr) // Находим минимальное значение
	maxDigit := max(arr) // Находим максимальное значение
	// Инициализируем срез для подсчета частоты появления каждого элемента
	count := make([]int, maxDigit-minDigit+1)

	// Заполняем срез count, где индекс будет соответствовать элементу из arr
	for _, num := range arr {
		count[num-minDigit]++
	}

	// Восстанавливаем отсортированный срез
	sortedArr := []int{}
	for i, val := range count {
		for j := 0; j < val; j++ {
			sortedArr = append(sortedArr, i+minDigit)
		}
	}

	return sortedArr
}
