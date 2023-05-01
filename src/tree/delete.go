package tree

import "bp-tree/src/node"

// Delete ...
func (t *Tree) Delete(key []byte) ([]byte, bool) {
	t.logger.Printf("Request received for deleting entry for key %s from tree\n", string(key))
	rootNode := t.GetRoot()
	return t.delete(rootNode, key)
}

func (t *Tree) delete(n *node.Node, key []byte) ([]byte, bool) {
	if n.IsLeaf {
		if !n.CheckForKey(key) {
			t.logger.Printf("Key %s is not present in tree\n", string(key))
			return nil, false
		}

		val, _ := n.GetValue(key)
		n.DeleteKVPair(key)
		t.nodeManager.Flush(n)
		t.logger.Printf("Entry corresponding to key %s successfully deleted from the tree\n", string(key))
		return val, true
	}

	child, _ := n.GetChildNode(key)
	childNode, err := t.nodeManager.Fetch(child)
	if err != nil {
		t.logger.Panicf("error retrieving node %d from persistence storage", child)
	}
	return t.delete(childNode, key)
}
