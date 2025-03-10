package main

import (
	bts "algs/Data_Structures/10_Binary_Search_Tree/1_Introduction"
	"fmt"
)

// Дано бинарное дерево поиска. Задача состоит в том, чтобы найти сумму всех элементов,
// меньшую и равную k-му наименьшему элементу.

func sum(root *bts.TreeNode, k int) int {
	ans := 0
	calculateSum(root, &k, &ans)
	return ans
}

func calculateSum(root *bts.TreeNode, k *int, ans *int) {
	if root == nil || *k == 0 {
		return
	}

	// Рекурсивно обходим левое поддерево
	if root.Left != nil {
		calculateSum(root.Left, k, ans)
	}

	if *k > 0 {
		*ans += root.Val
		*k--
	} else {
		return
	}

	// Рекурсивно обходим правое поддерево
	if root.Right != nil {
		calculateSum(root.Right, k, ans)
	}
}

func main() {
	root := bts.GenerateRandomTree(10)
	fmt.Println(sum(root, 3))
	bts.PreOrder(root)
}
