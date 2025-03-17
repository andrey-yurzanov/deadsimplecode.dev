package binarysearch_test

import (
	"slices"
	"testing"

	"deadsimplecode.dev/binarysearch"
)

func TestSuccessBinarySearch(t *testing.T) {
	result := binarysearch.BinarySearch([]int{2, 2, 2, 4, 5, 6}, 2)
	if slices.Compare(result, []int{2, 1, 0}) != 0 {
		t.Error("result should be 2")
	}
}

func TestNotFoundBinarySearch(t *testing.T) {
	result := binarysearch.BinarySearch([]int{1, 2, 3, 4, 5, 6}, 7)
	if len(result) != 0 {
		t.Error("result should be empty")
	}
}
