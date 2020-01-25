package binaryheap

// BinaryHeap is a priority queue.
// This will be a max-heap.
// TODO: generalizes BinaryHeap, make it being able
// to store any type of values from a total order set.
type BinaryHeap struct {
	data []int
}

// UnderflowError represents a heap underflow error.
type UnderflowError struct{}

func (err UnderflowError) Error() string {
	return "heap underflow"
}

func newBinaryHeapUnderflowError() UnderflowError {
	return UnderflowError{}
}

func parent(i int) int {
	return (i - 1) / 2
}

func left(i int) int {
	return 2*i + 1
}

func right(i int) int {
	return 2*i + 2
}

func (heap *BinaryHeap) swap(a, b int) {
	heap.data[a], heap.data[b] = heap.data[b], heap.data[a]
}

func (heap *BinaryHeap) maxHeapify(i int) {
	largest := i
	if l := left(i); l < heap.Len() && heap.data[largest] < heap.data[l] {
		largest = l
	}
	if r := right(i); r < heap.Len() && heap.data[largest] < heap.data[r] {
		largest = r
	}
	if largest != i {
		heap.swap(largest, i)
		heap.maxHeapify(largest)
	}
}

// NewBinaryHeap returns an empty binary heap.
func NewBinaryHeap() BinaryHeap {
	return BinaryHeap{
		data: make([]int, 0),
	}
}

// Len returns the number of elements in the heap.
func (heap *BinaryHeap) Len() int {
	return len(heap.data)
}

// Push pushes a new item into a binary heap.
func (heap *BinaryHeap) Push(item int) {
	heap.data = append(heap.data, item)
	for i := heap.Len() - 1; i > 0 && heap.data[parent(i)] < heap.data[i]; i = parent(i) {
		heap.swap(parent(i), i)
	}
}

// Pop pops out an item from a binary heap.
func (heap *BinaryHeap) Pop() (int, error) {
	if heap.Len() == 0 {
		return 0, newBinaryHeapUnderflowError()
	}
	result := heap.data[0]
	// Swap the root and the last leaf
	heap.swap(0, heap.Len()-1)
	// Shrink the heap
	heap.data = heap.data[:heap.Len()-1]
	heap.maxHeapify(0)

	return result, nil
}

// FromSlice constructs a binary heap from a slice
func FromSlice(s []int) BinaryHeap {
	heap := BinaryHeap{s}

	for i := parent(heap.Len() - 1); i >= 0; i-- {
		heap.maxHeapify(i)
	}

	return heap
}

// IntoSortedSlice returns the undelying data of the heap by a sorted slice in descending order.
func (heap *BinaryHeap) IntoSortedSlice() []int {
	s := make([]int, 0, heap.Len())
	for v, err := heap.Pop(); err == nil; v, err = heap.Pop() {
		s = append(s, v)
	}
	return s
}
