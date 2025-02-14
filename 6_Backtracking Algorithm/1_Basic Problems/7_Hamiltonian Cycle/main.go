package main

import (
	"fmt"
)

/*
Гамильтонов цикл — это замкнутый маршрут прохождения всех вершин выбранного графа строго по одному разу.
*/

/*
Гамильтонов цикл — это замкнутый маршрут прохождения всех вершин выбранного графа строго по одному разу.
Поиск гамильтонова цикла в графе - это хорошо известная NP-полная задача, которая означает,
что не существует известного эффективного алгоритма для ее решения для всех типов графов.
Однако она может быть решена для небольших или специфических типов графов.
*/

/*
Гамильтонов путь в графе G - это путь, который посещает каждую вершину G ровно один раз,
и гамильтонов путь не обязательно возвращается в начальную вершину. Это открытый путь.
*/

const V = 5

// isSafe проверяет, можно ли добавить вершину v в гамильтонов цикл
func isSafe(v int, graph [][]int, path []int, pos int) bool {
	// Проверяем, соединена ли текущая вершина с последней добавленной
	if graph[path[pos-1]][v] == 0 {
		return false
	}

	// Проверяем, не была ли эта вершина уже использована
	for i := 0; i < pos; i++ {
		if path[i] == v {
			return false
		}
	}
	return true
}

// hamCycleUtil рекурсивно ищет гамильтонов цикл
func hamCycleUtil(graph [][]int, path []int, pos int) bool {
	if pos == V {
		return graph[path[pos-1]][path[0]] == 1
	}

	for v := 1; v < V; v++ {
		if isSafe(v, graph, path, pos) {
			path[pos] = v
			if hamCycleUtil(graph, path, pos+1) {
				return true
			}
			// Откат (backtracking)
			path[pos] = -1
		}
	}
	return false
}

// hamCycle решает задачу поиска гамильтонова цикла
func hamCycle(graph [][]int) bool {
	path := make([]int, V)
	for i := range path {
		path[i] = -1
	}

	path[0] = 0
	if !hamCycleUtil(graph, path, 1) {
		fmt.Println("Решение не существует")
		return false
	}

	printSolution(path)
	return true
}

// printSolution выводит найденный гамильтонов цикл
func printSolution(path []int) {
	fmt.Println("Решение существует. Один из гамильтоновых циклов:")
	for _, v := range path {
		fmt.Print(v, " ")
	}
	fmt.Println(path[0]) // Замыкаем цикл
}

func main() {
	graph1 := [][]int{
		{0, 1, 0, 1, 0},
		{1, 0, 1, 1, 1},
		{0, 1, 0, 0, 1},
		{1, 1, 0, 0, 1},
		{0, 1, 1, 1, 0},
	}
	hamCycle(graph1)

	graph2 := [][]int{
		{0, 1, 0, 1, 0},
		{1, 0, 1, 1, 1},
		{0, 1, 0, 0, 1},
		{1, 1, 0, 0, 0},
		{0, 1, 1, 0, 0},
	}
	hamCycle(graph2)
}
