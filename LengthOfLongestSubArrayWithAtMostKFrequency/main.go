package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 1, 2, 3, 1, 2}
	k := 2
	fmt.Println(maxSubArrayLength(nums, k))
}
func maxSubArrayLength(nums []int, k int) int {
	left := 0
	maxLength := 0
	freq := make(map[int]int)
	for right := 0; right < len(nums); right++ {
		freq[nums[right]]++
		for freq[nums[right]] > k {
			freq[nums[left]]--
			left++
		}
		maxLength = max(maxLength, right-left+1)
	}
	return maxLength
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
