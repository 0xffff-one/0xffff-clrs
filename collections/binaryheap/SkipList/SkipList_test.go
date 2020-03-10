package SkipList

import (
	"testing"
)

func TestInsertion(t *testing.T) {
	var i uint64
	var val int = 1
	s := NEWSKIPLIST(6)
	for i = 1; i < 10; i++ {
		if !s.insert(i, val) {
			t.Error("[-]Insertion failed")
		}
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

func TestShow(t *testing.T) {
	var i uint64
	s := NEWSKIPLIST(6)
	for i = 1; i < 10; i++ {
		s.show()
	}
}
func TestDelete(t *testing.T) {
	var i uint64
	s := NEWSKIPLIST(6)
	for i = 1; i < 10; i++ {
		if !s.delete(i) {
			t.Error("[-]Deletion failed")
		}

	}
}
