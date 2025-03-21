package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var list = LinkedList{}

func (l *LinkedList) ToJSON() ([]byte, error) {
	l.mu.Lock()
	defer l.mu.Unlock()
	return json.Marshal(l.Head)
}

func main() {
	http.HandleFunc("/add", addEndHandler)
	http.HandleFunc("/addFront", addFrontHandler)
	http.HandleFunc("/list", listHandler)
	http.HandleFunc("/clear", clearHandler)
	http.Handle("/", http.FileServer(http.Dir("static")))

	fmt.Println("Server running at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
