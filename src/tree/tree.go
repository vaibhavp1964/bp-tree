package tree

import (
	"log"

	atomiccounter "github.com/vaibhavp1964/bp-tree/src/atomic_counter"
	"github.com/vaibhavp1964/bp-tree/src/nodemanager"
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
