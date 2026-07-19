package main

import "fmt"

func nextSmallestFromLeft(arr []int) []int {
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
func nextSmallestFromRight(arr []int) []int {
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
func nextLargestFromLeft(arr []int) []int {
	n := len(arr)
	ans := make([]int, n)
	var stack []int
	for i := 0; i < n; i++ {
		for len(stack) > 0 && arr[stack[len(stack)-1]] <= arr[i] {
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
func nextLargestFromRight(arr []int) []int {
	n := len(arr)
	ans := make([]int, n)
	var stack []int
	for i := n - 1; i >= 0; i-- {
		for len(stack) > 0 && arr[stack[len(stack)-1]] < arr[i] {
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
func findSumOfSubArrayMaximum(arr []int) int64 {
	sum := int64(0)
	left := nextLargestFromLeft(arr)
	right := nextLargestFromRight(arr)
	for i := 0; i < len(arr); i++ {
		startingPoint := i - left[i]
		endingPoint := right[i] - i
		sum += ((int64)(arr[i]) * (int64)(startingPoint) * (int64)(endingPoint))
	}
	return sum

}
func findSumofSubArrayMinimum(arr []int) int64 {
	sum := int64(0)
	left := nextSmallestFromLeft(arr)
	right := nextSmallestFromRight(arr)
	for i := 0; i < len(arr); i++ {
		startingPoint := i - left[i]
		endingPoint := right[i] - i
		sum += ((int64)(arr[i]) * (int64)(startingPoint) * (int64)(endingPoint))
	}
	return sum

}
func findSumOfSubArrayRanges(arr []int) int64 {
	sumOfSubArrayMinimum := findSumofSubArrayMinimum(arr)
	sumOfSubArrayMaximum := findSumOfSubArrayMaximum(arr)
	return sumOfSubArrayMaximum - sumOfSubArrayMinimum

}
func main() {
	arr := []int{1, 2, 3}
	fmt.Println(findSumOfSubArrayRanges(arr))
}
