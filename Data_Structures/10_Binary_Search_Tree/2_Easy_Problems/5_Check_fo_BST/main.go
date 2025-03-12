package main

import (
	bts "algs/Data_Structures/10_Binary_Search_Tree/1_Introduction"
	"sort"
)

// При наличии двоичного дерева задача состоит в том, чтобы преобразовать его в двоичное дерево поиска.

/*
Идея заключается в рекурсивном обходе бинарного дерева и сохранении узлов в массиве.
Отсортировать массив, выполнить обход дерева по порядку и обновить значение каждого узла до соответствующего
значения в дереве.
*/

func inorder(root *bts.TreeNode, nodes *[]int) {
	if root == nil {
		return
	}

	inorder(root.Left, nodes)
	*nodes = append(*nodes, root.Val)
	inorder(root.Right, nodes)
}

// Функция для преобразования бинарного дерева в BST
func constructBST(root *bts.TreeNode, nodes []int, index *int) {
	if root == nil {
		return
	}

	constructBST(root.Left, nodes, index)
	root.Val = nodes[*index]
	*index++
	constructBST(root.Right, nodes, index)
}

func binaryTreeToBST(root *bts.TreeNode) *bts.TreeNode {
	var nodes []int
	inorder(root, &nodes)
	sort.Ints(nodes) // Сортируем значения узлов

	index := 0
	constructBST(root, nodes, &index)

	return root
}

func main() {
	root := bts.GenerateRandomBT(10)
	root = binaryTreeToBST(root)
	bts.Inorder(root)
}
