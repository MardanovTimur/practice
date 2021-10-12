package main

import (
	"fmt"
	"yandex.practice/lib"
)

func main() {
	queueChannel := lib.NewQueueChannel(10)

	queueChannel.Push(12)
	queueChannel.Push(13)
	queueChannel.Push(14)
	v, _ := queueChannel.Pop()
	v, _ = queueChannel.Pop()
	fmt.Printf("%v\n", v)
}
