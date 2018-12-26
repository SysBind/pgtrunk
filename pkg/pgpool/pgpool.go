package pgpool

func GetPrimaryNode() Node {
	var node Node
	nodeCount := pcpNodeCount()
	for i := 0; i < nodeCount; i++ {
		node = pcpNodeInfo(i)
		if node.Role == "primary" { return node	}
	}
	return Node{}
}
