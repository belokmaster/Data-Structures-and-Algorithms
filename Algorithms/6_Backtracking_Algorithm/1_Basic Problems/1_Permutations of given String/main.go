package main

import (
	"fmt"
	"sort"
	"strings"
)

// Для заданной строки s задача состоит в том, чтобы вернуть все перестановки данной строки в лексикографически отсортированном порядке.
// Примечание: Перестановка — это перестановка всех элементов строки. Возможно дублирование перестановок.

// Time Complexity: O(n * n!)
// Auxiliary Space: O(n!)
func recurPermutation(index int, s []rune, ans *[]string) {
	if index == len(s) {
		*ans = append(*ans, string(s))
		return
	}

	for i := index; i < len(s); i++ {
		s[index], s[i] = s[i], s[index]
		recurPermutation(index+1, s, ans)
		s[index], s[i] = s[i], s[index]
	}
}

func findPermutation(s string) []string {
	ans := []string{}
	recurPermutation(0, []rune(s), &ans)
	sort.Strings(ans)
	return ans
}

func main() {
	var s string
	fmt.Scan(&s)

	res := findPermutation(s)
	fmt.Println(strings.Join(res, " "))
}
