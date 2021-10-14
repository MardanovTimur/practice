package main

import (
	"errors"
	"fmt"
	"strings"
)

type T interface{}
type V interface{}

type BNode struct {
	value T
	data  V
	left  *BRoot
	right *BRoot
}

func (bNode *BNode) GetValue() T {
	return bNode.value
}

func (bNode *BNode) GetData() V {
	return bNode.data
}

type BRoot struct {
	parent     *BRoot
	parentNode *BNode
	nodes      []*BNode
	capacity   int
	isRoot     bool
}

func (bRoot *BRoot) isLeaf() bool {
	var i = 0

	if len(bRoot.nodes) == 0 {
		return true
	}

	for i = 0; i < len(bRoot.nodes); i++ {
		if bRoot.nodes[i].left != nil {
			return false
		}
	}
	return bRoot.nodes[i-1].right == nil
}

func (bRoot *BRoot) isNodesEmpty() bool {
	return len(bRoot.nodes) == 0
}

func (bRoot *BRoot) getT() int {
	return bRoot.capacity
}

func (bRoot *BRoot) hasFreeSpace() bool {
	return len(bRoot.nodes) < 2*bRoot.getT()
}

func (bRoot *BRoot) splitRoot() *BRoot {
	medianNode := bRoot.nodes[bRoot.getT()]
	var leftNodes, rightNodes []*BNode = make([]*BNode, bRoot.getT()), make([]*BNode, len(bRoot.nodes[bRoot.getT()+1:]))
	copy(leftNodes, bRoot.nodes[:bRoot.getT()])
	copy(rightNodes, bRoot.nodes[bRoot.getT()+1:])

	medianNode.left = newRoot(bRoot, medianNode, leftNodes, bRoot.capacity)
	medianNode.right = newRoot(bRoot, medianNode, rightNodes, bRoot.capacity)

	if bRoot.isRoot {
		bRoot.nodes = make([]*BNode, 1)
		bRoot.nodes[0] = medianNode
		return bRoot
	} else if bRoot.parent != nil {
		medianNode.left.parent = bRoot.parent
		medianNode.right.parent = bRoot.parent
		bRoot.parent.insertNode(medianNode, true)
	}
	return bRoot.parent
}

func compareNodes(node1, node2 *BNode) int {
	switch node1.value.(type) {
	case int:
		if node1.value.(int) > node2.value.(int) {
			return 1
		} else if node1.value.(int) == node2.value.(int) {
			return 0
		} else {
			return -1
		}
	case string:
		return strings.Compare(node1.value.(string), node2.value.(string))
	default:
		return 0
	}
}

func (bRoot *BRoot) selectRoot(newNode *BNode) *BRoot {
	// Select root for insert node recursive
	if bRoot.isLeaf() {
		// handle initial situation when bRoot is a Root
		return bRoot
	}
	for _, node := range bRoot.nodes {
		if node.left == nil {
			continue
		}
		if compareNodes(newNode, node) <= 0 {
			return node.left.selectRoot(newNode)
		}
	}
	last := bRoot.nodes[len(bRoot.nodes)-1].right
	if last != nil {
		return last.selectRoot(newNode)
	}

	for _, node := range bRoot.nodes {
		fmt.Printf("%v ,", node)
	}
	fmt.Printf("\nBroot not leaf is: %v", bRoot)
	return nil
}

func (bRoot *BRoot) insertNode(node *BNode, reverse bool) {
	insertRoot := bRoot
	if !reverse {
		// Если идет операция вставки ключа, а не перестроения дерева обратное (рекурсивное)
		insertRoot = bRoot.selectRoot(node)
	}
	var insertP int = len(insertRoot.nodes)
	for i := 0; i < len(insertRoot.nodes); i++ {
		if compareNodes(node, insertRoot.nodes[i]) <= 0 {
			insertP = i
			break
		}
	}
	// Вставка новой ноды и обновление програничников
	newNodes := make([]*BNode, 0)
	if insertP > 0 {
		newNodes = insertRoot.nodes[:insertP]
		// Ставим пограничные руты
		insertRoot.nodes[insertP-1].right = node.left
	}
	newNodes = append(newNodes, node)
	if insertP < len(insertRoot.nodes) {
		newNodes = append(newNodes, insertRoot.nodes[insertP:]...)
		// Ставим пограничные руты
		insertRoot.nodes[insertP].left = node.right
	}
	insertRoot.nodes = newNodes

	if !insertRoot.hasFreeSpace() {
		insertRoot.splitRoot()
	}
}

func (bRoot *BRoot) AddKey(value T, data V) {
	var node *BNode = &BNode{value, data, nil, nil}

	bRoot.insertNode(node, false)
}

func (bRoot *BRoot) search(value T) (*BNode, error) {
	// var searchNode *BNode = &BNode{value, nil, nil, nil}
	// TODO finish it
	for _, node := range bRoot.nodes {
		if node.value == value {
			return node, nil
		}
	}
	return nil, errors.New(fmt.Sprintf("Key %v not found", value))
}

func newRoot(parentRoot *BRoot, parentNode *BNode, nodes []*BNode, capacity int) *BRoot {
	return &BRoot{
		parentRoot,
		parentNode,
		nodes,
		capacity,
		false,
	}
}

// capacity > 1;
func NewBTree(capacity int) (*BRoot, error) {
	if capacity <= 1 {
		return nil, errors.New("Capacity must greater than 2")
	}
	var broot *BRoot = newRoot(
		nil,
		nil,
		make([]*BNode, 0),
		capacity,
	)
	broot.isRoot = true
	return broot, nil
}

func (bRoot *BRoot) printBRootNodeValues() {
	fmt.Printf("[")
	var i int = 0
	for i = 0; i < len(bRoot.nodes); i++ {
		fmt.Printf("%v, ", bRoot.nodes[i].value)
	}
	fmt.Printf("]\n")
	for i = 0; i < len(bRoot.nodes); i++ {
		if bRoot.nodes[i].left != nil {
			bRoot.nodes[i].left.printBRootNodeValues()
		}
	}
	if bRoot.nodes[i-1].right != nil {
		bRoot.nodes[i-1].right.printBRootNodeValues()
	}
}

func main() {
	tree, _ := NewBTree(3)
	fmt.Printf("Tree created %v\n", tree)
	tree.AddKey(1, nil)
	tree.AddKey(2, nil)
	tree.AddKey(3, nil)
	tree.AddKey(4, nil)
	tree.AddKey(5, nil)
	tree.AddKey(6, nil)
	tree.AddKey(7, nil)
	tree.AddKey(-1, nil)
	tree.AddKey(6, nil)
	tree.AddKey(7, nil)
	tree.AddKey(-1, nil)
	tree.AddKey(6, nil)
	tree.AddKey(7, nil)
	tree.AddKey(-1, nil)
	tree.AddKey(6, nil)
	tree.AddKey(7, nil)
	tree.AddKey(-1, nil)
	tree.AddKey(6, nil)
	tree.AddKey(6, nil)
	tree.AddKey(6, nil)
	tree.printBRootNodeValues()
}
