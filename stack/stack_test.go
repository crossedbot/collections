package stack

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLen(t *testing.T) {
	s := New()
	expected := 0
	require.Equal(t, expected, s.Len())
	for expected < 10 {
		expected++
		s.Push(expected)
		require.Equal(t, expected, s.Len())
	}
}

func TestPop(t *testing.T) {
	s := New()
	for v := 0; v < 10; v++ {
		s.Push(v)
	}
	for expected := 9; expected > 0; expected-- {
		actual := s.Pop()
		require.Equal(t, expected, actual)
	}
}

func TestPush(t *testing.T) {
	s := New()
	expected := 0
	require.Equal(t, expected, s.Len())
	expected++
	s.Push(expected)
	require.Equal(t, expected, s.Len())
	actual := s.Pop()
	require.Equal(t, expected, actual)
}
