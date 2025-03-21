package main

import (
	"net/http"
	"strconv"
)

func addEndHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Только POST", http.StatusMethodNotAllowed)
		return
	}

	value, err := strconv.Atoi(r.FormValue("value"))
	if err != nil {
		http.Error(w, "Неверное число", http.StatusBadRequest)
		return
	}

	list.PushBack(value)
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

	list.PushBegin(value)
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
