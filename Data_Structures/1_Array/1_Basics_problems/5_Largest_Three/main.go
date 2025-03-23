package main

import (
	"fmt"
	"math/rand"
	"sort"
)

// Дан массив положительных целых чисел arr[] размером n.
// Задача состоит в том, чтобы найти третий по величине неповторяющийся элемент в массиве.
func findThirdMax(arr []int) int {
	n := len(arr)
	if n == 1 {
		return arr[0]
	}

	firstMax := arr[0]
	secondMax := arr[1]
	if secondMax > firstMax {
		firstMax, secondMax = secondMax, firstMax
	}

	if n == 2 {
		return secondMax
	}

	thirdMax := arr[2]
	if thirdMax > firstMax {
		temp := firstMax

		firstMax = thirdMax
		thirdMax = secondMax
		secondMax = temp
	} else if thirdMax > secondMax {
		secondMax, thirdMax = thirdMax, secondMax
	}

	for i := 3; i < n; i++ {
		if arr[i] > firstMax {
			temp := firstMax

			firstMax = arr[i]
			thirdMax = secondMax
			secondMax = temp
		} else if arr[i] > secondMax {
			thirdMax = secondMax
			secondMax = arr[i]
		} else if arr[i] > thirdMax {
			thirdMax = arr[i]
		}
	}

	return thirdMax
}

func main() {
	arr := []int{}
	for i := 0; i < 10; i++ {
		arr = append(arr, rand.Intn(100))
	}

	ans := findThirdMax(arr)
	fmt.Println(arr)
	fmt.Println(ans)
	sort.Ints(arr)
	fmt.Println(arr)
}
