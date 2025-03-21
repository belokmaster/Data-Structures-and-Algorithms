package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
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

func (l *LinkedList) Add(value int) {
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

func (l *LinkedList) ToJSON() ([]byte, error) {
	l.mu.Lock()
	defer l.mu.Unlock()
	return json.Marshal(l.Head)
}

var list = LinkedList{}

func addHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	var value int
	if _, err := fmt.Sscanf(r.FormValue("value"), "%d", &value); err != nil {
		http.Error(w, "Invalid number", http.StatusBadRequest)
		return
	}

	list.Add(value)
	w.WriteHeader(http.StatusOK)
}

func addFrontHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Только POST", http.StatusMethodNotAllowed)
		return
	}

	value, err := strconv.Atoi(r.FormValue("value"))
	if err != nil {
		http.Error(w, "Неверное число", http.StatusBadRequest)
		return
	}

	list.mu.Lock()
	defer list.mu.Unlock()

	newNode := &Node{Value: value, Next: list.Head}
	list.Head = newNode

	w.WriteHeader(http.StatusOK)
}

func listHandler(w http.ResponseWriter, r *http.Request) {
	data, err := list.ToJSON()
	if err != nil {
		http.Error(w, "Error converting to JSON", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

func clearHandler(w http.ResponseWriter, r *http.Request) {
	list.mu.Lock()
	defer list.mu.Unlock()
	list.Head = nil
	w.WriteHeader(http.StatusOK)
}

func main() {
	http.HandleFunc("/add", addHandler)
	http.HandleFunc("/addFront", addFrontHandler)
	http.HandleFunc("/list", listHandler)
	http.HandleFunc("/clear", clearHandler) // Новый обработчик
	http.Handle("/", http.FileServer(http.Dir("static")))

	fmt.Println("Server running at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
