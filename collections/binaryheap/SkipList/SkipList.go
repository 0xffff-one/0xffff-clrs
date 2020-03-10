package SkipList

import (
	"fmt"
	"math/rand"
	"time"
)

type ConcurrentSkipList struct {
	skipLists []*SKIPLIST
	level     int
}

// SKIPLIST skiplist
type SKIPLIST struct {
	level  int
	length int32
	head   *Node
	tail   *Node
	rand   *rand.Rand
	//vmutex  sync.RWMutex
}

type Node struct {
	index     uint64
	value     int
	nextNodes []*Node
}

func (s *SKIPLIST) createNode(index uint64, value, level int) *Node {

	newNode := &Node{index, value, make([]*Node, s.randomLevel())}
	return newNode
}

func (s *SKIPLIST) randomLevel() int {

	level := 1
	for ; level < s.level && s.rand.Uint32()&0x1 == 1; level++ {
	}
	return level

}

func (s *SKIPLIST) searchWithPreviousNodes(index uint64) ([]*Node, *Node) {
	// Store all previous value whose index is less than index and whose next value's index is larger than index.
	previousNodes := make([]*Node, s.level)

	currentNode := s.head

	// Iterate from top level to bottom level.
	for l := s.level - 1; l >= 0; l-- {
		// Iterate value util value's index is >= given index.
		// The max iterate count is skip list's length. So the worst O(n) is N.
		for currentNode.nextNodes[l] != s.tail && currentNode.nextNodes[l].index < index {
			currentNode = currentNode.nextNodes[l]
		}

		// When next value's index is >= given index, add current value whose index < given index.
		previousNodes[l] = currentNode
	}

	// Avoid point to tail which will occur panic in Insert and Delete function.
	// When the next value is tail.
	// The index is larger than the maximum index in the skip list or skip list's length is 0. Don't point to tail.
	// When the next value isn't tail.
	// Next value's index must >= given index. Point to it.
	if currentNode.nextNodes[0] != s.tail {
		currentNode = currentNode.nextNodes[0]
	}

	return previousNodes, currentNode
}

func (s *SKIPLIST) insert(index uint64, value int) {
	// Write lock and unlock.

	previousNodes, currentNode := s.searchWithPreviousNodes(index)

	// 如果相应index的节点已存在，直接改值
	if currentNode != s.head && currentNode.index == index {
		currentNode.value = value
		return
	}

	// Make a new value.
	newNode := s.createNode(index, value, s.randomLevel())

	// Adjust pointer. Similar to update linked list.
	for i := len(newNode.nextNodes) - 1; i >= 0; i-- {
		// Firstly, new value point to next value.
		newNode.nextNodes[i] = previousNodes[i].nextNodes[i]

		// Secondly, previous nodes point to new value.
		previousNodes[i].nextNodes[i] = newNode

		// Finally, in order to release the slice, point to nil.
		previousNodes[i] = nil
	}

	// atomic.AddInt32(&s.length, 1)

	// 如果previousNode的长度大于newNode，经历上述清空后还有继续清空previousNode
	for i := len(newNode.nextNodes); i < len(previousNodes); i++ {
		previousNodes[i] = nil
	}

	fmt.Println("[+]Insertion success.index=", index)
}

func (s *SKIPLIST) delete(index uint64) {

	previousNodes, currentNode := s.searchWithPreviousNodes(index)

	// If skip list length is 0 or could not find value with the given index.
	if currentNode != s.head && currentNode.index == index {
		// Adjust pointer. Similar to update linked list.
		for i := 0; i < len(currentNode.nextNodes); i++ {
			previousNodes[i].nextNodes[i] = currentNode.nextNodes[i]
			currentNode.nextNodes[i] = nil
			previousNodes[i] = nil
		}
	}

	for i := len(currentNode.nextNodes); i < len(previousNodes); i++ {
		previousNodes[i] = nil
	}
}
func (s *SKIPLIST) show() {
	var currentNode *Node = s.head
	for i := s.level - 1; i >= 0; i-- {
		fmt.Print("level:", i+1, " ")
		for {
			if currentNode != nil {
				fmt.Print(currentNode.value, " ")
				if currentNode.nextNodes[i] != nil {
					fmt.Print("-->")
				}
				currentNode = currentNode.nextNodes[i]
			} else {
				break
			}

		}
		fmt.Printf("\n")
		currentNode = s.head
	}
}

//NEWSKIPLIST Create skiplist
func NEWSKIPLIST(level int) *SKIPLIST {

	list := &SKIPLIST{}

	if level <= 0 {

		level = 32

	}

	list.level = level

	list.head = &Node{nextNodes: make([]*Node, level, level)}

	list.tail = nil

	list.rand = rand.New(rand.NewSource(time.Now().UnixNano()))

	for index := range list.head.nextNodes {

		list.head.nextNodes[index] = list.tail

	}

	return list

}
