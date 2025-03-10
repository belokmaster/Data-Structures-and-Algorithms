package main

import (
	bts "algs/Data_Structures/10_Binary_Search_Tree/1_Introduction"
	"fmt"
)

// Найти второй по величине элемент

func findSecondLargest(root *bts.TreeNode) int {
	if root == nil || (root.Left == nil && root.Right == nil) {
		return -1 // Дерево пустое или содержит только один узел
	}

	current := root
	var prev *bts.TreeNode

	// Идем до самого правого узла (максимального значения)
	for current.Right != nil {
		prev = current
		current = current.Right
	}

	// Если у максимального узла есть левое поддерево, второе по величине значение — максимальное значение в этом поддереве
	if current.Left != nil {
		current = current.Left
		for current.Right != nil {
			current = current.Right
		}
		return current.Val
	}

	// Если левого поддерева нет, возвращаем значение родителя
	if prev != nil {
		return prev.Val
	}

	return -1
}

func main() {
	root := bts.GenerateRandomTree(10)
	fmt.Println(findSecondLargest(root))
	bts.PreOrder(root)
}
