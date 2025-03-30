package main

import (
	"fmt"
	"math"
	"sort"
)

// Дан целочисленный массив из N элементов. Необходимо найти максимальную сумму двух элементов,
// которая будет максимально близка к нулю.
// Если у нас есть два или более способов вычислить сумму двух элементов,
// ближайших к нулю, верните максимальную сумму.
func closestPairSum(arr []int) int {
	sort.Ints(arr)
	n := len(arr)
	if n < 2 {
		return 0
	}

	left := 0
	right := n - 1
	res := arr[left] + arr[right]

	for i := 0; i < n; i++ {
		for left <= right {
			current := arr[left] + arr[right]
			if current == 0 {
				return 0
			}

			if int((math.Abs(float64(current)))) < int(math.Abs(float64(res))) {
				res = current
			}

			if current < 0 {
				left++
			} else {
				right--
			}
		}
	}

	return res
}

func main() {
	arr := []int{-10, -3, 0, 5, 9, 20}
	fmt.Println(arr)
	ans := closestPairSum(arr)
	fmt.Println(ans)
}
