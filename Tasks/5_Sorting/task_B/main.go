package main

import (
	"fmt"
	"sort"
)

type Gnom struct {
	a     int
	b     int
	index int
}

func main() {
	var n int
	fmt.Scan(&n)

	gnoms := make([]Gnom, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&gnoms[i].a)
		gnoms[i].index = i + 1
	}

	for i := 0; i < n; i++ {
		fmt.Scan(&gnoms[i].b)
	}

	sort.Slice(gnoms, func(i, j int) bool {
		return gnoms[i].b > gnoms[j].b
	})

	totalTime := 0
	ans := true
	for i := 0; i < n; i++ {
		if totalTime >= gnoms[i].b {
			ans = false
			break
		}
		totalTime += gnoms[i].a
	}

	if !(ans) {
		fmt.Println(-1)
		return
	}

	for i := 0; i < n; i++ {
		fmt.Print(gnoms[i].index, " ")
	}
	fmt.Println()
}
