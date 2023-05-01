package tree

import (
	keyvalue "bp-tree/src/key_value"
	"bytes"
)

// RangeQuery ...
func (t *Tree) RangeQuery(startKey, endKey []byte) []keyvalue.KeyValue {
	n := t.GetRoot()
	node := t.getNodeForKey(n, startKey)

	kvPairs := make([]keyvalue.KeyValue, 0)

	for {
		for _, kv := range node.KVPairs {
			if bytes.Compare(kv.Key, endKey) > 0 {
				return kvPairs
			}
			if bytes.Compare(startKey, kv.Key) <= 0 {
				kvPairs = append(kvPairs, kv)
			}
		}
		nodeIndex := node.Right

		if nodeIndex == -1 {
			return kvPairs
		}

		fetchedNode, err := t.nodeManager.Fetch(nodeIndex)
		if err != nil {
			t.logger.Fatalf("error retreiving node %d from persistence storage: %s", nodeIndex, err)
		}
		node = fetchedNode
	}
}
