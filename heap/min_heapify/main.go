package main

import "fmt"

func minHeapify(arr []int, i, heapSize int) {
	l := 2*i + 1
	r := 2*i + 2
	smallest := i
	if l < heapSize && arr[l] < arr[smallest] {
		smallest = l

	}
	if r < heapSize && arr[r] < arr[smallest] {
		smallest = r
	}
	if smallest != i {
		arr[i], arr[smallest] = arr[smallest], arr[i]
		minHeapify(arr, smallest, heapSize)
	}

}
func buildMinHeap(arr []int) {
	n := len(arr)
	for i := n/2 - 1; i >= 0; i-- {
		minHeapify(arr, i, n)
	}
}
func main() {
	arr := []int{3, 1, 6, 5, 2, 4}
	fmt.Println("Before min heap:", arr)
	buildMinHeap(arr)
	fmt.Println("After min heap:", arr)

}
