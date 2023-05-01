package tree

import "bp-tree/src/node"

// GetRoot ...
func (t *Tree) GetRoot() *node.Node {
	node, err := t.nodeManager.Fetch(t.rootNodeIndex)
	if err != nil {
		t.logger.Fatalf("")
	}
	return node
}
