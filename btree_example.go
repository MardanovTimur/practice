package main

import (
	"fmt"
	"yandex.practice/lib"
)

func main() {
	tree, _ := lib.NewBTree(3)
	fmt.Printf("Tree created %v\n", tree)
	for i := 0; i < 100; i++ {
		tree.AddKey(i, nil)
	}
	tree.AddKey(99, nil)
	tree.AddKey(99, nil)

	element, err := tree.Search(99)
	if err == nil {
		fmt.Printf("Found %v, \n", element)
	} else {
		fmt.Println(err)
	}
}
