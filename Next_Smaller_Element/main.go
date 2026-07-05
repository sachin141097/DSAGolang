package main

import "fmt"

func findNextSmaller(arr []int) []int {
	n := len(arr)
	ans := make([]int, n)
	stack := make([]int, 0, n)
	for i := n - 1; i >= 0; i-- {
		for len(stack) > 0 && stack[len(stack)-1] >= arr[i] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			ans[i] = -1
		} else {
			ans[i] = stack[len(stack)-1]
		}
		stack = append(stack, arr[i])
	}
	return ans

}
func main() {
	arr := []int{3, 5, 1, 7, 5, 9}
	fmt.Println(findNextSmaller(arr))
}
