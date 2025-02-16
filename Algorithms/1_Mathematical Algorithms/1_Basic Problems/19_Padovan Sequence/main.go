package main

import "fmt"

// Последовательность Падована похожа на последовательность Фибоначчи с аналогичной рекурсивной структурой.
// Рекурсивная формула равна:   P(n) = P(n-2) + P(n-3)

func pad(n int) int {
	pPrevPrev := 1
	pPrev := 1
	pNext := 1

	pCurr := 1

	for i := 3; i < n; i++ {
		pNext = pPrev + pPrevPrev
		pPrevPrev = pPrev
		pPrev = pCurr
		pCurr = pNext
	}

	return pNext
}

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Println(pad(n))
}
