package SkipList

import (
	"testing"
)

func TestInsertion(t *testing.T) {
	head := SkipListNode{}
	list := SkipList{head: &head, tail: nil}
	for i := 0; i <= 100; i++ {
		if (*SkipList).insert(&list, i, i) {
			t.Error(`Insertion failed`)
		}
	}
}
