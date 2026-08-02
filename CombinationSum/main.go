package main

import "fmt"

func findTargetSum(candidates []int, startIndex int, target int, result *[]int, ans *[][]int) {
	if target < 0 {
		return
	}
	if target == 0 {
		//Make a copy because result will keep changing during backtracking
		combinations := make([]int, len(*result))
		copy(combinations, *result)
		*ans = append(*ans, combinations)
	}
	for i := startIndex; i < len(candidates); i++ {
		*result = append(*result, candidates[i])
		findTargetSum(candidates, i, target-candidates[i], result, ans)
		*result = (*result)[:len(*result)-1]
	}
}
func combinationSum(candidates []int, target int) [][]int {
	var ans [][]int
	var result []int
	findTargetSum(candidates, 0, target, &result, &ans)
	return ans
}
func main() {
	candidates := []int{2, 3, 6, 7}
	target := 7
	ans := combinationSum(candidates, target)
	fmt.Println(ans)
}
