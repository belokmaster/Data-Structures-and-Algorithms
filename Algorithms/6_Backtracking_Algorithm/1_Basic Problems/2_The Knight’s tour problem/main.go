package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Дана N* N доска, на которой конь размещен на первом блоке пустой доски.
// Двигаясь в соответствии с правилами шахмат, конь должен посетить каждую клетку ровно один раз.
// Выведите порядок расположения каждой ячейки, в которой они были посещены.

/*
Обратный поиск работает поэтапно для решения проблем. Как правило, мы начинаем с пустого вектора решений и
один за другим добавляем элементы.

Когда мы добавляем элемент, мы проверяем, не нарушает ли добавление текущего элемента ограничение проблемы,
если это так, то мы удаляем элемент и пробуем другие альтернативы. Если ни одна из альтернатив не работает,
мы переходим к предыдущему этапу и удаляем элемент, добавленный на предыдущем этапе.
Если мы возвращаемся к начальной стадии, то говорим, что решения не существует.

Если добавление элемента не нарушает ограничений, мы рекурсивно добавляем элементы один за другим.
Если вектор решения становится полным, мы печатаем решение.
*/

// Time Complexity: O(8^n^2)
// Auxiliary Space: O(n^2)

const N = 8

// Проверка, является ли ход допустимым
func isSafe(x, y int, board [N][N]int) bool {
	return x >= 0 && x < N && y >= 0 && y < N && board[x][y] == -1
}

// Рекурсивная функция поиска пути
func solveKTUtil(x, y, movei int, board *[N][N]int, xMove, yMove [8]int) bool {
	if movei == N*N {
		return true
	}

	for k := 0; k < 8; k++ {
		nextX := x + xMove[k]
		nextY := y + yMove[k]
		if isSafe(nextX, nextY, *board) {
			board[nextX][nextY] = movei
			if solveKTUtil(nextX, nextY, movei+1, board, xMove, yMove) {
				return true
			}
			board[nextX][nextY] = -1 // Откат хода
		}
	}
	return false
}

// Вывод результата
func printSolution(board [N][N]int) {
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			fmt.Printf("%2d ", board[i][j])
		}
		fmt.Println()
	}
}

// Основная функция решения задачи
func solveKT() {
	var board [N][N]int
	for i := range board {
		for j := range board[i] {
			board[i][j] = -1
		}
	}

	xMove := [8]int{2, 1, -1, -2, -2, -1, 1, 2}
	yMove := [8]int{1, 2, 2, 1, -1, -2, -2, -1}

	// Начальная позиция коня
	board[0][0] = 0

	if solveKTUtil(0, 0, 1, &board, xMove, yMove) {
		printSolution(board)
	} else {
		fmt.Println("Решение не существует")
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	solveKT()
}
