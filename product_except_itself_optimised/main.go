package main

import "fmt"

func productexceptItself(arr []int) []int {
	result := make([]int, len(arr))
	left := make([]int, len(arr))
	right := make([]int, len(arr))
	left[0] = 1
	for i := 1; i < len(arr); i++ {
		left[i] = left[i-1] * arr[i-1]
	}
	right[len(arr)-1] = 1
	for j := len(arr) - 2; j >= 0; j-- {
		right[j] = right[j+1] * arr[j+1]
	}
	for i := 0; i < len(arr); i++ {
		result[i] = left[i] * right[i]
	}
	return result
}
func main() {
	arr := []int{10, 20, 30}
	fmt.Println(productexceptItself(arr))
}
