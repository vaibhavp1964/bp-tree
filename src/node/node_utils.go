package node

import (
	keyvalue "bp-tree/src/key_value"
	"fmt"
	"log"
	"strconv"
	"strings"
	// "strings"
)

// NewLeafNode ...
func NewLeafNode(index, order int, logger *log.Logger) *Node {
	logger.Println("Received request to create new leaf node with index:", index)
	return &Node{
		Index:    index,
		Order:    order,
		IsLeaf:   true,
		Children: nil,
		KVPairs:  make([]keyvalue.KeyValue, 0),
		Right:    -1,
		logger:   logger,
	}
}

// NewInternalNode ...
func NewInternalNode(index, order int, logger *log.Logger) *Node {
	logger.Println("Received request to create new internal node with index:", index)
	return &Node{
		Index:    index,
		Order:    order,
		IsLeaf:   false,
		KVPairs:  make([]keyvalue.KeyValue, 0),
		Children: make([]int, 0),
		Right:    -1,
		logger:   logger,
	}
}

// PrintKeys ...
func (n *Node) PrintKeys() {
	keys := make([]string, 0)
	for _, kv := range n.KVPairs {
		keys = append(keys, string(kv.Key))
	}
	if n.IsLeaf {
		fmt.Println("keys for leaf node", n.Index, "=", strings.Join(keys, ","))
		return
	}
	fmt.Println("keys for node", n.Index, "=", strings.Join(keys, ","))
}

// GetChildren ...
func (n *Node) GetChildren() []int {
	if n.IsLeaf {
		return nil
	}

	children := make([]int, 0)
	for _, c := range n.Children {
		children = append(children, c)
	}
	return children
}

// PrintChildren ...
func (n *Node) PrintChildren() {
	children := make([]string, 0)
	for _, c := range n.Children {
		children = append(children, strconv.Itoa(c))
	}
	if n.IsLeaf {
		fmt.Println("children for leaf node", n.Index, "=", strings.Join(children, ","))
		return
	}
	fmt.Println("children for internal node", n.Index, "=", strings.Join(children, ","))
}
