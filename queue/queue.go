package queue

import ()

type Queue interface {
	Len() int

	Pop() interface{}

	Push(value interface{})
}

type queue struct {
	data []interface{}
}

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
