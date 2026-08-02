package main

import "fmt"

func findTargetSum(candidates []int, startIndex int, target int, ans *[][]int, result *[]int, k int) {
	if target < 0 {
		return
	}
	if target == 0 && len(*result) == k {
		temp := make([]int, len(*result))
		copy(temp, *result)
		*ans = append(*ans, temp)
		return
	}
	for i := startIndex; i < len(candidates); i++ {
		*result = append(*result, candidates[i])
		findTargetSum(candidates, i+1, target-candidates[i], ans, result, k)
		*result = (*result)[:len(*result)-1]
	}

}
func combinationSum3(k, target int) [][]int {
	var ans [][]int
	var result []int
	candidates := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	findTargetSum(candidates, 0, target, &ans, &result, k)
	return ans
}
func main() {
	k, target := 3, 7
	fmt.Println(combinationSum3(k, target))
}
