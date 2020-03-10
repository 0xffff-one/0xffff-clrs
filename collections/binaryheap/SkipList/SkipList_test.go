package SkipList

import (
	"testing"
)

func TestInsertion(t *testing.T) {
	var i uint64
	var val int = 1
	s := NEWSKIPLIST(6)
	for i = 1; i < 10; i++ {
		s.insert(i, val)
		val++
	}
}

func TestSearch(t *testing.T) {
	var i uint64
	s := NEWSKIPLIST(6)
	for i = 1; i < 10; i++ {
		s.searchWithPreviousNodes(i)
	}
}

func TestDelete(t *testing.T) {
	var i uint64
	s := NEWSKIPLIST(6)
	for i = 1; i < 10; i++ {
		s.delete(i)
	}
}
