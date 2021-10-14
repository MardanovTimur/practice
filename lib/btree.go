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

func upgradeRoots(bRoot *BRoot) {
	for _, node := range bRoot.nodes {
		if node.left != nil {
			subRoot := node.left
			subRoot.parent = bRoot
			subRootR := node.right
			subRootR.parent = bRoot
		}
	}
}

func (bRoot *BRoot) splitRoot() *BRoot {
	mNode := bRoot.nodes[bRoot.getT()]
	medianNode := &BNode{mNode.value, mNode.data, nil, nil}

	var leftNodes, rightNodes []*BNode = make([]*BNode, 0), make([]*BNode, 0)
	leftNodes = append(leftNodes, bRoot.nodes[:bRoot.getT()]...)
	rightNodes = append(rightNodes, bRoot.nodes[bRoot.getT()+1:]...)

	medianNode.left = newRoot(bRoot, medianNode, leftNodes, bRoot.capacity)
	medianNode.right = newRoot(bRoot, medianNode, rightNodes, bRoot.capacity)

	upgradeRoots(medianNode.left)
	upgradeRoots(medianNode.right)

	if bRoot.isRoot {
		bRoot.nodes = []*BNode{medianNode}
		return bRoot
	} else {
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

func selectRoot(bRoot *BRoot, newNode *BNode) *BRoot {
	// Select root for insert node recursive
	if bRoot.isLeaf() {
		// handle initial situation when bRoot is a Root
		return bRoot
	}
	for _, node := range bRoot.nodes {
		if compareNodes(newNode, node) <= 0 {
			return selectRoot(node.left, newNode)
		}
	}
	last := bRoot.nodes[len(bRoot.nodes)-1].right
	if last != nil {
		return selectRoot(last, newNode)
	}
	return nil
}

func (bRoot *BRoot) insertNode(node *BNode, reverse bool) {
	var insertRoot *BRoot = bRoot
	if !reverse {
		// Если идет операция вставки ключа, а не перестроения дерева обратное (рекурсивное)
		insertRoot = selectRoot(bRoot, node)
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
		newNodes = append(newNodes, insertRoot.nodes[:insertP]...)
		// Ставим пограничные руты
		if reverse {
			newNodes[insertP-1].right = node.left
		}
	}
	newNodes = append(newNodes, node)
	if insertP < len(insertRoot.nodes) {
		newNodes = append(newNodes, insertRoot.nodes[insertP:]...)
		// Ставим пограничные руты
		if reverse {
			newNodes[insertP+1].left = node.right
		}
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
	// fmt.Printf("] %v - %v\n", bRoot, bRoot.parent)
	fmt.Printf("]\n")

	fmt.Print("  ")
	if bRoot.nodes[0].right != nil {
		bRoot.nodes[0].left.printBRootNodeValues()
	}
	for i = 0; i < len(bRoot.nodes); i++ {
		if bRoot.nodes[i].right != nil {
			bRoot.nodes[i].right.printBRootNodeValues()
		}
	}
}

func main() {
	tree, _ := NewBTree(3)
	fmt.Printf("Tree created %v\n", tree)
	for i := 0; i < 100; i++ {
		tree.AddKey(i, nil)
	}
	tree.AddKey(99, nil)
	tree.AddKey(99, nil)
	tree.printBRootNodeValues()
}
