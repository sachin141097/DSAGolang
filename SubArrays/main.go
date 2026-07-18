package main

import "fmt"

/*
1
1 2
1 2 3
1 2 3 4
1 2 3 4 5
2
2 3
2 3 4
2 3 4 5
3
3 4
3 4 5
4
4 5
5
*/

func findSubArrays(arr []int) [][]int {
	var ans [][]int
	var temp []int
	n := len(arr)
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			temp = []int{}
			for k := i; k <= j; k++ {
				temp = append(temp, arr[k])
			}
			ans = append(ans, temp)
		}
	}
	return ans
}
func main() {
	arr := []int{3, 1, 2, 5}
	fmt.Println(findSubArrays(arr))
}
