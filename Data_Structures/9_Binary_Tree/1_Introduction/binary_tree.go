package binaryTree

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func NewNode(value int) *TreeNode {
	return &TreeNode{Val: value, Left: nil, Right: nil}
}

// (Обход в глубину: Левый → Корень → Правый)
func InOrderDFS(node *TreeNode) {
	if node == nil {
		return
	}

	InOrderDFS(node.Left)
	fmt.Print(node.Val, " ")
	InOrderDFS(node.Right)
}

// (Обход в глубину: Корень → Левый → Правый)
func PreOrderDFS(node *TreeNode) {
	if node == nil {
		return
	}
	fmt.Print(node.Val, " ")
	PreOrderDFS(node.Left)
	PreOrderDFS(node.Right)
}

// (Обход в глубину: Левый → Правый → Корень)
func PostOrderDFS(node *TreeNode) {
	if node == nil {
		return
	}

	PostOrderDFS(node.Left)
	PostOrderDFS(node.Right)
	fmt.Print(node.Val, " ")
}

// Обход в ширину для проверки структуры дерева
func LevelOrder(root *TreeNode) {
	if root == nil {
		return
	}

	queue := []*TreeNode{root}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		fmt.Print(node.Val, " ")

		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}

	fmt.Println()
}

// Функция вставки нового узла в бинарное дерево
func Insert(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: key}
	}

	queue := []*TreeNode{root}

	// Выполняем обход уровней, пока не найдем свободное место
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		// Если левый потомок пуст, вставляем новый узел сюда
		if node.Left == nil {
			node.Left = &TreeNode{Val: key}
			break
		} else {
			queue = append(queue, node.Left)
		}

		// Если правый потомок пуст, вставляем новый узел сюда
		if node.Right == nil {
			node.Right = &TreeNode{Val: key}
			break
		} else {
			queue = append(queue, node.Right)
		}
	}

	return root
}

func Search(root *TreeNode, target int) bool {
	if root == nil {
		return false
	}

	queue := []*TreeNode{root}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if node.Val == target {
			return true
		}

		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}

	return false
}

// Функция для нахождения минимального узла в дереве
func findMin(root *TreeNode) *TreeNode {
	current := root
	for current.Left != nil {
		current = current.Left
	}
	return current
}

// Функция удаления узла из бинарного дерева
func DeleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}

	// Если значение узла меньше ключа, ищем в правом поддереве
	if key < root.Val {
		root.Right = DeleteNode(root.Right, key)
	} else if key > root.Val { // Если больше, ищем в левом поддереве
		root.Left = DeleteNode(root.Left, key)
	} else {
		// Если нашли узел для удаления

		// Если у узла нет детей
		if root.Left == nil && root.Right == nil {
			return nil
		}

		if root.Left == nil {
			return root.Right
		}

		if root.Right == nil {
			return root.Left
		}

		// Если у узла есть два ребенка
		// Найдем минимальный элемент в правом поддереве (или максимальный в левом)
		minRight := findMin(root.Right)
		root.Val = minRight.Val
		// Удалим минимальный узел из правого поддерева
		root.Right = DeleteNode(root.Right, minRight.Val)
	}

	return root
}

func MaxDeapth(root *TreeNode) int {
	level := 0
	if root == nil {
		return level
	}

	queue := []*TreeNode{root}
	for len(queue) > 0 {
		size := len(queue)

		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		level++
	}

	return level
}
