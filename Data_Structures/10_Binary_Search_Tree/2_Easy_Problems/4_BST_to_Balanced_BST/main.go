package main

import (
	bts "algs/Data_Structures/10_Binary_Search_Tree/1_Introduction"
	"fmt"
)

// Дано бинарное дерево поиска, которое может быть несбалансированными, задача состоит в том,
// чтобы преобразовать его в сбалансированное BTS, что имеет минимально возможную высоту.

func inorderTraversal(root *bts.TreeNode, nodes *[]int) {
	if root == nil {
		return
	}

	inorderTraversal(root.Left, nodes)
	*nodes = append(*nodes, root.Val)
	inorderTraversal(root.Right, nodes)
}

func buildBalancedBTS(nodes []int, start, end int) *bts.TreeNode {
	if start > end {
		return nil
	}

	mid := (start + end) / 2
	root := &bts.TreeNode{Val: nodes[mid]}
	root.Left = buildBalancedBTS(nodes, start, mid-1)
	root.Right = buildBalancedBTS(nodes, mid+1, end)

	return root
}

func balanceBST(root *bts.TreeNode) *bts.TreeNode {
	nodes := []int{}
	inorderTraversal(root, &nodes)
	return buildBalancedBTS(nodes, 0, len(nodes)-1)
}

func main() {
	root := bts.GenerateRandomTree(10)
	bts.PreOrder(root)
	fmt.Println()
	root = balanceBST(root)
	bts.PreOrder(root)
}
