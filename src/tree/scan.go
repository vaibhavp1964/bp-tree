package tree

import (
	keyvalue "bp-tree/src/key_value"
)

// Scan ...
func (t *Tree) Scan() []keyvalue.KeyValue {
	kvSet := make([]keyvalue.KeyValue, 0)

	node, _ := t.nodeManager.Fetch(t.leftMost)
	for {
		keys := make([]string, 0)
		for _, kv := range node.KVPairs {
			keys = append(keys, string(kv.Key))
		}

		kvSet = append(kvSet, node.KVPairs...)

		nodeIndex := node.Right

		if nodeIndex == -1 {
			return kvSet
		}

		fetchedNode, err := t.nodeManager.Fetch(nodeIndex)
		if err != nil {
			t.logger.Fatalf("error retreiving node %d from persistence storage: %s", nodeIndex, err)
		}
		node = fetchedNode
	}
}
