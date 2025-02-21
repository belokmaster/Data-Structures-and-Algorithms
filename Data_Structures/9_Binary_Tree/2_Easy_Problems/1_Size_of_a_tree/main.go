package main

import (
	binaryTree "algs/Data_Structures/9_Binary_Tree/1_Introduction"
	"fmt"
)

// Задача состоит в том, чтобы определить размер бинарного дерева.
// Размер дерева - это количество узлов, присутствующих в дереве.

// Временная сложность: O(n), где n - количество узлов в бинарном дереве.
// Вспомогательное пространство: O(h), где h - высота дерева.
func countNodes(root *binaryTree.TreeNode) int {
	if root == nil {
		return 0
	}

	left := countNodes(root.Left)
	right := countNodes(root.Right)

	return left + right + 1
}

func main() {
	root := &binaryTree.TreeNode{Val: 1}
	root = binaryTree.Insert(root, 2)
	root = binaryTree.Insert(root, 3)
	root = binaryTree.Insert(root, 4)
	root = binaryTree.Insert(root, 5)
	root = binaryTree.Insert(root, 8)
	root = binaryTree.Insert(root, 6)
	root = binaryTree.Insert(root, 7)
	fmt.Println(countNodes(root))
}
