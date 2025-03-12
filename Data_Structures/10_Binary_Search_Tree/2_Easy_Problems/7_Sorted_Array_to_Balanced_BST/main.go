package main

import (
	bts "algs/Data_Structures/10_Binary_Search_Tree/1_Introduction"
	"math/rand"
	"time"
)

// Сортированный массив в BST

func sortedArrayToBST(nums []int) *bts.TreeNode {
	if len(nums) == 0 {
		return nil
	}

	mid := len(nums) / 2
	root := &bts.TreeNode{Val: nums[mid]}

	root.Left = sortedArrayToBST(nums[:mid])
	root.Right = sortedArrayToBST(nums[mid+1:])

	return root
}

func generateRandomArray(n, min, max int) []int {
	rand.Seed(time.Now().UnixNano()) // Инициализируем генератор случайных чисел
	arr := make([]int, n)

	for i := 0; i < n; i++ {
		arr[i] = rand.Intn(max-min+1) + min // Генерируем число в диапазоне [min, max]
	}

	return arr
}

func main() {
	arr := generateRandomArray(10, 1, 100)
	root := sortedArrayToBST(arr)
	bts.Inorder(root)
}
