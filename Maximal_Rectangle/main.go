package main

import "fmt"

func nextSmallestFromLeft(height []int) []int {
	n := len(height)
	ans := make([]int, n)
	var stack []int
	for i := 0; i < n; i++ {
		for len(stack) > 0 && height[stack[len(stack)-1]] >= height[i] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			ans[i] = -1
		} else {
			ans[i] = stack[len(stack)-1]
		}
		stack = append(stack, i)
	}
	return ans

}
func nextSmallestFromRight(height []int) []int {
	n := len(height)
	ans := make([]int, n)
	var stack []int
	for i := n - 1; i >= 0; i-- {
		for len(stack) > 0 && height[stack[len(stack)-1]] >= height[i] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			ans[i] = n
		} else {
			ans[i] = stack[len(stack)-1]
		}
		stack = append(stack, i)
	}
	return ans

}
func findLargestRectangle(height []int) int {
	maxArea := 0
	n := len(height)
	right := nextSmallestFromRight(height)
	left := nextSmallestFromLeft(height)
	for i := 0; i < n; i++ {
		area := (height[i]) * (right[i] - left[i] - 1)
		if area > maxArea {
			maxArea = area
		}
	}
	return maxArea

}
func findMaximalRectangle(matrix [][]int) int {
	rows := len(matrix)
	cols := len(matrix[0])
	prefixedMatrix := make([][]int, rows)
	for i := range prefixedMatrix {
		prefixedMatrix[i] = make([]int, cols)
	}
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if i == 0 {
				prefixedMatrix[i][j] = matrix[i][j]
			} else if matrix[i][j] == 1 {
				prefixedMatrix[i][j] = prefixedMatrix[i-1][j] + matrix[i][j]
			}
		}
	}
	maxArea := 0
	for i := 0; i < rows; i++ {
		area := findLargestRectangle(prefixedMatrix[i])
		if area > maxArea {
			maxArea = area
		}
	}
	return maxArea

}
func main() {
	matrix := [][]int{
		{1, 0, 1, 0, 0},
		{1, 0, 1, 1, 1},
		{1, 1, 1, 1, 1},
		{1, 0, 0, 1, 0},
	}
	fmt.Println(findMaximalRectangle(matrix))
}
