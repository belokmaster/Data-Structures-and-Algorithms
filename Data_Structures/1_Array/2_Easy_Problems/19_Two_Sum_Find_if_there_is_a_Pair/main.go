package main

import (
	"fmt"
	"math/rand"
	"sort"
)

// Даны массив arr[] из n целых чисел и целевое значение. Задача состоит в том,
// чтобы определить, есть ли в массиве пара элементов, сумма которых равна целевому значению.
func binarySearch(arr []int, left, right, target int) bool {
	for left <= right {
		mid := left + (right-left)/2
		if arr[mid] == target {
			return true
		}
		if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return false
}

func twoSum(arr []int, target int) bool {
	sort.Ints(arr)
	for i := 0; i < len(arr); i++ {
		compl := target - arr[i]
		if binarySearch(arr, i+1, len(arr)-1, compl) {
			return true
		}
	}
	return false
}

func main() {
	arr := []int{}
	for i := 0; i < 10; i++ {
		arr = append(arr, rand.Intn(200)-100)
	}
	fmt.Println(arr)
	ans := twoSum(arr, 2)
	fmt.Println(ans)
}
