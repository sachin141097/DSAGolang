/*
Given an integer array nums, return true if any value appears more than once in the array, otherwise return false.

Example 1:

Input: nums = [1, 2, 3, 3]

Output: true

Example 2:

Input: nums = [1, 2, 3, 4]

Output: false

Time Complexity: O(N)
Space Complexity: O(N)
*/
package main

import "fmt"

func hasDuplicate(nums []int) bool {
	seen := make(map[int]bool)
	for _, num := range nums {
		if seen[num] {
			return true
		}
		seen[num] = true
	}
	return false
}
func main() {
	nums := []int{1, 2, 3, 3, 5}
	fmt.Println(hasDuplicate(nums))
}
