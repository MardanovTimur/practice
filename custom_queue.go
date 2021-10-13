package main

import (
	"fmt"
	"yandex.practice/lib"
)

func main() {
	queue := lib.NewQueue(10)
	queue.Push("123")
	queue.Push("ABS")
	queue.Push(14)
	v, _ := queue.Pop()
	fmt.Printf("Array: value: %v\n", v)
	v, _ = queue.Pop()
	fmt.Printf("Array: value: %v\n", v)

	queueChannel := lib.NewQueueChannel(10)

	queueChannel.Push("A")
	queueChannel.Push("B")
	queueChannel.Push(14)
	v, _ = queueChannel.Pop()
	fmt.Printf("Channel: value: %v\n", v)
	v, _ = queueChannel.Pop()
	fmt.Printf("Channel: value: %v\n", v)
}
