package main

import "fmt"

func maxHeapify(arr []int, i, heapsize int) {
	l := 2*i + 1
	r := 2*i + 2
	largest := i
	if l < heapsize && arr[l] > arr[largest] {
		largest = l
	}
	if r < heapsize && arr[r] > arr[largest] {
		largest = r
	}
	if largest != i {
		arr[i], arr[largest] = arr[largest], arr[i]
		maxHeapify(arr, largest, heapsize)

	}

}
func buildMaxheap(arr []int) {
	n := len(arr)
	for i := n/2 - 1; i >= 0; i-- {
		maxHeapify(arr, i, n)
	}

}
func main() {
	arr := []int{3, 1, 6, 5, 2, 4}
	fmt.Println("Before max heapify:", arr)
	buildMaxheap(arr)
	fmt.Println("After max heapify:", arr)

}
