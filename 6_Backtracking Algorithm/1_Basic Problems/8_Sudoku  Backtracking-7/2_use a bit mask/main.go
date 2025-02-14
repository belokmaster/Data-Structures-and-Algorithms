package main

import (
	"fmt"
)

/*
Дано неполное судоку в виде матрицы mat[][] порядка 9*9, задача состоит в том, чтобы завершить судоку.
Решение судоку должно удовлетворять  всем следующим правилам :

Каждая из цифр  1-9 должна встречаться ровно один раз в каждой строке.
Каждая из цифр  1-9 должна встречаться ровно один раз в каждом столбце.
Каждая из цифр  1-9 должна встречаться ровно один раз в каждом из 9  3x3 подполей сетки.
*/

/*
В приведенном выше подходе функция isSafe() , которая используется для проверки того,
безопасно ли помещать число num в ячейку (i, j), ищет num в каждой строке , столбце и поле .
Идея состоит в том, чтобы оптимизировать это с помощью битовой маски .
Для этого создайте три массива rows[], cols[], boxs[] размером n, чтобы отметить использованное
значение в строке , столбце и поле соответственно. Элемент row[i] отмечает число, уже использованное
в строке i, и то же самое делают cols[] и boxs[] для столбцов и полей. Чтобы отметить число num строки i ,
 установите бит num слева от row[ i ] и действуйте аналогично для cols[] и boxs[] .
 Аналогично, чтобы снять отметку со значения num , снимите биты , установленные на текущем шаге.
*/

const N = 9

func isSafe(board *[N][N]int, row, col, num int, rowMask, colMask, boxMask *[N]int) bool {
	boxIndex := (row/3)*3 + col/3
	if (rowMask[row]&(1<<num)) != 0 || (colMask[col]&(1<<num)) != 0 || (boxMask[boxIndex]&(1<<num)) != 0 {
		return false
	}
	return true
}

func solveSudokuRec(board *[N][N]int, row, col int, rowMask, colMask, boxMask *[N]int) bool {
	if row == N-1 && col == N {
		return true
	}
	if col == N {
		row++
		col = 0
	}
	if board[row][col] != 0 {
		return solveSudokuRec(board, row, col+1, rowMask, colMask, boxMask)
	}

	for num := 1; num <= N; num++ {
		if isSafe(board, row, col, num, rowMask, colMask, boxMask) {
			board[row][col] = num
			boxIndex := (row/3)*3 + col/3
			rowMask[row] |= (1 << num)
			colMask[col] |= (1 << num)
			boxMask[boxIndex] |= (1 << num)

			if solveSudokuRec(board, row, col+1, rowMask, colMask, boxMask) {
				return true
			}

			board[row][col] = 0
			rowMask[row] &^= (1 << num)
			colMask[col] &^= (1 << num)
			boxMask[boxIndex] &^= (1 << num)
		}
	}

	return false
}

func solveSudoku(board *[N][N]int) {
	var rowMask, colMask, boxMask [N]int

	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			if board[i][j] != 0 {
				num := board[i][j]
				boxIndex := (i/3)*3 + j/3
				rowMask[i] |= (1 << num)
				colMask[j] |= (1 << num)
				boxMask[boxIndex] |= (1 << num)
			}
		}
	}

	solveSudokuRec(board, 0, 0, &rowMask, &colMask, &boxMask)
}

func printBoard(board [N][N]int) {
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			fmt.Printf("%d ", board[i][j])
		}
		fmt.Println()
	}
}

func main() {
	board := [N][N]int{
		{3, 0, 6, 5, 0, 8, 4, 0, 0},
		{5, 2, 0, 0, 0, 0, 0, 0, 0},
		{0, 8, 7, 0, 0, 0, 0, 3, 1},
		{0, 0, 3, 0, 1, 0, 0, 8, 0},
		{9, 0, 0, 8, 6, 3, 0, 0, 5},
		{0, 5, 0, 0, 9, 0, 6, 0, 0},
		{1, 3, 0, 0, 0, 0, 2, 5, 0},
		{0, 0, 0, 0, 0, 0, 0, 7, 4},
		{0, 0, 5, 2, 0, 6, 3, 0, 0},
	}

	solveSudoku(&board)
	printBoard(board)
}
