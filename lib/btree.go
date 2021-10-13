package lib

import (
	"errors"
	"fmt"
	"strings"
)

type BNode struct {
	value interface{}
	data  interface{}
	left  *BRoot
	right *BRoot
}

type BRoot struct {
	parent *BRoot
	nodes  []*BNode
}

func (bRoot *BRoot) search(value interface{}) (*BNode, error) {
	for _, node := range bRoot.nodes {
		if node.value == value {
			return node, nil
		}
		switch node.value.(type) {
		case int:
			if node.value.(int) > value.(int) {
				return node.left.search(value)
			}
		case string:
			if strings.Compare(node.value.(string), value.(string)) == 1 {
				return node.left.search(value)
			}
		}
	}
	return nil, errors.New(fmt.Sprintf("Key %v not found", value))
}

// capacity > 1;
func NewBTree(capacity int) (*BRoot, error) {
	if capacity > 1 {
		return nil, errors.New("Capacity must greater than 2")
	}
	var broot *BRoot = &BRoot{
		nil,
		make([]*BNode, 2*capacity-1),
	}
	return broot, nil
}
