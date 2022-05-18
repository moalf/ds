package queue

import "fmt"

type Queue []uint

func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}

func (q *Queue) Enqueue(elem uint) {
	*q = append(*q, elem)
}

func (q *Queue) Dequeue() {
	if !(*q).IsEmpty() {
		*q = append((*q)[1:len(*q)])
	} else {
		fmt.Println("Queue is empty")
	}
}
