package queue

import ()

// Queue represents an interface to queue collection
type Queue interface {
	// Len returns the length of the queue
	Len() int

	// Pop returns and removes the next item in the queue
	Pop() interface{}

	// Push pushes a new item onto the queue
	Push(value interface{})
}

// queue implements the Queue collection
type queue struct {
	data []interface{}
}

// New returns a new Queue
func New() Queue {
	return &queue{}
}

func (q *queue) Len() int {
	return len(q.data)
}

func (q *queue) Pop() interface{} {
	if q.Len() == 0 {
		return nil
	}
	item := q.data[0]
	q.data = q.data[1:]
	return item
}

func (q *queue) Push(value interface{}) {
	q.data = append(q.data, value)
}
