package dsw

import (
	"math/rand"
	"testing"
)

// insert puts an int into the selected node
func insert(t *Tree, v int) *Tree {
	if t == nil {
		return &Tree{nil, nil, nil, v}
	}
	if v < t.Val {
		t.Left = insert(t.Left, v)
	}
	t.Right = insert(t.Right, v)
	return t
}

// new builds a random tree with k, 2k, ..., 10k int vals
func new(k int) *Tree {
	var t *Tree
	for _, v := range rand.Perm(10) {
		t = insert(t, (1+v)*k)
	}
	return t
}

// TestRandomTree builds a naive tree like the one in the Go tour
func TestRandomTree(t *testing.T) {
	tree := new(10)
	if tree.Left == nil || tree.Right == nil {
		t.Fatal("new tree not created successfully")
	}
	t.Log("OK")
}
