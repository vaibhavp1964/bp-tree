package node

import (
	keyvalue "bp-tree/src/key_value"
	"bytes"
	"fmt"
	"log"
	"sort"
)

// Node ...
type Node struct {
	Index    int
	IsLeaf   bool
	Order    int
	Children []int
	KVPairs  []keyvalue.KeyValue
	Right    int
	logger   *log.Logger
}

// CheckForKey ...
func (n *Node) CheckForKey(key []byte) bool {
	for _, kv := range n.KVPairs {
		if kv.HasSameKey(key) {
			return true
		}
	}
	return false
}

// GetValue ...
func (n *Node) GetValue(key []byte) ([]byte, error) {
	// return error for internal node and absence of key
	if !n.IsLeaf {
		return nil, fmt.Errorf("value cannot be fetched from an internal node")
	}

	for _, kv := range n.KVPairs {
		if kv.HasSameKey(key) {
			return kv.Value, nil
		}
	}

	return nil, fmt.Errorf("no value found for key %s", string(key))
}

// InsertKVPair ...
func (n *Node) InsertKVPair(key, value []byte) error {
	kv := keyvalue.GetKeyValue(string(key), string(value))

	// if key already present, modify that particular keyvalue pair
	if n.CheckForKey(key) {
		index, _ := n.GetKeyValueIndex(key)
		n.KVPairs[index] = kv
		return nil
	}

	// if node is full, return error
	if n.IsFull() {
		return fmt.Errorf("Node completely filled. Cannot tolerate any more insertions")
	}

	n.KVPairs = append(n.KVPairs, kv)
	sort.SliceStable(n.KVPairs, func(i, j int) bool {
		return bytes.Compare(n.KVPairs[i].Key, n.KVPairs[j].Key) < 0
	})
	return nil
}

// IsFull ...
func (n *Node) IsFull() bool {
	return len(n.KVPairs) == n.Order
}

// GetKeyValueIndex ...
func (n *Node) GetKeyValueIndex(key []byte) (int, bool) {
	for index, kv := range n.KVPairs {
		if kv.HasSameKey(key) {
			return index, true
		}
	}
	return -1, false
}

// CanInsertKeyValue ...
func (n *Node) CanInsertKeyValue(key []byte) bool {
	return !n.IsFull() || n.CheckForKey(key)
}

// GetChildNode ...
func (n *Node) GetChildNode(key []byte) (int, error) {
	if n.IsLeaf {
		return -1, fmt.Errorf("Cannot return child node for leaf node")
	}

	largestKey := n.KVPairs[len(n.KVPairs)-1]
	if bytes.Compare(key, largestKey.Key) >= 0 {
		return n.Children[len(n.KVPairs)], nil
	}

	for i, kv := range n.KVPairs {
		if bytes.Compare(key, kv.Key) < 0 {
			return n.Children[i], nil
		}
	}

	return -1, nil
}

// GetChildNodeIndex ...
func (n *Node) GetChildNodeIndex(node *Node) (int, bool) {
	if n.IsLeaf {
		return -1, false
	}

	for i, index := range n.Children {
		if index == node.Index {
			return i, true
		}
	}
	return -1, false
}

// GetFirstKey ...
func (n *Node) GetFirstKey() []byte {
	if len(n.KVPairs) == 0 {
		return []byte("")
	}

	return n.KVPairs[0].Key
}

// DeleteKVPair ...
func (n *Node) DeleteKVPair(key []byte) bool {
	newKVPairs := make([]keyvalue.KeyValue, 0)

	for _, kv := range n.KVPairs {
		if bytes.Compare(kv.Key, key) != 0 {
			newKVPairs = append(newKVPairs, kv)
		}
	}

	n.KVPairs = newKVPairs
	return true
}
