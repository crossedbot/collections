package queue

import (
	"container/heap"
	"time"
)

type pqItem struct {
	Exp   time.Time
	Index int
	Key   string
	Value interface{}
}

type pqHeap []*pqItem

func (h *pqHeap) Len() int {
	c := *h
	return len(c)
}

func (h *pqHeap) Less(i, j int) bool {
	// Put NULL values at the start of the queue to be removed upon POP
	c := *h
	a := c[i]
	if a == nil {
		return true
	}
	b := c[j]
	if b == nil {
		return false
	}
	// Otherwise order by expiration time
	return a.Exp.Before(b.Exp)
}

func (h *pqHeap) Push(x interface{}) {
	if item, ok := x.(*pqItem); ok {
		item.Index = h.Len()
		*h = append(*h, item)
	}
}

func (h *pqHeap) Pop() interface{} {
	c := *h
	l := len(c)
	if l == 0 {
		return nil
	}
	item := c[l-1]
	c = c[:l-1]
	*h = c
	return item
}

func (h *pqHeap) Swap(i, j int) {
	c := *h
	c[i], c[j] = c[j], c[i]
	c[i].Index = i
	c[j].Index = j
	*h = c
}

type PriorityQueue interface {
	Add(key string, value interface{}, ttl time.Duration)
	Delete(key string)
	DeleteExpired(t time.Time) int
	Get(key string, ttl time.Duration) interface{}
	Len() int
}

type pq struct {
	Heap *pqHeap
	// XXX we could use an unsigned integer but strings are more versatile
	Map map[string]*pqItem
}

func NewPriorityQueue() PriorityQueue {
	return &pq{
		Heap: new(pqHeap),
		Map:  make(map[string]*pqItem, 0),
	}
}

func (q *pq) Add(key string, value interface{}, ttl time.Duration) {
	item := &pqItem{
		Exp:   time.Now().Add(ttl),
		Key:   key,
		Value: value,
	}
	heap.Push(q.Heap, item)
	q.Map[key] = item
}

func (q *pq) Delete(key string) {
	if item, ok := q.Map[key]; ok {
		delete(q.Map, key)
		heap.Remove(q.Heap, item.Index)
	}
}

func (q *pq) DeleteExpired(now time.Time) int {
	count := 0
	for q.Heap.Len() != 0 {
		item := (*q.Heap)[0]
		if now.Before(item.Exp) {
			break
		}
		delete(q.Map, item.Key)
		heap.Pop(q.Heap)
		count++
	}
	return count
}

func (q *pq) Get(key string, ttl time.Duration) interface{} {
	if item, ok := q.Map[key]; ok {
		item.Exp = time.Now().Add(ttl)
		heap.Fix(q.Heap, item.Index)
		return item.Value
	}
	return nil
}

func (q *pq) Len() int {
	return q.Heap.Len()
}
