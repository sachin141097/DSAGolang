package main

import (
	"fmt"
	"slices"
)

/*
Time Complexity: O(n^3)
*/

func longestConsecutiveSequence(arr []int) int {
	longest := 0
	for _, num := range arr {
		currentNum := num
		currentStreak := 1
		for contains(arr, currentNum+1) {
			currentNum++
			currentStreak++

		}
		if currentStreak > longest {
			longest = currentStreak
		}

	}
	return longest
}
func contains(arr []int, currentNum int) bool {
	return slices.Contains(arr, currentNum)
}
func main() {
	arr := []int{1, 2, 3, 4}
	fmt.Println(longestConsecutiveSequence(arr))
}
