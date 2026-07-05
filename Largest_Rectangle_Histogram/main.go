package main

import "fmt"

func findNextSmallerFromLeft(height []int) []int {
	n := len(height)
	ans := make([]int, n)
	stack := make([]int, 0, n)
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
func findNextSmallerFromRight(height []int) []int {
	n := len(height)
	ans := make([]int, n)
	stack := make([]int, 0, n)
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
func main() {
	height := []int{3, 5, 1, 7, 5, 9}
	right := findNextSmallerFromRight(height)
	left := findNextSmallerFromLeft(height)
	maxArea := 0
	for i := 0; i < len(height); i++ {
		currentArea := height[i] * (right[i] - left[i] - 1)
		if currentArea > maxArea {
			maxArea = currentArea
		}
	}
	fmt.Println(maxArea)
}
