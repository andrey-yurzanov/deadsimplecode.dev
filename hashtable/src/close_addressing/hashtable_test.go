package hashtable_test

import (
	"fmt"
	"testing"

	"deadsimplecode.dev/hashtable"
)

func TestPut(t *testing.T) {
	count := 100

	htable := hashtable.HashTable[int, any]{}
	for i := 0; i < count; i++ {
		htable.Put(i, i)
	}

	if htable.GetSize() != uint64(count) {
		t.Error("size should be equals: " + fmt.Sprint(count))
	}

	for i := 0; i < count; i++ {
		if htable.Get(i) != i {
			t.Error("value not found by key: " + fmt.Sprint(i))
		}
	}
}

func TestDelete(t *testing.T) {
	count := 100

	htable := hashtable.HashTable[int, any]{}
	for i := 0; i < count; i++ {
		htable.Put(i, i)
	}

	for i := 0; i < count; i++ {
		htable.Delete(i)
	}

	if htable.GetSize() != 0 {
		t.Error("size should be equals zero")
	}
}
