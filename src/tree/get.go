package tree

import (
	"bp-tree/src/node"
	"fmt"
)

// Get ...
func (t *Tree) Get(key []byte) ([]byte, error) {
	t.logger.Printf("Request received for finding value for key %s from tree\n", string(key))
	n := t.GetRoot()
	return t.getValue(n, key)
}

func (t *Tree) getValue(n *node.Node, key []byte) ([]byte, error) {
	if n.IsLeaf {
		if !n.CheckForKey(key) {
			t.logger.Printf("No value found for key %s in the tree\n", string(key))
			return nil, fmt.Errorf("No value found for key %s in the tree", string(key))
		}
		return n.GetValue(key)
	}

	nodeIndex, _ := n.GetChildNode(key)
	childNode, err := t.nodeManager.Fetch(nodeIndex)
	if err != nil {
		t.logger.Panicf("error retrieving node %d from persistence storage", nodeIndex)
	}
	return t.getValue(childNode, key)
}

func (t *Tree) getNodeForKey(n *node.Node, key []byte) *node.Node {
	if n.IsLeaf {
		return n
	}

	child, _ := n.GetChildNode(key)
	childNode, err := t.nodeManager.Fetch(child)
	if err != nil {
		t.logger.Panicf("error retrieving node %d from persistence storage", child)
	}
	return t.getNodeForKey(childNode, key)
}
