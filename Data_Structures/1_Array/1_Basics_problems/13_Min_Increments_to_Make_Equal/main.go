package main

import (
	"fmt"
	"math/rand"
	"slices"
)

// Вам дан массив из n элементов, и вы должны найти количество операций,
// необходимых для того, чтобы сделать все элементы массива равными.
// При этом за одну операцию можно увеличить элемент на k.
// Если сделать все элементы равными невозможно, выведите -1.
func minOps(arr []int, k int) int {
	maxVal := slices.Max(arr)
	res := 0

	for _, x := range arr {
		if (maxVal-x)%k != 0 {
			return -1
		}
		res += (maxVal - x) / k
	}

	return res
}

func main() {
	arr := []int{}
	for i := 0; i < 10; i++ {
		arr = append(arr, rand.Intn(100))
	}

	fmt.Println(minOps(arr, 1))
}
