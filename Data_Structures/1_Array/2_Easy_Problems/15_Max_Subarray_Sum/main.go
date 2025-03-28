package main

import (
	"fmt"
	"math/rand"
)

func sum(arr []int) int {
	ans := 0
	for _, num := range arr {
		ans += num
	}
	return ans
}

// Дан массив arr[], задача состоит в том, чтобы найти подмассив с максимальной суммой и вернуть её.
func maxSubarraySum(arr []int) ([]int, int) {
	maxSum := arr[0]
	currSum := 0
	start, end, tempStart := 0, 0, 0

	for i, num := range arr {
		if currSum <= 0 {
			currSum = num
			tempStart = i
		} else {
			currSum += num
		}

		if currSum > maxSum {
			maxSum = currSum
			start = tempStart
			end = i
		}
	}

	return arr[start : end+1], maxSum
}

func kadanAlg(arr []int) int {
	if len(arr) == 0 {
		return 0
	}

	res := arr[0]       // Глобальный максимум
	maxEnding := arr[0] // Локальная сумма

	for i := 1; i < len(arr); i++ {
		maxEnding = max(maxEnding+arr[i], arr[i]) // Продолжаем или начинаем новый подмассив
		res = max(res, maxEnding)                 // Обновляем глобальный максимум
	}

	return res
}

func main() {
	arr := []int{}
	for i := 0; i < 10; i++ {
		arr = append(arr, rand.Intn(200)-100)
	}
	fmt.Println(arr)
	ansKadan := kadanAlg(arr)
	fmt.Println("Kadan' sum is", ansKadan)
	ans, sumAns := maxSubarraySum(arr)
	fmt.Println(ans, sumAns)

}
