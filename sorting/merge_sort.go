package main

func merge(arr1, arr2 []int) []int {
	var merged []int
	i, j := 0, 0

	// Слияние двух отсортированных массивов
	for i < len(arr1) && j < len(arr2) {
		if arr1[i] <= arr2[j] {
			merged = append(merged, arr1[i])
			i++
		} else {
			merged = append(merged, arr2[j])
			j++
		}
	}
	// Добавляем оставшиеся элементы
	for i < len(arr1) {
		merged = append(merged, arr1[i])
		i++
	}
	for j < len(arr1) {
		merged = append(merged, arr2[j])
		j++
	}

	return merged
}

func MergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	mid := len(arr) / 2
	left := MergeSort(arr[:mid])
	right := MergeSort(arr[mid:])

	return merge(left, right)
}
