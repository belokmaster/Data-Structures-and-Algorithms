package main

import "fmt"

func main() {
	tree := &BinSearchTree{}
	keys := []int{50, 30, 70, 20, 40, 60, 80}

	for _, key := range keys {
		tree.Insert(key)
	}
	tree.Insert(100)
	tree.Insert(45)
	fmt.Println(tree.Root)
	fmt.Println("InOrder traversal:")
	tree.InOrder()

	fmt.Println("Search 40:", tree.Search(40))
	fmt.Println("Search 25:", tree.Search(25))
}
