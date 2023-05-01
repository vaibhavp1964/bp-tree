package tree

// ConstructTreeFromFile ...
func ConstructTreeFromFile(fileName string, order, blockSize int, rootNodeIndex, leftMost, numberOfNodes int) *Tree {
	return RecreateTree(fileName, order, blockSize, rootNodeIndex, leftMost, numberOfNodes)
}
