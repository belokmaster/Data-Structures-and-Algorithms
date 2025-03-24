package main

import (
	"fmt"
	"math/rand"
)

// Дано массив prices[] размером n, обозначающий стоимость акций в каждый день.
// Задача состоит в том, чтобы найти максимальную общую прибыль, если мы можем покупать
// и продавать акции любое количество раз. Примечание: мы можем продавать только те акции,
// которые купили ранее, и не можем владеть несколькими акциями в один день.
func maximumProfit(arr []int) int {
	n := len(arr)
	lMin := arr[0]
	lMax := arr[0]
	res := 0

	i := 0
	for i < n-1 {
		for i < n-1 && arr[i] >= arr[i+1] {
			i++
		}
		lMin = arr[i]

		for i < n-1 && arr[i] <= arr[i+1] {
			i++
		}
		lMax = arr[i]

		res += lMax - lMin
	}

	return res
}

func maximumProfit2(arr []int) int {
	res := 0
	for i := 1; i < len(arr); i++ {
		if arr[i] > arr[i-1] {
			res += arr[i] - arr[i-1]
		}
	}
	return res
}

func main() {
	arr := []int{}
	for i := 0; i < 10; i++ {
		arr = append(arr, rand.Intn(100))
	}
	fmt.Println(arr)
	ans := maximumProfit(arr)
	ans2 := maximumProfit2(arr)
	fmt.Println(ans, ans2)
}
