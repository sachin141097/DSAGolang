package main

import "fmt"

func productExceptItself(nums []int) []int {
	result := []int{}
	var ans int
	for i := 0; i < len(nums); i++ {
		ans = 1
		for j := 0; j < len(nums); j++ {
			if i == j {
				continue
			}
			ans *= nums[j]
		}
		result = append(result, ans)
	}
	return result
}
func main() {
	nums := []int{1, 2, 3, 0}
	fmt.Println(productExceptItself(nums))
}
