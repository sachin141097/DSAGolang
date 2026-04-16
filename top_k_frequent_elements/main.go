package main

import (
	"fmt"
	"sort"
)

/*
Given an integer array nums and an integer k, return the k most frequent elements within the array.

The test cases are generated such that the answer is always unique.

You may return the output in any order.

Example 1:

Input: nums = [1,2,2,3,3,3], k = 2

Output: [2,3]
Example 2:

Input: nums = [7,7], k = 1

Output: [7]
*/
func findtopK(nums []int, k int) []int {
	freq := make(map[int]int)
	for _, num := range nums {
		freq[num]++
	}
	keys := []int{}
	for num := range freq {
		keys = append(keys, num)
	}
	//sort keys by their frequency value descending
	sort.Slice(keys, func(i, j int) bool { return freq[keys[i]] > freq[keys[j]] })
	return keys[:k]

}
func main() {
	nums := []int{7, 7}
	k := 1
	fmt.Println(findtopK(nums, k))
}
