package main

import "fmt"

func rottenOranges(grid [][]int) int {
	rows := len(grid)
	cols := len(grid[0])
	queue := [][]int{}
	fresh := 0
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if grid[r][c] == 2 {
				queue = append(queue, []int{r, c})
			} else if grid[r][c] == 1 {
				fresh++
			}
		}
	}
	minutes := 0
	visited := make([][]bool, rows)
	for i := range visited {
		visited[i] = make([]bool, cols)
	}
	//BFS
	for len(queue) > 0 && fresh > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			cell := queue[0]
			queue = queue[1:]
			r, c := cell[0], cell[1]
			//visit up
			if r-1 >= 0 && grid[r-1][c] == 1 && !visited[r-1][c] {
				visited[r-1][c] = true
				fresh--
				queue = append(queue, []int{r - 1, c})
			}
			//visit down
			if r+1 < rows && grid[r+1][c] == 1 && !visited[r+1][c] {
				visited[r+1][c] = true
				fresh--
				queue = append(queue, []int{r + 1, c})
			}
			//visit left
			if c-1 >= 0 && grid[r][c-1] == 1 && !visited[r][c-1] {
				visited[r][c-1] = true
				fresh--
				queue = append(queue, []int{r, c - 1})
			}
			//visit right
			if c+1 < cols && grid[r][c+1] == 1 && !visited[r][c+1] {
				visited[r][c+1] = true
				fresh--
				queue = append(queue, []int{r, c + 1})
			}

		}
		minutes++
	}
	if fresh > 0 {
		return -1
	}
	return minutes
}
func main() {
	grid := [][]int{{2, 1, 1}, {1, 1, 0}, {0, 1, 1}}
	fmt.Println(rottenOranges(grid))
}
