package main

import "fmt"

func main() {
	heap := &MinHeap{}
	heap.Insert(1)
	heap.Insert(100)
	heap.Insert(50)
	heap.Insert(30)
	heap.Insert(20)
	fmt.Println(heap)
}
