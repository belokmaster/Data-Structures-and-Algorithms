package main

import (
	binaryTree "algs/Data_Structures/9_Binary_Tree/1_Introduction"
	"fmt"
	"math"
)

// найти макс элемент в бинарном дереве.

// Временная сложность: O(n), где n - количество узлов в бинарном дереве.
// Вспомогательное пространство: O(n).
// Выполняется рекурсивный вызов. Каждый узел обрабатывается один раз, и,
// учитывая объем стека, сложность пространства будет равна O(N).
func findMin(root *binaryTree.TreeNode) int {
	if root == nil {
		return math.MaxInt
	}

	res := root.Val
	left := findMin(root.Left)
	right := findMin(root.Right)

	if left < res {
		res = left
	}
	if right < res {
		res = right
	}

	return res
}

// итеративно через BFS
func findMinIterative(root *binaryTree.TreeNode) int {
	if root == nil {
		return math.MaxInt
	}

	queue := []*binaryTree.TreeNode{root}
	res := math.MaxInt

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if node.Val < res {
			res = node.Val
		}

		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}

	return res
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
	fmt.Println(findMin(root))
	fmt.Println(findMinIterative(root))
}
