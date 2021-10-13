package lib

import "errors"

type IQueue interface {
	Push(interface{})
	Pop() (interface{}, error)
}

type Queue struct {
	arr  []interface{}
	size int
}

func (q *Queue) Push(a interface{}) error {
	if len(q.arr) >= q.size {
		return errors.New("Queue overflow")
	}
	q.arr = append(q.arr, a)
	return nil
}

func (q *Queue) Pop() (interface{}, error) {
	if len(q.arr) <= 0 {
		return -1, errors.New("Queue underflow")
	}
	var item interface{} = q.arr[0]
	q.arr = q.arr[1:]
	return item, nil
}

func NewQueue(size int) *Queue {
	return &Queue{make([]interface{}, 0), size}
}
