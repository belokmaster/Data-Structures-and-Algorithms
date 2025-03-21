package main

import (
	"net/http"
	"sync"
)

type Node struct {
	Value int   `json:"value"`
	Next  *Node `json:"next,omitempty"`
}

type LinkedList struct {
	mu   sync.Mutex
	Head *Node
}

func (l *LinkedList) PushBack(value int) {
	l.mu.Lock()
	defer l.mu.Unlock()

	newNode := &Node{Value: value}
	if l.Head == nil {
		l.Head = newNode
	} else {
		current := l.Head
		for current.Next != nil {
			current = current.Next
		}
		current.Next = newNode
	}
}

func (l *LinkedList) PushBegin(value int) {
	l.mu.Lock()
	defer l.mu.Unlock()

	newNode := &Node{Value: value, Next: list.Head}
	list.Head = newNode
}

func clearHandler(w http.ResponseWriter, r *http.Request) {
	list.mu.Lock()
	defer list.mu.Unlock()

	list.Head = nil
	w.WriteHeader(http.StatusOK)
}
