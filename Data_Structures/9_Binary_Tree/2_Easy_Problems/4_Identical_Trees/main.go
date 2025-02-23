package main

import (
	binaryTree "algs/Data_Structures/9_Binary_Tree/1_Introduction"
	"fmt"
)

// При наличии двух бинарных деревьев задача состоит в том, чтобы определить, идентичны ли они оба или нет.
// Два дерева идентичны, если они содержат одинаковые данные и расположение данных также одинаковое.

// используя bfs
func isIdenticalBFS(r1, r2 *binaryTree.TreeNode) bool {
	if r1 == nil && r2 == nil {
		return true
	}

	if r1 == nil || r2 == nil {
		return false
	}

	// Используем два среза как очереди
	queue1 := []*binaryTree.TreeNode{r1}
	queue2 := []*binaryTree.TreeNode{r2}

	for len(queue1) > 0 && len(queue2) > 0 {
		node1 := queue1[0]
		node2 := queue2[0]

		queue1 = queue1[1:]
		queue2 = queue2[1:]

		// Проверяем значение текущих узлов
		if node1.Val != node2.Val {
			return false
		}

		// Проверяем левых потомков
		if (node1.Left != nil) != (node2.Left != nil) {
			return false
		}
		if node1.Left != nil {
			queue1 = append(queue1, node1.Left)
			queue2 = append(queue2, node2.Left)
		}

		// Проверяем правых потомков
		if (node1.Right != nil) != (node2.Right != nil) {
			return false
		}
		if node1.Right != nil {
			queue1 = append(queue1, node1.Right)
			queue2 = append(queue2, node2.Right)
		}
	}

	return len(queue1) == 0 && len(queue2) == 0
}

// Временная сложность: O(n), где n - количество узлов в большем из двух деревьев, поскольку каждый узел посещается один раз.
// Вспомогательное пространство: O(1), обход Морриса временно изменяет дерево, устанавливая “потоки”
// для обхода узлов без использования рекурсии или стека.

// Morris Traversal для проверки идентичности двух деревьев.
func isIdenticalMorris(r1, r2 *binaryTree.TreeNode) bool {
	for r1 != nil && r2 != nil {
		// Если значения узлов различаются, деревья не идентичны
		if r1.Val != r2.Val {
			return false
		}

		// Morris Traversal для первого дерева (r1)
		if r1.Left == nil {
			r1 = r1.Right
		} else {
			pre := r1.Left
			for pre.Right != nil && pre.Right != r1 {
				pre = pre.Right
			}

			if pre.Right == nil {
				pre.Right = r1
				r1 = r1.Left
			} else {
				pre.Right = nil
				r1 = r1.Right
			}
		}

		// Morris Traversal для второго дерева (r2)
		if r2.Left == nil {
			r2 = r2.Right
		} else {
			pre := r2.Left
			for pre.Right != nil && pre.Right != r2 {
				pre = pre.Right
			}

			if pre.Right == nil {
				pre.Right = r2
				r2 = r2.Left
			} else {
				pre.Right = nil
				r2 = r2.Right
			}
		}
	}

	// Оба указателя должны быть nil, иначе деревья разные
	return r1 == nil && r2 == nil
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

	fmt.Println(isIdenticalBFS(root1, root)) // true
	fmt.Println(isIdenticalMorris(root1, root))
}
