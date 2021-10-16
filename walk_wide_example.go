package main

import (
	"fmt"
	"yandex.practice/lib"
)

func main() {
	bnode := lib.NewBNode(5)
	bnode.Left = lib.NewBNode(3)
	bnode.Right = lib.NewBNode(7)
	bnode.Left.Left = lib.NewBNode(2)
	bnode.Left.Right = lib.NewBNode(4)
	bnode.Right.Left = lib.NewBNode(6)
	bnode.Right.Right = lib.NewBNode(8)

	walkedNodes := bnode.WalkWide()
	fmt.Printf("%v\n", walkedNodes)
}
