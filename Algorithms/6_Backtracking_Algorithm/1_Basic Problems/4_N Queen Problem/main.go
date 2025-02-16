package main

import "fmt"

/*
Учитывая целое число n, задача состоит в том, чтобы найти решение задачи о n ферзях,
где n ферзей расположены на n * n шахматной доске таким образом, что никакие два ферзя
не могут атаковать друг друга. N ферзей - это задача разместить N шахматных ферзей на
шахматной доске N × N таким образом, чтобы ни один из двух ферзей не атаковал друг друга.
*/

// Максимальное количество ферзей, которое может быть расставлено на шахматной доске размером n × n, равно n.

// Time Complexity: O(n*n!)
// размещение это по стандарту n!, а - n это проверка isSafe
// Auxiliary Space: O(n^2). Массив для хранения доски - n * n => n^2

// Функция isSafe проверяет три направления (вертикаль и две диагонали),
// чтобы убедиться, что ферзя можно безопасно поставить в board[row][col].
func isSafe(board [][]int, row, col int) bool {
	n := len(board)

	// Проверяем этот столбец сверху вниз
	for i := 0; i < row; i++ {
		if board[i][col] == 1 {
			return false
		}
	}

	// Проверяем верхнюю левую диагональ
	for i, j := row-1, col-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if board[i][j] == 1 {
			return false
		}
	}

	// Проверяем верхнюю правую диагональ
	for i, j := row-1, col+1; i >= 0 && j < n; i, j = i-1, j+1 {
		if board[i][j] == 1 {
			return false
		}
	}

	// Если все проверки пройдены, значит можно ставить ферзя
	return true
}

// placeQueens рекурсивно размещает ферзей на доске
func placeQueens(board [][]int, row int) bool {
	n := len(board)

	// Базовый случай: если расставлены все n ферзей
	if row == n {
		return true
	}

	// Пробуем поставить ферзя в текущую строку в каждую колонку
	for col := 0; col < n; col++ {
		if isSafe(board, row, col) {
			board[row][col] = 1 // Ставим ферзя

			// Рекурсивно пытаемся разместить остальных ферзей
			if placeQueens(board, row+1) {
				return true
			}

			// Откатываем назад, если не получилось
			board[row][col] = 0
		}
	}

	return false
}

// nQueen решает задачу N-ферзей и возвращает массив с колонками ферзей
func nQueen(n int) []int {
	// Создаём доску n x n, заполненную нулями
	board := make([][]int, n)
	for i := range board {
		board[i] = make([]int, n)
	}

	// Если есть решение
	if placeQueens(board, 0) {
		// Список номеров колонок, где стоят ферзи
		ans := []int{}
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if board[i][j] == 1 {
					ans = append(ans, j+1)
				}
			}
		}
		return ans
	}

	return []int{-1} // Если решений нет
}

func main() {
	var n int
	fmt.Scan(&n)

	ans := nQueen(n)
	for _, col := range ans {
		fmt.Print(col, " ")
	}
}
