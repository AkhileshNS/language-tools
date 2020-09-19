/*
 ** Stack
 ** New - new Stack(value)
 ** Push - .push(value)
 ** Pop - .pop()
 ** Peek - .peek()
 */

// =============================================
package main

import "fmt"

type stack []int

func newStack(value int) stack {
	return stack{value}
}

func (s *stack) push(value int) {
	*s = append(*s, value)
}

func (s *stack) pop() int {
	top := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return top
}

func (s stack) peek() int {
	return s[len(s)-1]
}

func main() {
	s := newStack(1)
	fmt.Println(s)
	s.push(10)
	fmt.Println(s)
	fmt.Println(s.pop())
	fmt.Println(s.peek())
	fmt.Println(s)
}
