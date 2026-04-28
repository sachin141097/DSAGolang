package main

import "fmt"

func dfs(row int, col int, visited [][]bool, grid [][]int, path *[]string, row0 int, col0 int) {
	visited[row][col] = true
	*path = append(*path, fmt.Sprintf("%d,%d", row-row0, col-col0))
	rows := len(grid)
	cols := len(grid[0])

	if row-1 >= 0 && grid[row-1][col] == 1 && !visited[row-1][col] {
		dfs(row-1, col, visited, grid, path, row0, col0)
	}
	if row+1 < rows && grid[row+1][col] == 1 && !visited[row+1][col] {
		dfs(row+1, col, visited, grid, path, row0, col0)
	}
	if col-1 >= 0 && grid[row][col-1] == 1 && !visited[row][col-1] {
		dfs(row, col-1, visited, grid, path, row0, col0)
	}
	if col+1 < cols && grid[row][col+1] == 1 && !visited[row][col+1] {
		dfs(row, col+1, visited, grid, path, row0, col0)
	}
}
func countDistinctIslands(grid [][]int) int {
	rows := len(grid)
	cols := len(grid[0])
	visited := make([][]bool, rows)
	for i := range visited {
		visited[i] = make([]bool, cols)
	}
	seen := make(map[string]bool)
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if !visited[i][j] && grid[i][j] == 1 {
				var path []string
				dfs(i, j, visited, grid, &path, i, j)
				seen[fmt.Sprint(path)] = true
			}
		}
	}
	return len(seen)
}
func main() {
	grid := [][]int{{1, 1, 0, 0, 0}, {1, 1, 0, 0, 0}, {0, 0, 0, 1, 1}, {0, 0, 0, 1, 1}}
	fmt.Println(countDistinctIslands(grid))
}
