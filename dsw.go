// Package dsw implements Stout-Warren, 1986 which "takes an arbitrary binary
// search tree and rebalances it to form another of optimal shape, using time
// linear in the number of nodes and only a constant amount of space
// (beyond that used to store the initial tree)".
// See http://web.eecs.umich.edu/~qstout/pap/CACM86.pdf for details.
package dsw

// TreeOps interface exposes our DFW implementation
type TreeOps interface {
	Balance(Tree) (*Tree, error)
}

// Tree struct represents our tree in memory
type Tree struct {
	Root  *Node
	Left  *Tree
	Right *Tree
	Val   int
}

// Node is one element in the tree
type Node struct {
	Parent   *Node
	Entries  []*Entry
	Children []*Node
}

// Entry is a single key-val pair in the node
type Entry struct {
	Key interface{}
	Val interface{}
}

func (t *Tree) treeToVine(n *Node) (*Tree, error) {
	return t, nil
}

func (t *Tree) vineToTree(n *Node, s int) (*Tree, error) {
	return t, nil
}

func (t *Tree) compress(n *Node, c int) (*Tree, error) {
	return t, nil
}
