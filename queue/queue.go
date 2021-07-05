package queue

import "fmt"

type Queue struct {
	queue []int
}

func (q *Queue) Enqueue(d int) {
	q.queue = append(q.queue, d)
}

func (q *Queue) Dequeue() int {
	dequeued := q.queue[0]
	q.queue = q.queue[1:]
	return dequeued
}

func (q *Queue) Front() int {
	return q.queue[0]
}

func (q *Queue) Size() int {
	return len(q.queue)
}

func (q *Queue) Display() {
	if q.Size() == 0 {
		fmt.Println("Queue is empty")
		return
	}
	for _, e := range q.queue {
		fmt.Printf("%+v ~>", e)
	}
	fmt.Println()
}
