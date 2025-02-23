package main

import (
	binaryTree "algs/Data_Structures/9_Binary_Tree/1_Introduction"
	"fmt"
)

// Дано два бинарных дерева , задача состоит в том, чтобы проверить, являются ли два дерева зеркальным отражением
// друг друга или нет.
// Функция для проверки зеркальности двух деревьев
func areMirrors(root1, root2 *binaryTree.TreeNode) bool {
	// Если оба корня пусты, они зеркальны
	if root1 == nil && root2 == nil {
		return true
	}

	// Если только один корень пуст, деревья не зеркальны
	if root1 == nil || root2 == nil {
		return false
	}

	// Создаем два стека для одновременного обхода
	stk1 := []*binaryTree.TreeNode{root1}
	stk2 := []*binaryTree.TreeNode{root2}

	// Пока оба стека не пусты
	for len(stk1) > 0 && len(stk2) > 0 {
		// Извлекаем узлы из обоих стеков
		curr1 := stk1[len(stk1)-1]
		stk1 = stk1[:len(stk1)-1]
		curr2 := stk2[len(stk2)-1]
		stk2 = stk2[:len(stk2)-1]

		// Проверяем, одинаковы ли данные узлов
		if curr1.Val != curr2.Val {
			return false
		}

		// Проверяем зеркальность поддеревьев
		if curr1.Left != nil && curr2.Right != nil {
			stk1 = append(stk1, curr1.Left)
			stk2 = append(stk2, curr2.Right)
		} else if curr1.Left != nil || curr2.Right != nil {
			return false
		}

		if curr1.Right != nil && curr2.Left != nil {
			stk1 = append(stk1, curr1.Right)
			stk2 = append(stk2, curr2.Left)
		} else if curr1.Right != nil || curr2.Left != nil {
			return false
		}
	}

	// Если оба стека пусты, деревья зеркальны
	return len(stk1) == 0 && len(stk2) == 0
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

	root1 := &binaryTree.TreeNode{Val: 1}
	root1 = binaryTree.Insert(root1, 2)
	root1 = binaryTree.Insert(root1, 3)
	root1 = binaryTree.Insert(root1, 4)
	root1 = binaryTree.Insert(root1, 5)
	root1 = binaryTree.Insert(root1, 8)
	root1 = binaryTree.Insert(root1, 6)
	root1 = binaryTree.Insert(root1, 7)

	fmt.Println(areMirrors(root1, root))
}
