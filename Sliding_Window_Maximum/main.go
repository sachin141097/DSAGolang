package main

import "fmt"

func maximumSlidingWindow(arr []int, k int) []int {
	n := len(arr)
	var dq []int
	result := make([]int, n-k+1)
	for i := 0; i < n; i++ {
		for len(dq) > 0 && dq[0] <= i-k {
			dq = dq[1:]
		}
		for len(dq) > 0 && arr[dq[len(dq)-1]] < arr[i] {
			dq = dq[:len(dq)-1]
		}
		dq = append(dq, i)
		if i >= k-1 {
			result[i-k+1] = arr[dq[0]]
		}
	}
	return result

}
func main() {
	arr, k := []int{1, 3, -1, -3, 5, 3, 6, 7}, 3
	fmt.Println(maximumSlidingWindow(arr, k))
}
