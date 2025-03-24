package main

import (
	"fmt"
	"math/rand"
	"sort"
)

func rearrangeArray(arr []int) []int {
	n := len(arr)
	sort.Ints(arr)

	ans := make([]int, n)
	ptr1, ptr2 := 0, n-1

	for i := 0; i < n; i++ {
		if (i+1)%2 == 0 {
			ans[i] = arr[ptr2]
			ptr2--
		} else {
			ans[i] = arr[ptr1]
			ptr1++
		}
	}

	return ans
}

func main() {
	arr := []int{}
	for i := 0; i < 10; i++ {
		arr = append(arr, rand.Intn(100))
	}
	fmt.Println(arr)
	ans := rearrangeArray(arr)
	fmt.Println(ans)
}
