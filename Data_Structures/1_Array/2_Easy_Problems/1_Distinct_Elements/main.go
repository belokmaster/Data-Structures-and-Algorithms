package main

import (
	"fmt"
	"math/rand"
)

// Дан целочисленный массив arr[], выведите все уникальные элементы этого массива
func findDistinct(arr []int) []int {
	distincts := []int{}
	hashMap := make(map[int]int)
	for _, num := range arr {
		hashMap[num]++
	}

	for key, val := range hashMap {
		if val == 1 {
			distincts = append(distincts, key)
		}
	}

	return distincts
}

func main() {
	arr := []int{}
	for i := 0; i < 10; i++ {
		arr = append(arr, rand.Intn(100))
	}
	fmt.Println(arr)
	ans := findDistinct(arr)
	fmt.Println(ans)
}
