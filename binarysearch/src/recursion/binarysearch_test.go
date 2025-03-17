package binarysearch_test

import (
	"testing"

	"deadsimplecode.dev/binarysearch"
)

func TestSuccessBinarySearch(t *testing.T) {
	result := binarysearch.BinarySearch([]int{1, 2, 3, 4, 5, 6}, 2)
	if result != 1 {
		t.Error("result should be 1")
	}
}

func TestNotFoundBinarySearch(t *testing.T) {
	result := binarysearch.BinarySearch([]int{1, 2, 3, 4, 5, 6}, 7)
	if result != -1 {
		t.Error("result should be -1")
	}
}
