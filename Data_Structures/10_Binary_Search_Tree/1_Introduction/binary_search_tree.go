package binarySearchTree

//package main

import (
	"fmt"
	"math/rand"
	"time"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Функция вставки нового узла в бинарное дерево поиска
func Insert(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: key}
	}

	if key == root.Val {
		return root
	} else if key < root.Val {
		root.Left = Insert(root.Left, key)
	} else {
		root.Right = Insert(root.Right, key)
	}

	return root
}

func InsertIterative(root *TreeNode, key int) *TreeNode {
	temp := &TreeNode{Val: key}
	if root == nil {
		return temp
	}

	var prev *TreeNode
	parent := root

	for parent != nil {
		prev = parent
		if key < parent.Val {
			parent = parent.Left
		} else if key > parent.Val {
			parent = parent.Right
		} else {
			return root
		}
	}

	if key < prev.Val {
		prev.Left = temp
	} else {
		prev.Right = temp
	}

	return root
}

func Search(root *TreeNode, key int) bool {
	if root == nil {
		return false
	}

	if root.Val == key {
		return true
	}

	if root.Val < key {
		return Search(root.Right, key)
	}
	return Search(root.Left, key)
}

// Функция поиска минимального значения в дереве
func minValueNode(node *TreeNode) *TreeNode {
	current := node
	for current.Left != nil {
		current = current.Left
	}
	return current
}

// Функция удаления узла
func Delete(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return root
	}

	if key < root.Val {
		root.Left = Delete(root.Left, key)
	} else if key > root.Val {
		root.Right = Delete(root.Right, key)
	} else {
		// Узел с одним или без потомков
		if root.Left == nil {
			return root.Right
		} else if root.Right == nil {
			return root.Left
		}

		// Узел с двумя потомками: ищем минимальный узел в правом поддереве
		temp := minValueNode(root.Right)
		root.Val = temp.Val
		root.Right = Delete(root.Right, temp.Val)
	}
	return root
}

// Функция для красивого вывода дерева
func PrintTree(root *TreeNode, space int, level int) {
	if root == nil {
		return
	}

	space += 5
	PrintTree(root.Right, space, level+1)

	fmt.Print("\n")
	for i := 5; i < space; i++ {
		fmt.Print(" ")
	}
	fmt.Println(root.Val)

	PrintTree(root.Left, space, level+1)
}

func Inorder(root *TreeNode) {
	if root == nil {
		return
	}

	Inorder(root.Left)
	fmt.Print(root.Val, " ")
	Inorder(root.Right)
}

func PreOrder(root *TreeNode) {
	if root == nil {
		return
	}

	fmt.Print(root.Val, " ")
	PreOrder(root.Left)
	PreOrder(root.Right)
}

func PostOrder(root *TreeNode) {
	if root == nil {
		return
	}

	PostOrder(root.Left)
	PostOrder(root.Right)
	fmt.Print(root.Val, " ")
}

func MinValueBST(root *TreeNode) int {
	if root == nil {
		return -1
	}

	current := root
	for current.Left != nil {
		current = current.Left
	}

	return current.Val
}

func MaxValueBST(root *TreeNode) int {
	if root == nil {
		return -1
	}

	current := root
	for current.Right != nil {
		current = current.Right
	}

	return current.Val
}

// Функция для вставки значения в бинарное дерево
func InsertBT(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}

	if val < root.Val {
		root.Left = InsertBT(root.Left, val)
	} else {
		root.Right = InsertBT(root.Right, val)
	}

	return root
}

// Функция для создания случайного бинарного дерева из n элементов
func GenerateRandomBT(n int) *TreeNode {
	if n == 0 {
		return nil
	}

	rand.Seed(time.Now().UnixNano())
	root := &TreeNode{Val: rand.Intn(100)}
	for i := 1; i < n; i++ {
		InsertBT(root, rand.Intn(100))
	}

	return root
}

// Функция для генерации случайного бинарного дерева
func GenerateRandomTree(nodeCount int) *TreeNode {
	rand.Seed(time.Now().UnixNano()) // Инициализация генератора случайных чисел

	var root *TreeNode
	for i := 0; i < nodeCount; i++ {
		val := rand.Intn(100) // Генерация случайного значения от 0 до 99
		root = Insert(root, val)
	}
	return root
}

/*
func main() {
	root := &TreeNode{Val: 50}
	root = InsertIterative(root, 60)
	root = Insert(root, 70)
	root = Insert(root, 30)
	root = Insert(root, 20)
	root = Insert(root, 100)
	root = Insert(root, 1)
	root = Insert(root, 21)
	root = Delete(root, 50)
	PrintTree(root, 0, 0)
	Inorder(root)
	fmt.Println()
	PreOrder(root)
}
*/
