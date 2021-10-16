package lib

type BNodeV interface{}

type BNode struct {
	Value  BNodeV
	Parent *BNode
	Left   *BNode
	Right  *BNode
}

func NewBNode(v BNodeV) *BNode {
	return &BNode{
		v,
		nil,
		nil,
		nil,
	}
}
