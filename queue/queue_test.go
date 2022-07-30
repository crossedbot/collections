package queue

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLen(t *testing.T) {
	q := New()
	expected := 0
	require.Equal(t, expected, q.Len())
	for expected < 10 {
		expected++
		q.Push(expected)
		require.Equal(t, expected, q.Len())
	}
}

func TestPop(t *testing.T) {
	q := New()
	for v := 0; v < 10; v++ {
		q.Push(v)
	}
	for expected := 0; expected < 10; expected++ {
		actual := q.Pop()
		require.Equal(t, expected, actual)
	}
}

func TestPush(t *testing.T) {
	q := New()
	expected := 0
	require.Equal(t, expected, q.Len())
	expected++
	q.Push(expected)
	require.Equal(t, expected, q.Len())
	actual := q.Pop()
	require.Equal(t, expected, actual)
}
