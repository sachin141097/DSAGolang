package main

import "fmt"

func dfs(board [][]byte, row int, col int, visited [][]bool, rows int, cols int) {
	visited[row][col] = true
	if row-1 >= 0 && board[row-1][col] == 'O' && !visited[row-1][col] {
		dfs(board, row-1, col, visited, rows, cols)
	}
	if row+1 < rows && board[row+1][col] == 'O' && !visited[row+1][col] {
		dfs(board, row+1, col, visited, rows, cols)
	}
	if col-1 >= 0 && board[row][col-1] == 'O' && !visited[row][col-1] {
		dfs(board, row, col-1, visited, rows, cols)
	}
	if col+1 < cols && board[row][col+1] == 'O' && !visited[row][col+1] {
		dfs(board, row, col+1, visited, rows, cols)
	}
}
func solve(board [][]byte) {
	rows := len(board)
	cols := len(board[0])
	visited := make([][]bool, rows)
	for i := range visited {
		visited[i] = make([]bool, cols)
	}

	//Traverse from bounary and mark safe '0'
	for col := 0; col < cols; col++ {
		if board[0][col] == 'O' && !visited[0][col] {
			dfs(board, 0, col, visited, rows, cols)
		}
		if board[rows-1][col] == 'O' && !visited[rows-1][col] {
			dfs(board, rows-1, col, visited, rows, cols)
		}
	}

	for row := 0; row < rows; row++ {
		if board[row][0] == 'O' && !visited[row][0] {
			dfs(board, row, 0, visited, rows, cols)
		}
		if board[row][cols-1] == 'O' && !visited[row][cols-1] {
			dfs(board, row, cols-1, visited, rows, cols)
		}
	}
	//Flip surrounded 'O' to 'X'
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if board[i][j] == 'O' && !visited[i][j] {
				board[i][j] = 'X'
			}
		}
	}
}
func main() {
	board := [][]byte{
		{'X', 'X', 'X', 'X'},
		{'X', 'O', 'O', 'X'},
		{'X', 'X', 'O', 'X'},
		{'X', 'O', 'X', 'X'},
	}
	solve(board)
	for _, row := range board {
		fmt.Println(string(row))
	}
}
