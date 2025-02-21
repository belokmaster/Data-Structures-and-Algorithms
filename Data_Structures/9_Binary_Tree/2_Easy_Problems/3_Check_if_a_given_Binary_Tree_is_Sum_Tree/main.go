package main

import (
	binaryTree "algs/Data_Structures/9_Binary_Tree/1_Introduction"
	"fmt"
)

// При наличии двоичного дерева задача состоит в том, чтобы проверить, является ли оно суммарным деревом.
// Суммарное дерево - это двоичное дерево, в котором значение узла равно сумме узлов,
// присутствующих в его левом поддереве и правом поддереве.
// Пустое дерево является суммарным деревом, и сумма по пустому дереву может быть равна 0.
// Конечный узел также считается суммарным деревом.

// Временная сложность: O(n^2), где n - количество узлов в бинарном дереве.
// Вспомогательное пространство: O(h)

// Функция для проверки, является ли дерево суммарным деревом
func isSumTree(root *binaryTree.TreeNode) bool {
	if root == nil {
		return true
	}

	queue := []*binaryTree.TreeNode{root}
	// Массив для хранения результатов проверки для каждого узла
	sumResults := make(map[*binaryTree.TreeNode]int)

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}

		// Рассчитываем сумму для текущего узла (левая + правая части)
		leftSum := 0
		rightSum := 0
		if node.Left != nil {
			leftSum = sumResults[node.Left]
		}
		if node.Right != nil {
			rightSum = sumResults[node.Right]
		}

		// Если узел не лист и его значение не равно сумме его детей, то возвращаем false
		if (node.Left != nil || node.Right != nil) && node.Val != leftSum+rightSum {
			return false
		}

		// Сохраняем сумму текущего узла (оно равно сумме значений его детей + его собственное значение)
		sumResults[node] = node.Val + leftSum + rightSum
	}

	return true
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
	fmt.Println(isSumTree(root))
}
