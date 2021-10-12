package lib

import "errors"

type IQueueChannel interface {
	Push(int)
	Pop() (int, error)
}

type QueueChannel struct {
	channel chan int
	size    int
}

func (q *QueueChannel) Push(a int) error {
	if len(q.channel) >= q.size {
		return errors.New("Queue overflow")
	}
	q.channel <- a
	return nil
}

func (q *QueueChannel) Pop() (int, error) {
	if len(q.channel) <= 0 {
		return -1, errors.New("Queue underflow")
	}
	return <-q.channel, nil
}

func NewQueueChannel(size int) *QueueChannel {
	return &QueueChannel{make(chan int, size), size}
}
