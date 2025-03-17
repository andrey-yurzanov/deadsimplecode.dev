package linkedlist_test

import (
	"testing"

	linkedlist "deadsimplecode.dev/linkedlist"
)

func TestInsertFirst(t *testing.T) {
	list := linkedlist.LinkedList[int]{}

	for i := 0; i < 100; i++ {
		list.InsertFirst(i)
		if list.GetFirst() != i {
			t.Error("values not equals")
		}
	}
}

func TestInsertLast(t *testing.T) {
	list := linkedlist.LinkedList[int]{}

	for i := 0; i < 100; i++ {
		list.InsertLast(i)
		if list.GetLast() != i {
			t.Error("values not equals")
		}
	}
}

func TestRemoveFirst(t *testing.T) {
	list := linkedlist.LinkedList[int]{}

	for i := 0; i < 100; i++ {
		list.InsertFirst(i)
		list.RemoveFirst()
		if int(list.Len()) != 0 {
			t.Error("list should be empty")
		}
	}
}

func TestRemoveLast(t *testing.T) {
	list := linkedlist.LinkedList[int]{}

	for i := 0; i < 100; i++ {
		list.InsertLast(i)
		list.RemoveLast()
		if int(list.Len()) != 0 {
			t.Error("list should be empty")
		}
	}
}
