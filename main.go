package main

import (
	"fmt"

	"github.com/moalf/ds/stack"
)

var s stack.Stack

func main() {

	fmt.Printf("Empty? %v;  Stack: %v\n", s.IsEmpty(), s)

	s.Push(3)
	s.Push(10)
	s.Push(23)
	fmt.Printf("Empty? %v;  Stack: %v\n", s.IsEmpty(), s)
	fmt.Println(s.Peek())
	s.Pop()
	fmt.Printf("Empty? %v;  Stack: %v\n", s.IsEmpty(), s)
	fmt.Println(s.Peek())
	s.Pop()
	fmt.Printf("Empty? %v;  Stack: %v\n", s.IsEmpty(), s)
	fmt.Println(s.Peek())
	s.Pop()
	fmt.Printf("Empty? %v;  Stack: %v\n", s.IsEmpty(), s)
	fmt.Println(s.Peek())
}
