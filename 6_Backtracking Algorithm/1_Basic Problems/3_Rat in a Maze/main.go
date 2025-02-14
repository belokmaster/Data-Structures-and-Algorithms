package main

import "fmt"

/*
Рассмотрим крысу, помещенную в (0, 0) квадратную матрицу порядка N * N.
Она должна достичь пункта назначения в (N – 1, N – 1). Найдите все возможные пути, по которым крыса
может добраться от источника до пункта назначения. Крыса может двигаться в следующих направлениях:
"U" (вверх), "D" (вниз), "L" (влево), "R" (вправо). Значение 0 в ячейке матрицы означает,
что она заблокирована и rat не может перейти в нее, в то время как значение 1 в ячейке матрицы означает,
что rat может пройти через нее. Возвращает список путей в лексикографически возрастающем порядке.
Примечание: В пути ни одна ячейка не может быть посещена более одного раза.
Если исходная ячейка равна 0, крыса не может перейти в какую-либо другую ячейку.
*/

// Time Complexity: O(3^(m*n)), потому что на каждой ячейке мы должны попробовать 3 разных направления, так как мы не будем проверять ячейку, из которой мы перешли в последний ход.
// Auxiliary Space: O(m*n), Максимальная глубина дерева рекурсии (вспомогательное пространство).

// все возможные пути
var direction = "DLRU"
var dr = []int{1, 0, 0, -1}
var dc = []int{0, -1, 1, 0}

// Проверка, можно ли попасть в клетку
func isValid(row, col, n int, maze [][]int) bool {
	return row >= 0 && col >= 0 && row < n && col < n && maze[row][col] == 1
}

// Поиск всех возможных путей
func findPath(row, col int, maze [][]int, n int, ans *[]string, currentPath *string) {
	if row == n-1 && col == n-1 {
		*ans = append(*ans, *currentPath)
		return
	}

	// Отмечаем клетку как посещённую
	maze[row][col] = 0

	for i := 0; i < 4; i++ {
		newRow, newCol := row+dr[i], col+dc[i]
		if isValid(newRow, newCol, n, maze) {
			*currentPath += string(direction[i])
			findPath(newRow, newCol, maze, n, ans, currentPath)
			*currentPath = (*currentPath)[:len(*currentPath)-1] // Откат при возврате
		}
	}

	// Освобождаем клетку
	maze[row][col] = 1
}

func main() {
	maze := [][]int{
		{1, 0, 0, 0},
		{1, 1, 1, 1},
		{1, 1, 1, 0},
		{0, 1, 1, 1},
	}

	n := len(maze)
	var result []string
	currentPath := ""

	if maze[0][0] != 0 && maze[n-1][n-1] != 0 {
		findPath(0, 0, maze, n, &result, &currentPath)
	}

	if len(result) == 0 {
		fmt.Println(-1)
	} else {
		for _, path := range result {
			fmt.Print(path, " ")
		}
		fmt.Println()
	}
}
