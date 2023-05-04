package tree

import (
	"bytes"
	"sort"

	keyvalue "github.com/vaibhavp1964/bp-tree/src/key_value"
	"github.com/vaibhavp1964/bp-tree/src/node"
)

// Put ...
func (t *Tree) Put(key, value []byte) error {
	t.logger.Printf("Request received for putting key %s, value %s in tree\n", string(key), string(value))
	n := t.GetRoot()
	returnNode, partition, err := t.putInTree(n, key, value)
	if err != nil {
		return err
	}

	if partition == nil {
		return nil
	}

	newRootNode := node.NewInternalNode(t.counter.Get(), t.order, t.logger)

	newRootNode.KVPairs = append(newRootNode.KVPairs, keyvalue.GetKeyValue(string(partition), ""))
	newRootNode.Children = append(newRootNode.Children, n.Index)
	newRootNode.Children = append(newRootNode.Children, returnNode.Index)

	sort.Slice(newRootNode.Children, func(i, j int) bool {
		iNode, err := t.nodeManager.Fetch(newRootNode.Children[i])
		if err != nil {
			t.logger.Panicf("error retrieving node %d from persistence storage", newRootNode.Children[i])
		}
		jNode, err := t.nodeManager.Fetch(newRootNode.Children[j])
		if err != nil {
			t.logger.Panicf("error retrieving node %d from persistence storage", newRootNode.Children[j])
		}
		return bytes.Compare(iNode.GetFirstKey(), jNode.GetFirstKey()) < 0
	})
	t.rootNodeIndex = newRootNode.Index

	t.nodeManager.Flush(newRootNode)

	return nil
}

func (t *Tree) putInTree(n *node.Node, key, value []byte) (*node.Node, []byte, error) {
	// for leaf node
	if n.IsLeaf {
		// if key already exists, just replace value
		if n.CheckForKey(key) {
			index, _ := n.GetKeyValueIndex(key)
			n.KVPairs[index] = keyvalue.GetKeyValue(string(key), string(value))
			t.nodeManager.Flush(n)
			return nil, nil, nil
		}

		// leaf node not saturated. KV pair can be inserted
		if !n.IsFull() {
			err := n.InsertKVPair(key, value)
			if err != nil {
				t.logger.Fatalf("Cannot insert key %s value %s in tree due to error %s", string(key), string(value), err)
			}
			t.nodeManager.Flush(n)
			return nil, nil, err
		}

		// leaf node saturated. New node needs to be created
		newNode := t.partitionLeafNode(n, key, value)
		t.nodeManager.Flush(newNode)
		t.nodeManager.Flush(n)
		return newNode, newNode.KVPairs[0].Key, nil
	}

	// find appropriate child node to go to
	child, _ := n.GetChildNode(key)
	childNode, err := t.nodeManager.Fetch(child)

	if err != nil {
		t.logger.Panicf("error retrieving node %d from persistence storage", child)
	}

	// go to leaf node
	newNode, partitionKey, err := t.putInTree(childNode, key, value)
	if err != nil {
		t.logger.Fatalf("Cannot insert key %s value %s in tree due to error %s", string(key), string(value), err)
		return nil, nil, err
	}

	// if no partition key is received back, return
	if partitionKey == nil {
		return nil, nil, nil
	}

	// if partition key is received but internal node is not saturated, insert kv pair, manage pointers and return
	if !n.IsFull() {
		err := n.InsertKVPair(partitionKey, nil)
		n.Children = append(n.Children, newNode.Index)
		sort.Slice(n.Children, func(i, j int) bool {
			iNode, err := t.nodeManager.Fetch(n.Children[i])
			if err != nil {
				t.logger.Panicf("error retrieving node %d from persistence storage", n.Children[i])
			}
			jNode, err := t.nodeManager.Fetch(n.Children[j])
			if err != nil {
				t.logger.Panicf("error retrieving node %d from persistence storage", n.Children[j])
			}
			return bytes.Compare(iNode.GetFirstKey(), jNode.GetFirstKey()) < 0
		})
		t.nodeManager.Flush(n)
		return nil, nil, err
	}

	// partition internal node
	n.KVPairs = append(n.KVPairs, keyvalue.GetKeyValue(string(partitionKey), ""))
	n.Children = append(n.Children, newNode.Index)
	newInternalNode, internalPartitionKey := t.partitionInternalNode(n)
	t.nodeManager.Flush(n)
	t.nodeManager.Flush(newInternalNode)
	return newInternalNode, internalPartitionKey, nil
}

func (t *Tree) partitionInternalNode(n *node.Node) (*node.Node, []byte) {
	children := make([]int, 0)
	children = append(children, n.Children...)
	sort.Slice(children, func(i, j int) bool {
		iNode, err := t.nodeManager.Fetch(children[i])
		if err != nil {
			t.logger.Panicf("error retrieving node %d from persistence storage", children[i])
		}
		jNode, err := t.nodeManager.Fetch(children[j])
		if err != nil {
			t.logger.Panicf("error retrieving node %d from persistence storage", children[j])
		}
		return bytes.Compare(iNode.GetFirstKey(), jNode.GetFirstKey()) < 0
	})

	kvpairs := make([]keyvalue.KeyValue, 0)
	kvpairs = append(kvpairs, n.KVPairs...)
	sort.Slice(kvpairs, func(i, j int) bool {
		return bytes.Compare(kvpairs[i].Key, kvpairs[j].Key) < 0
	})

	newNode := node.NewInternalNode(t.counter.Get(), t.order, t.logger)

	mid := len(kvpairs) / 2
	n.KVPairs = make([]keyvalue.KeyValue, 0)
	n.KVPairs = append(n.KVPairs, kvpairs[:mid]...)

	newNode.KVPairs = make([]keyvalue.KeyValue, 0)
	newNode.KVPairs = append(newNode.KVPairs, kvpairs[mid+1:]...)

	partition := kvpairs[mid]
	keys := make([]string, 0)
	for _, kv := range kvpairs {
		keys = append(keys, string(kv.Key))
	}

	mid++
	n.Children = make([]int, 0)
	n.Children = append(n.Children, children[:mid]...)
	newNode.Children = make([]int, 0)
	newNode.Children = append(newNode.Children, children[mid:]...)

	return newNode, partition.Key
}

func (t *Tree) partitionLeafNode(n *node.Node, key, value []byte) *node.Node {
	newNode := node.NewLeafNode(t.counter.Get(), t.order, t.logger)

	newKV := keyvalue.GetKeyValue(string(key), string(value))

	// divide kv pairs
	kvpairs := make([]keyvalue.KeyValue, 0)
	kvpairs = append(kvpairs, n.KVPairs...)
	kvpairs = append(kvpairs, newKV)
	sort.Slice(kvpairs, func(i, j int) bool {
		return bytes.Compare(kvpairs[i].Key, kvpairs[j].Key) < 0
	})
	mid := len(kvpairs) / 2

	n.KVPairs = make([]keyvalue.KeyValue, 0)
	n.KVPairs = append(n.KVPairs, kvpairs[:mid]...)
	newNode.KVPairs = make([]keyvalue.KeyValue, 0)
	newNode.KVPairs = append(newNode.KVPairs, kvpairs[mid:]...)

	// readjust right pointer
	newNode.Right = n.Right
	n.Right = newNode.Index

	return newNode
}

func insertAtIndex(slice []*node.Node, child *node.Node, index int) []*node.Node {
	if len(slice) == index { // nil or empty slice or after last element
		return append(slice, child)
	}
	slice = append(slice[:index+1], slice[index:]...) // index < len(a)
	slice[index] = child
	return slice
}
