package main

import (
	bts "algs/Data_Structures/10_Binary_Search_Tree/1_Introduction"
)

func inorderTraversal(root *bts.TreeNode, inorderArr *[]int) {
	if root == nil {
		return
	}

	inorderTraversal(root.Left, inorderArr)
	*inorderArr = append(*inorderArr, root.Val)
	inorderTraversal(root.Right, inorderArr)
}

func preorderFill(root *bts.TreeNode, inorderArr []int, index *int) {
	if root == nil {
		return
	}

	root.Val = inorderArr[*index]
	*index++

	preorderFill(root.Left, inorderArr, index)
	preorderFill(root.Right, inorderArr, index)
}

// convertBSTtoMinHeap конвертирует BST в Min Heap
func convertBSTtoMinHeap(root *bts.TreeNode) {
	var inorderArr []int

	// Шаг 1: Собираем значения BST в отсортированный массив
	inorderTraversal(root, &inorderArr)

	// Шаг 2: Используем Preorder Traversal для заполнения дерева
	index := 0
	preorderFill(root, inorderArr, &index)
}

func main() {
	root := bts.GenerateRandomTree(10)
	convertBSTtoMinHeap(root)
	bts.Inorder(root)
}
