package main

import "fmt"

func main() {

	var n int
	fmt.Scanf("%d\n", &n)

	for i:=0; i<n; i++ {

		var lines []Stack
		var curLine string
		var x, y int
		fmt.Scanf("%d %d\n", &x, &y)
		fmt.Scanln(&curLine)


	}

}

type Stack struct {
	top *Element
	size int
}

type Element struct {
	value interface{}
	next *Element
}

// Return the stack's length
func (s *Stack) Len() int {
	return s.size
}

// Push a new element onto the stack
func (s *Stack) Push(value interface{}) {
	s.top = &Element{value, s.top}
	s.size++
}

// Remove the top element from the stack and return it's value
// If the stack is empty, return nil
func (s *Stack) Pop() (value interface{}) {
	if s.size > 0 {
		value, s.top = s.top.value, s.top.next
		s.size--
		return
	}
	return nil
}