package main

import "fmt"

// Node представляет узел бинарного дерева поиска
type Node struct {
	Key   int
	Left  *Node
	Right *Node
}

// BinSearchTree представляет бинарное дерево поиска
type BinSearchTree struct {
	Root *Node
}

// insertNode рекурсивно добавляет новый узел в дерево
func insertNode(node *Node, key int) *Node {
	if node == nil {
		return &Node{Key: key}
	}

	if key < node.Key {
		node.Left = insertNode(node.Left, key)
	} else {
		node.Right = insertNode(node.Right, key)
	}
	return node
}

func (tree *BinSearchTree) Insert(key int) {
	tree.Root = insertNode(tree.Root, key)
}

// searchNode рекурсивно ищет узел в дереве
func searchNode(node *Node, key int) bool {
	if node == nil {
		return false
	}
	if key < node.Key {
		return searchNode(node.Left, key)
	} else if key > node.Key {
		return searchNode(node.Right, key)
	}
	return true
}

// Search ищет узел с заданным ключом в дереве
func (tree *BinSearchTree) Search(key int) bool {
	return searchNode(tree.Root, key)
}

// inOrderTraversal рекурсивно выполняет инфиксный обход дерева
func inOrderTraversal(node *Node) {
	if node != nil {
		inOrderTraversal(node.Left)
		fmt.Print(node.Key, " ")
		inOrderTraversal(node.Right)
	}
}

// InOrder печатает ключи узлов в инфиксном порядке
func (tree *BinSearchTree) InOrder() {
	inOrderTraversal(tree.Root)
	fmt.Println()
}
