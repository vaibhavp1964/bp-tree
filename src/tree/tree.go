package tree

import (
	atomiccounter "bp-tree/src/atomic_counter"
	"bp-tree/src/nodemanager"
	"log"
)

// Tree ...
type Tree struct {
	fileName      string
	rootNodeIndex int
	order         int
	blockSize     int
	nodeManager   *nodemanager.NodeManager
	logger        *log.Logger
	counter       *atomiccounter.AtomicCounter
	leftMost      int
}
