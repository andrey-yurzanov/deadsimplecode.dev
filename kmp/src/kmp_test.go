package kmp_test

import (
	"slices"
	"testing"

	"deadsimplecode.dev/kmp"
)

func TestPrefix(t *testing.T) {
	prefix := kmp.Prefix([]rune("ababcabxzab"))
	if slices.Compare(prefix, []int{0, 0, 1, 2, 0, 1, 2, 0, 0, 1, 2}) != 0 {
		t.Error("prefix is incorrect")
	}
}

func TestSuccessSearch(t *testing.T) {
	pattern := []rune("ababcabxzab")
	text := []rune("qwesdababcabxzabqwe")

	index := kmp.Search(text, pattern)
	if slices.Compare(pattern, text[index:index+len(pattern)]) != 0 {
		t.Error("pattern not found")
	}
}

func TestFailureSearch(t *testing.T) {
	pattern := []rune("xwdfcvqw")
	text := []rune("qwesdababcabxzabqwe")

	index := kmp.Search(text, pattern)
	if index != -1 {
		t.Error("index is incorrect")
	}
}

func TestSuccessSearchAll(t *testing.T) {
	pattern := []rune("ababcabxzab")
	text := []rune("qwesdababcabxzabqweababcabxzabqweqababcabxzab")

	indexes := kmp.SearchAll(text, pattern)
	for _, index := range indexes {
		if slices.Compare(pattern, text[index:index+len(pattern)]) != 0 {
			t.Error("pattern not found")
		}
	}
}

func TestFailureSearchAll(t *testing.T) {
	pattern := []rune("xwdfcvqw")
	text := []rune("qwesdababcabxzabqwe")

	indexes := kmp.SearchAll(text, pattern)
	if len(indexes) > 0 {
		t.Error("index is incorrect")
	}
}
