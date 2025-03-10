package main

import (
	bts "algs/Data_Structures/10_Binary_Search_Tree/1_Introduction"
	"fmt"
)

// Даны два значения n1 и n2, где n1 < n2, и корневой указатель на дерево бинарного поиска.
// Задача состоит в том, чтобы найти все ключи дерева в диапазоне от n1 до n2 в порядке возрастания.

func printNearNodes(root *bts.TreeNode, low, high int) []int {
	ans := []int{}
	current := root

	for current != nil {
		if current.Left == nil {
			// Если левого поддерева нет, проверяем текущий узел
			if current.Val >= low && current.Val <= high {
				ans = append(ans, current.Val)
			}
			current = current.Right
		} else {
			// Находим inorder-предшественника текущего узла
			prev := current.Left
			for prev.Right != nil && prev.Right != current {
				prev = prev.Right
			}

			if prev.Right == nil {
				// Создаем связь с текущим узлом и переходим в левое поддерево
				prev.Right = current
				current = current.Left
			} else {
				// Восстанавливаем структуру дерева и проверяем текущий узел
				prev.Right = nil
				if current.Val >= low && current.Val <= high {
					ans = append(ans, current.Val)
				}
				current = current.Right
			}
		}
	}

	return ans
}

func main() {
	root := bts.GenerateRandomTree(10)
	fmt.Println(printNearNodes(root, 7, 80))
	bts.PreOrder(root)
}
