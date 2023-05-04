package tree

import "github.com/vaibhavp1964/bp-tree/src/node"

// Traverse ...
func (t *Tree) Traverse(n *node.Node) {
	if n.IsLeaf {
		n.PrintKeys()
	}

	for _, c := range n.Children {
		node, err := t.nodeManager.Fetch(c)
		if err != nil {
			t.logger.Panicf("error retrieving node %d from persistence storage", c)
		}
		t.Traverse(node)
	}
}
