package lib

func (node *BNode) WalkWide() []BNodeV {
	walkedNodes := make([]BNodeV, 0)
	queue := NewQueue(100000)
	queue.Push(node)
	for {
		qNode, err := queue.Pop()
		if err != nil {
			break
		}
		walkedNodes = append(walkedNodes, qNode.(*BNode).Value)
		if qNode.(*BNode).Left != nil {
			queue.Push(qNode.(*BNode).Left)
		}
		if qNode.(*BNode).Right != nil {
			queue.Push(qNode.(*BNode).Right)
		}
	}
	return walkedNodes
}
