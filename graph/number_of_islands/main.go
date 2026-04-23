package main

import "fmt"

func dfs(i int, j int, grid [][]byte, visited [][]bool, N int, M int) {
	visited[i][j] = true
	//visit up
	if i-1 >= 0 && grid[i-1][j] == '1' && !visited[i-1][j] {
		dfs(i-1, j, grid, visited, N, M)
	}
	//visit down
	if i+1 < N && grid[i+1][j] == '1' && !visited[i+1][j] {
		dfs(i+1, j, grid, visited, N, M)
	}
	//visit left
	if j-1 >= 0 && grid[i][j-1] == '1' && !visited[i][j-1] {
		dfs(i, j-1, grid, visited, N, M)
	}
	//visit right
	if j+1 < M && grid[i][j+1] == '1' && !visited[i][j+1] {
		dfs(i, j+1, grid, visited, N, M)
	}
	//visit up-left
	if i-1 >= 0 && j-1 >= 0 && grid[i-1][j-1] == '1' && !visited[i-1][j-1] {
		dfs(i-1, j-1, grid, visited, N, M)
	}
	//visit down-right
	if i+1 < N && j+1 < M && grid[i+1][j+1] == '1' && !visited[i+1][j+1] {
		dfs(i+1, j+1, grid, visited, N, M)
	}
	//visit up-right
	if i-1 >= 0 && j+1 < M && grid[i-1][j+1] == '1' && !visited[i-1][j+1] {
		dfs(i-1, j+1, grid, visited, N, M)
	}
	//viist down-left
	if i+1 < N && j-1 >= 0 && grid[i+1][j-1] == '1' && !visited[i+1][j-1] {
		dfs(i+1, j-1, grid, visited, N, M)
	}

}
func NumIslands(grid [][]byte) int {
	N := len(grid)
	M := len(grid[0])
	visited := make([][]bool, N)
	for i := range visited {
		visited[i] = make([]bool, M)
	}
	components := 0
	for i := range grid {
		for j := range grid[i] {
			if !visited[i][j] && grid[i][j] == '1' {
				dfs(i, j, grid, visited, N, M)
				components += 1
			}
		}
	}
	return components

}
func main() {
	grid := [][]byte{
		{'1', '1', '1', '0', '1'},
		{'1', '0', '0', '0', '0'},
		{'1', '1', '1', '0', '1'},
		{'0', '0', '0', '1', '1'},
	}

	// Function call to find the number of islands in given grid
	ans := NumIslands(grid)

	// Output
	fmt.Println("The total islands in given grids are:", ans)
}
