package queue

import (
	"container/heap"
	"time"
)

// pqItem represents a priority queue item
type pqItem struct {
	Exp   time.Time   // expiration time
	Index int         // queue index
	Key   string      // item key
	Value interface{} // item value
}

// pqHeap implements a prioritized "mini-heap". Where a heap is a tree with the
// property that each node is the minimum-valued node in its subtree; see the
// container/heap package for more information
type pqHeap []*pqItem

// Len returns the length of the heap
func (h *pqHeap) Len() int {
	c := *h
	return len(c)
}

// Less returns true if the item at index i will expired before the item at
// index j. If any item is nil, it is considered "lesser" than the other item
// and true/false is returned for index i/j respectively
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

// Push pushes and appends the given item at the end of the prioritized heap
func (h *pqHeap) Push(x interface{}) {
	if item, ok := x.(*pqItem); ok {
		item.Index = h.Len()
		*h = append(*h, item)
	}
}

// Pop returns and removes the item at index len(heap)-1; IE. the item at the
// end of the heap
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

// Swap swaps the position of the items at index i and j
func (h *pqHeap) Swap(i, j int) {
	c := *h
	c[i], c[j] = c[j], c[i]
	c[i].Index = i
	c[j].Index = j
	*h = c
}

// PriorityQueue represents an interface to a priority queue for items with a
// Time-To-Live (TTL)
type PriorityQueue interface {
	// Add adds an item to the priority queue with the given value and TTL
	Add(key string, value interface{}, ttl time.Duration)

	// Delete deletes the item with the given key from the priority queue
	Delete(key string)

	// DeleteExpired deletes all items from the priority queue that have
	// expired based on their TTL
	DeleteExpired(t time.Time) int

	// Get returns the item for a given key and increases its TTL on
	// retrieval
	Get(key string, ttl time.Duration) interface{}

	// Len returns the length of the priority queue
	Len() int
}

// pq implements a PriorityQueue for items with expiration times
type pq struct {
	Heap *pqHeap            // The priority queue heap
	Map  map[string]*pqItem // Map of priority queue items and their keys
}

// NewPriorityQueue returns a new PriorityQueue
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
