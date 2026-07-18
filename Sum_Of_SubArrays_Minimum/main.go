package main

import "fmt"

const MOD = 1_000_000_007

func findNextSmallestFromLeft(arr []int) []int {
	n := len(arr)
	ans := make([]int, n)
	var stack []int
	for i := 0; i < n; i++ {
		for len(stack) > 0 && arr[stack[len(stack)-1]] >= arr[i] {
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
func findNextSmallestFromRight(arr []int) []int {
	n := len(arr)
	ans := make([]int, n)
	var stack []int
	for i := n - 1; i >= 0; i-- {
		for len(stack) > 0 && arr[stack[len(stack)-1]] > arr[i] {
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
func findSum(arr []int) int {
	n := len(arr)
	left := findNextSmallestFromLeft(arr)
	right := findNextSmallestFromRight(arr)
	var sum int64 = 0
	for i := 0; i < n; i++ {
		startingPoint := i - left[i]
		endingPoint := right[i] - i
		sum += ((int64)(arr[i]) * (int64)(startingPoint) * (int64)(endingPoint)) % MOD
	}
	return int(sum % MOD)
}
func main() {
	arr := []int{3, 1, 2, 5}
	fmt.Println(findSum(arr))
}
