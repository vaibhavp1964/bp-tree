package nodemanager

import (
	"bp-tree/src/diskmanager"
	"bp-tree/src/node"
	"bytes"
	"encoding/gob"
	"log"
)

// NodeManager ...
type NodeManager struct {
	FileName  string
	BlockSize int
	DM        diskmanager.DiskManager
	Logger    *log.Logger
}

// Flush ...
func (nm *NodeManager) Flush(n *node.Node) error {
	// calculate offset
	offset := n.Index * nm.BlockSize

	// serialize node
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(n)

	if err != nil {
		nm.Logger.Fatalln("error encountered in serializing node data:", err)
		return err
	}

	additional := make([]byte, nm.BlockSize-buf.Len())
	for i := range additional {
		additional[i] = byte(' ')
	}
	additional = append(buf.Bytes(), additional...)

	err = nm.DM.Write(nm.FileName, offset, additional)
	if err != nil {
		nm.Logger.Fatalf("error encountered while writing to file %s: %s", nm.FileName, err)
	}
	return err
}

// Fetch ...
func (nm *NodeManager) Fetch(index int) (*node.Node, error) {
	// calculate offset
	offset := index * nm.BlockSize

	// get data from diskManager
	b, err := nm.DM.Read(nm.FileName, offset, nm.BlockSize)
	if err != nil {
		nm.Logger.Printf("Error encountered while reading from file %s: %s", nm.FileName, err)
		return nil, err
	}

	b = bytes.Trim(b, " ")

	// deserialize bytes
	buf := bytes.NewBuffer(b)
	dec := gob.NewDecoder(buf)

	node := &node.Node{}
	if err := dec.Decode(node); err != nil {
		nm.Logger.Printf("Error encountered while deserializing node data: %s", err)
		return nil, err
	}

	return node, nil
}
