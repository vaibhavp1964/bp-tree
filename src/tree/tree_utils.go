package tree

import (
	atomiccounter "bp-tree/src/atomic_counter"
	"bp-tree/src/diskmanager"
	"bp-tree/src/node"
	"bp-tree/src/nodemanager"
	"log"
	"os"
)

// NewTree ...
func NewTree(fileName string, order, blockSize int) *Tree {
	checkAndCreateFileIfNecessary(fileName)

	counter := atomiccounter.NewCounter()

	logger := log.Default()

	// create new node and assign it to tree
	rootNode := node.NewLeafNode(counter.Get(), order, logger)

	diskmanager := diskmanager.NewDiskManagerImpl(logger)

	// create new nodemanager object and assign it to tree
	nodemanager := &nodemanager.NodeManager{
		FileName:  fileName,
		BlockSize: blockSize,
		DM:        diskmanager,
		Logger:    logger,
	}
	nodemanager.Flush(rootNode)

	return &Tree{
		fileName:      fileName,
		rootNodeIndex: rootNode.Index,
		order:         order,
		blockSize:     blockSize,
		nodeManager:   nodemanager,
		logger:        logger,
		counter:       counter,
		leftMost:      rootNode.Index,
	}
}

// RecreateTree ...
func RecreateTree(fileName string, order, blockSize, rootNodeIndex, leftMostNodeIndex, numberOfNodes int) *Tree {
	logger := log.Default()
	counter := atomiccounter.NewCounterFromIndex(numberOfNodes)

	diskmanager := diskmanager.NewDiskManagerImpl(logger)

	// create new nodemanager object and assign it to tree
	nodemanager := &nodemanager.NodeManager{
		FileName:  fileName,
		BlockSize: blockSize,
		DM:        diskmanager,
		Logger:    logger,
	}

	return &Tree{
		fileName:      fileName,
		rootNodeIndex: rootNodeIndex,
		order:         order,
		blockSize:     blockSize,
		nodeManager:   nodemanager,
		logger:        logger,
		counter:       counter,
		leftMost:      leftMostNodeIndex,
	}
}

func checkAndCreateFileIfNecessary(fileName string) error {
	_, err := os.Stat(fileName)

	if err != nil {
		_, err = os.Create(fileName)
		return err
	}

	return os.Truncate(fileName, 0)
}
