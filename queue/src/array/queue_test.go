package queue_test

import (
	"testing"

	collection "deadsimplecode.dev/queue/array"
)

func TestSuccess(t *testing.T) {
	queue := collection.Queue[int]{}
	for i := 0; i < 100; i++ {
		queue.Insert(i)
	}

	for i := 0; i < 50; i++ {
		if queue.Remove() != i {
			t.Error("Value not equals")
		}
	}

	for i := 100; i < 200; i++ {
		queue.Insert(i)
	}

	for i := 50; i < 200; i++ {
		if queue.Remove() != i {
			t.Error("Value not equals")
		}
	}
}

func TestFail(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("Panic is expected")
		}
	}()

	queue := collection.Queue[int]{}
	queue.Remove()
}
