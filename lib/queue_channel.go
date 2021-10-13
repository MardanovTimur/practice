package lib

import "errors"

type IQueueChannel interface {
	Push(interface{})
	Pop() (interface{}, error)
}

type QueueChannel struct {
	channel chan interface{}
	size    int
}

func (q *QueueChannel) Push(a interface{}) error {
	if len(q.channel) >= q.size {
		return errors.New("Queue overflow")
	}
	q.channel <- a
	return nil
}

func (q *QueueChannel) Pop() (interface{}, error) {
	if len(q.channel) <= 0 {
		return -1, errors.New("Queue underflow")
	}
	return <-q.channel, nil
}

func NewQueueChannel(size int) *QueueChannel {
	return &QueueChannel{make(chan interface{}, size), size}
}
