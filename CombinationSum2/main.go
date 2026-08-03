package main

import (
	"fmt"
	"sort"
)

func findTargetSum(candidates []int, startIndex int, result *[]int, ans *[][]int, target int) {
	if target < 0 {
		return
	}
	if target == 0 {
		temp := make([]int, len(*result))
		copy(temp, *result)
		*ans = append(*ans, temp)
		return
	}
	for i := startIndex; i < len(candidates); i++ {
		if i > startIndex && candidates[i] == candidates[i-1] {
			continue
		}
		*result = append(*result, candidates[i])
		findTargetSum(candidates, i+1, result, ans, target-candidates[i])
		*result = (*result)[:len(*result)-1]
	}
}

func combinationSum2(candidates []int, target int) [][]int {
	sort.Slice(candidates, func(i, j int) bool {
		return candidates[i] < candidates[j]
	})
	var ans [][]int
	var result []int
	findTargetSum(candidates, 0, &result, &ans, target)
	return ans
}
func main() {
	candidates := []int{10, 1, 2, 7, 6, 1, 5}
	target := 8
	fmt.Println(combinationSum2(candidates, target))
}
