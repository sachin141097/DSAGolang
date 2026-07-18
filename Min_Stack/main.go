package main

import "fmt"

type entry struct {
	val int
	min int
}
type MinStack struct {
	stack []entry
}

func (s *MinStack) Push(val int) {
	if len(s.stack) == 0 {
		s.stack = append(s.stack, entry{val, val})
		return
	}
	currentMin := s.stack[len(s.stack)-1].min
	if val < currentMin {
		s.stack = append(s.stack, entry{val, val})
	} else {
		s.stack = append(s.stack, entry{val, currentMin})
	}
}
func (s *MinStack) Pop() {
	s.stack = s.stack[:len(s.stack)-1]
}
func (s *MinStack) Top() int {
	return s.stack[len(s.stack)-1].val
}
func (s *MinStack) GetMin() int {
	return s.stack[len(s.stack)-1].min
}
func main() {
	var ms MinStack
	ms.Push(5)
	fmt.Println("push(5)  -> top:", ms.Top(), " min:", ms.GetMin())

	ms.Push(3)
	fmt.Println("push(3)  -> top:", ms.Top(), " min:", ms.GetMin())

}
