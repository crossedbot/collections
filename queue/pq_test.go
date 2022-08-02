package queue

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestPqHeapLen(t *testing.T) {
	h := new(pqHeap)
	expected := 0
	require.Equal(t, expected, h.Len())

	for expected < 10 {
		expected++
		item := &pqItem{Value: expected}
		h.Push(item)
		require.Equal(t, expected, h.Len())
	}
}

func TestPqHeapLess(t *testing.T) {
	now := time.Now()
	h := new(pqHeap)

	item1 := &pqItem{Value: "one", Exp: now}
	h.Push(item1)
	item2 := &pqItem{Value: "two", Exp: now.Add(time.Second * 30)}
	h.Push(item2)

	require.True(t, h.Less(0, 1))
	require.False(t, h.Less(1, 0))
}

func TestPqHeapPush(t *testing.T) {
	now := time.Now()
	h := new(pqHeap)

	item1 := &pqItem{Value: "one", Exp: now}
	h.Push(item1)
	item2 := &pqItem{Value: "two", Exp: now.Add(time.Second * 30)}
	h.Push(item2)
	require.Equal(t, 2, h.Len())

	actual := (*h)[0]
	require.Equal(t, item1.Value, actual.Value)
	require.Equal(t, item1.Exp.UnixNano(), actual.Exp.UnixNano())
	require.Zero(t, actual.Index)

	actual = (*h)[1]
	require.Equal(t, item2.Value, actual.Value)
	require.Equal(t, item2.Exp.UnixNano(), actual.Exp.UnixNano())
	require.Equal(t, 1, actual.Index)
}

func TestPqHeapPop(t *testing.T) {
	now := time.Now()
	h := new(pqHeap)

	item1 := &pqItem{Value: "one", Exp: now}
	h.Push(item1)
	item2 := &pqItem{Value: "two", Exp: now.Add(time.Second * 30)}
	h.Push(item2)
	require.Equal(t, 2, h.Len())

	actual, ok := h.Pop().(*pqItem)
	require.True(t, ok)
	require.Equal(t, item2.Value, actual.Value)
	require.Equal(t, item2.Exp.UnixNano(), actual.Exp.UnixNano())
	require.Equal(t, 1, actual.Index)

	actual, ok = h.Pop().(*pqItem)
	require.True(t, ok)
	require.Equal(t, item1.Value, actual.Value)
	require.Equal(t, item1.Exp.UnixNano(), actual.Exp.UnixNano())
	require.Zero(t, actual.Index)
}

func TestPqHeapSwap(t *testing.T) {
	now := time.Now()
	h := new(pqHeap)

	item1 := &pqItem{Value: "one", Exp: now}
	h.Push(item1)
	item2 := &pqItem{Value: "two", Exp: now.Add(time.Second * 30)}
	h.Push(item2)
	require.Equal(t, 2, h.Len())

	h.Swap(0, 1)

	actual := (*h)[0]
	require.Equal(t, item2.Value, actual.Value)
	require.Equal(t, item2.Exp.UnixNano(), actual.Exp.UnixNano())
	require.Zero(t, actual.Index)

	actual = (*h)[1]
	require.Equal(t, item1.Value, actual.Value)
	require.Equal(t, item1.Exp.UnixNano(), actual.Exp.UnixNano())
	require.Equal(t, 1, actual.Index)
}

func TestPqAdd(t *testing.T) {
	key := "key"
	value := "value"
	ttl := time.Second * 30
	q := &pq{
		Heap: new(pqHeap),
		Map:  make(map[string]*pqItem, 0),
	}
	q.Add(key, value, ttl)
	require.Equal(t, 1, q.Heap.Len())
	require.Equal(t, 1, len(q.Map))

	actual := (*q.Heap)[0]
	require.Equal(t, key, actual.Key)
	require.Equal(t, value, actual.Value)
	require.Zero(t, actual.Index)

	actual, ok := q.Map[key]
	require.True(t, ok)
	require.Equal(t, value, actual.Value)
	require.Zero(t, actual.Index)
}

func TestPqDelete(t *testing.T) {
	key := "key"
	value := "value"
	ttl := time.Second * 30
	q := &pq{
		Heap: new(pqHeap),
		Map:  make(map[string]*pqItem, 0),
	}
	q.Add(key, value, ttl)
	require.Equal(t, 1, q.Heap.Len())
	require.Equal(t, 1, len(q.Map))
	q.Delete(key)
	require.Zero(t, q.Heap.Len())
	require.Zero(t, len(q.Map))
}

func TestPqDeleteExpired(t *testing.T) {
	key := "key"
	value := "value"
	ttl := time.Millisecond * 100
	q := &pq{
		Heap: new(pqHeap),
		Map:  make(map[string]*pqItem, 0),
	}
	q.Add(key, value, ttl)
	require.Equal(t, 1, q.Heap.Len())
	require.Equal(t, 1, len(q.Map))
	time.Sleep(ttl)
	actual := q.DeleteExpired(time.Now())
	require.Equal(t, 1, actual)
	require.Zero(t, q.Heap.Len())
	require.Zero(t, len(q.Map))
}

func TestPqGet(t *testing.T) {
	key := "key"
	value := "value"
	ttl := time.Second * 30
	q := &pq{
		Heap: new(pqHeap),
		Map:  make(map[string]*pqItem, 0),
	}
	q.Add(key, value, ttl)
	actual, ok := q.Get(key, ttl).(string)
	require.True(t, ok)
	require.Equal(t, value, actual)
}

func TestPqLen(t *testing.T) {
	key := "key"
	value := "value"
	ttl := time.Second * 30
	q := &pq{
		Heap: new(pqHeap),
		Map:  make(map[string]*pqItem, 0),
	}
	q.Add(key, value, ttl)
	require.Equal(t, 1, q.Len())
}
