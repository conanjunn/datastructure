package trie

import (
	"testing"
)

func TestTrie(t *testing.T) {
	a := NewTrie()
	a.Add("abc")
	ret := a.Contains("abc")
	t.Logf("%v", ret)
}
