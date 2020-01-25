package binaryheap

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func (heap *BinaryHeap) walkImpl(root int, f func(root *int, l *int, r *int)) {
	if heap == nil || heap.Len() == 0 {
		return
	}

	var triple = [3]*int{&heap.data[root]}

	if l := left(root); l < heap.Len() {
		triple[1] = &heap.data[l]
		heap.walkImpl(l, f)
	}
	if r := right(root); r < heap.Len() {
		triple[2] = &heap.data[r]
		heap.walkImpl(r, f)
	}

	f(triple[0], triple[1], triple[2])
}

func (heap *BinaryHeap) walk(f func(root *int, l *int, r *int)) {
	heap.walkImpl(0, f)
}

func (heap *BinaryHeap) verify(t *testing.T) {
	heap.walk(func(root *int, l *int, r *int) {
		if l != nil {
			assert.Greater(t, *root, *l)
		}
		if r != nil {
			assert.Greater(t, *root, *r)
		}
	})
}

func TestFromSlice(t *testing.T) {
	cases := [][]int{
		{1, 2, 3},
		{3, 1, 2},
		{1, 3, 4, 2},
		{2, 1, 4, 3},
		{5, 2, 4, 1, 3},
		{5, 1, 4, 2, 3},
		{1, 2, 3, 4, 5},
	}
	for _, slice := range cases {
		heap := FromSlice(slice)
		heap.verify(t)
	}
}

func TestPush(t *testing.T) {
	cases := [][]int{
		{1, 2, 3},
		{3, 1, 2},
		{1, 3, 4, 2},
		{2, 1, 4, 3},
		{5, 2, 4, 1, 3},
		{5, 1, 4, 2, 3},
		{1, 2, 3, 4, 5},
	}

	for _, slice := range cases {
		heap := NewBinaryHeap()
		for _, v := range slice {
			heap.Push(v)
			heap.verify(t)
		}
	}
}

func TestPop(t *testing.T) {
	cases := [][]int{
		{1, 2, 3},
		{3, 1, 2},
		{1, 3, 4, 2},
		{2, 1, 4, 3},
		{5, 2, 4, 1, 3},
		{5, 1, 4, 2, 3},
		{1, 2, 3, 4, 5},
	}
	for _, slice := range cases {
		heap := FromSlice(slice)
		for v, err := heap.Pop(); err == nil; v, err = heap.Pop() {
			heap.walk(func(root *int, _ *int, _ *int) {
				assert.Greater(t, v, *root)
			})
			heap.verify(t)
		}
	}
}

func TestIntoSortedSlice(t *testing.T) {
	cases := [][]int{
		{1, 2, 3},
		{3, 1, 2},
		{1, 3, 4, 2},
		{2, 1, 4, 3},
		{5, 2, 4, 1, 3},
		{5, 1, 4, 2, 3},
		{1, 2, 3, 4, 5},
	}

	for _, slice := range cases {
		heap := FromSlice(slice)
		s := heap.IntoSortedSlice()
		// TODO: extract the logic of determining whether a slice is sorted.
		for i := 1; i < len(s); i++ {
			assert.Less(t, s[i], s[i-1])
		}
	}
}
