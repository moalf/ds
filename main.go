package main

import (
	"fmt"

	"github.com/moalf/ds/queue"
	"github.com/moalf/ds/stack"
)

var s stack.Stack
var q queue.Queue

func main() {
	fmt.Println("DS Demo")

	fmt.Println("Stack ***")
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

	fmt.Println("Queue ***")
	fmt.Printf("Empty? %v; Queue: %v\n", q.IsEmpty(), q)
	q.Enqueue(10)
	fmt.Printf("Empty? %v; Queue: %v\n", q.IsEmpty(), q)
	q.Enqueue(20)
	fmt.Printf("Empty? %v; Queue: %v\n", q.IsEmpty(), q)
	q.Enqueue(30)
	fmt.Printf("Empty? %v; Queue: %v\n", q.IsEmpty(), q)
	q.Enqueue(40)
	fmt.Printf("Empty? %v; Queue: %v\n", q.IsEmpty(), q)
	q.Dequeue()
	fmt.Printf("Empty? %v; Queue: %v\n", q.IsEmpty(), q)
	q.Dequeue()
	fmt.Printf("Empty? %v; Queue: %v\n", q.IsEmpty(), q)
	q.Dequeue()
	fmt.Printf("Empty? %v; Queue: %v\n", q.IsEmpty(), q)
	q.Dequeue()
	fmt.Printf("Empty? %v; Queue: %v\n", q.IsEmpty(), q)

	q.Enqueue(100)
	fmt.Printf("Empty? %v; Queue: %v\n", q.IsEmpty(), q)
	q.Dequeue()
	fmt.Printf("Empty? %v; Queue: %v\n", q.IsEmpty(), q)

	fmt.Println("Done1")
	fmt.Println("Done2")
}
